package keeper

import (
	"fmt"

	"cosmossdk.io/core/store"
	"cosmossdk.io/math"
	"cosmossdk.io/store/prefix"
	dbm "github.com/cometbft/cometbft-db"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sunriselayer/sunrise/x/liquiditypool/types"
)

type SwapStrategy interface {
	GetSqrtTargetPrice(nextTickSqrtPrice math.LegacyDec) math.LegacyDec
	ComputeSwapWithinBucketOutGivenIn(sqrtPriceCurrent, sqrtPriceTarget math.LegacyDec, liquidity, amountRemainingIn math.LegacyDec) (sqrtPriceNext math.LegacyDec, amountInConsumed, amountOutComputed, feeChargeTotal math.LegacyDec)
	ComputeSwapWithinBucketInGivenOut(sqrtPriceCurrent, sqrtPriceTarget math.LegacyDec, liquidity, amountRemainingOut math.LegacyDec) (sqrtPriceNext math.LegacyDec, amountOutConsumed, amountInComputed, feeChargeTotal math.LegacyDec)
	NextTickIterator(ctx sdk.Context, poolId uint64, tickIndex int64) dbm.Iterator
	GetLiquidityDeltaSign(liquidityDelta math.LegacyDec) math.LegacyDec
	NextTickAfterCrossing(nextTick int64) (updatedNextTick int64)
	ValidateSqrtPrice(sqrtPriceLimit math.LegacyDec, currentSqrtPrice math.LegacyDec) error
	BaseForQuote() bool
}

func New(baseForQuote bool, sqrtPriceLimit math.LegacyDec, storeService store.KVStoreService, feeRate math.LegacyDec) SwapStrategy {
	if baseForQuote {
		return &baseForQuoteStrategy{sqrtPriceLimit: sqrtPriceLimit, storeService: storeService, feeRate: feeRate}
	}
	return &quoteForBaseStrategy{sqrtPriceLimit: sqrtPriceLimit, storeService: storeService, feeRate: feeRate}
}

func GetPriceLimit(baseForQuote bool) math.LegacyDec {
	if baseForQuote {
		return types.MinSpotPrice
	}
	return types.MaxSpotPrice
}

func GetSqrtPriceLimit(priceLimit math.LegacyDec, baseForQuote bool) (math.LegacyDec, error) {
	if priceLimit.IsZero() {
		if baseForQuote {
			return types.MinSqrtPrice, nil
		}
		return types.MaxSqrtPrice, nil
	}

	if priceLimit.LT(types.MinSpotPrice) || priceLimit.GT(types.MaxSpotPrice) {
		return math.LegacyDec{}, types.ErrPriceOutOfBound
	}

	sqrtPriceLimit, err := priceLimit.ApproxSqrt()
	if err != nil {
		return math.LegacyDec{}, err
	}

	return sqrtPriceLimit, nil
}

func getFeeRateOverOneMinusFeeRate(feeRate math.LegacyDec) math.LegacyDec {
	return feeRate.QuoRoundUp(math.LegacyOneDec().Sub(feeRate))
}

func computeFeeChargeFromInAmount(amountIn math.LegacyDec, feeRateOveroneMinusFeeRate math.LegacyDec) math.LegacyDec {
	return amountIn.MulRoundUp(feeRateOveroneMinusFeeRate)
}

func computeFeeChargePerSwapStepOutGivenIn(hasReachedTarget bool, amountIn, amountSpecifiedRemaining, feeRate math.LegacyDec) math.LegacyDec {
	if feeRate.IsZero() {
		return math.LegacyZeroDec()
	} else if feeRate.IsNegative() {
		panic(fmt.Errorf("fee rate must be non-negative, was (%s)", feeRate))
	}

	var feeChargeTotal math.LegacyDec
	if hasReachedTarget {
		feeChargeTotal = computeFeeChargeFromInAmount(amountIn, getFeeRateOverOneMinusFeeRate(feeRate))
	} else {
		feeChargeTotal = amountSpecifiedRemaining.Sub(amountIn)
	}

	if feeChargeTotal.IsNegative() {
		panic(fmt.Errorf("fee rate charge must be non-negative, was (%s)", feeChargeTotal))
	}

	return feeChargeTotal
}

type baseForQuoteStrategy struct {
	sqrtPriceLimit math.LegacyDec
	storeService   store.KVStoreService
	feeRate        math.LegacyDec
}

var _ SwapStrategy = (*baseForQuoteStrategy)(nil)

func (s baseForQuoteStrategy) BaseForQuote() bool { return true }

func (s baseForQuoteStrategy) GetSqrtTargetPrice(nextTickSqrtPrice math.LegacyDec) math.LegacyDec {
	if nextTickSqrtPrice.LT(s.sqrtPriceLimit) {
		return s.sqrtPriceLimit
	}
	return nextTickSqrtPrice
}

func (s baseForQuoteStrategy) ComputeSwapWithinBucketOutGivenIn(sqrtPriceCurrent, sqrtPriceTarget math.LegacyDec, liquidity, amountBaseInRemaining math.LegacyDec) (math.LegacyDec, math.LegacyDec, math.LegacyDec, math.LegacyDec) {
	amountBaseIn := types.CalcAmountBaseDelta(liquidity, sqrtPriceTarget, sqrtPriceCurrent, true)
	amountBaseInAfterFee := amountBaseInRemaining.Mul(math.LegacyOneDec().Sub(s.feeRate))

	var sqrtPriceNext math.LegacyDec
	if amountBaseInAfterFee.GTE(amountBaseIn) {
		sqrtPriceNext = sqrtPriceTarget
	} else {
		sqrtPriceNext = types.GetNextSqrtPriceFromAmountBaseInRoundingUp(sqrtPriceCurrent, liquidity, amountBaseInAfterFee)
	}

	hasReachedTarget := sqrtPriceTarget.Equal(sqrtPriceNext)
	if !hasReachedTarget {
		amountBaseIn = types.CalcAmountBaseDelta(liquidity, sqrtPriceNext, sqrtPriceCurrent, true)
	}

	amountQuoteOut := types.CalcAmountQuoteDelta(liquidity, sqrtPriceNext, sqrtPriceCurrent, false)

	feeChargeTotal := computeFeeChargePerSwapStepOutGivenIn(hasReachedTarget, amountBaseIn, amountBaseInRemaining, s.feeRate)
	return sqrtPriceNext, amountBaseIn, amountQuoteOut, feeChargeTotal
}

func (s baseForQuoteStrategy) ComputeSwapWithinBucketInGivenOut(sqrtPriceCurrent, sqrtPriceTarget math.LegacyDec, liquidity, amountQuoteRemainingOut math.LegacyDec) (math.LegacyDec, math.LegacyDec, math.LegacyDec, math.LegacyDec) {
	amountQuoteOut := types.CalcAmountQuoteDelta(liquidity, sqrtPriceTarget, sqrtPriceCurrent, false)

	var sqrtPriceNext math.LegacyDec
	if amountQuoteRemainingOut.GTE(amountQuoteOut) {
		sqrtPriceNext = sqrtPriceTarget
	} else {
		sqrtPriceNext = types.GetNextSqrtPriceFromAmountQuoteOutRoundingDown(sqrtPriceCurrent, liquidity, amountQuoteRemainingOut)
	}

	hasReachedTarget := sqrtPriceTarget.Equal(sqrtPriceNext)

	if !hasReachedTarget {
		amountQuoteOut = types.CalcAmountQuoteDelta(liquidity, sqrtPriceNext, sqrtPriceCurrent, false)
	}

	amountBaseIn := types.CalcAmountBaseDelta(liquidity, sqrtPriceNext, sqrtPriceCurrent, true)

	feeChargeTotal := computeFeeChargeFromInAmount(amountBaseIn, getFeeRateOverOneMinusFeeRate(s.feeRate))

	if amountQuoteOut.GT(amountQuoteRemainingOut) {
		amountQuoteOut = amountQuoteRemainingOut
	}

	return sqrtPriceNext, amountQuoteOut, amountBaseIn, feeChargeTotal
}

func (s baseForQuoteStrategy) NextTickIterator(ctx sdk.Context, poolId uint64, currentTickIndex int64) dbm.Iterator {
	storeAdapter := runtime.KVStoreAdapter(s.storeService.OpenKVStore(ctx))
	prefixBz := types.KeyTickPrefixByPoolId(poolId)
	prefixStore := prefix.NewStore(storeAdapter, prefixBz)
	startKey := types.TickIndexToBytes(currentTickIndex + 1)
	iter := prefixStore.ReverseIterator(nil, startKey)

	for ; iter.Valid(); iter.Next() {
		tick, err := types.TickIndexFromBytes(iter.Key())
		if err != nil {
			iter.Close()
			panic(fmt.Errorf("invalid tick index (%s): %v", string(iter.Key()), err))
		}
		if tick <= currentTickIndex {
			break
		}
	}
	return iter
}

func (s baseForQuoteStrategy) GetLiquidityDeltaSign(deltaLiquidity math.LegacyDec) math.LegacyDec {
	return deltaLiquidity.Neg()
}

func (s baseForQuoteStrategy) NextTickAfterCrossing(nextTick int64) int64 {
	return nextTick - 1
}

func (s baseForQuoteStrategy) ValidateSqrtPrice(sqrtPrice math.LegacyDec, currentSqrtPrice math.LegacyDec) error {
	if sqrtPrice.GT(currentSqrtPrice) || sqrtPrice.LT(types.MinSqrtPrice) {
		return types.ErrInvalidSqrtPrice
	}
	return nil
}

type quoteForBaseStrategy struct {
	sqrtPriceLimit math.LegacyDec
	storeService   store.KVStoreService
	feeRate        math.LegacyDec
}

var _ SwapStrategy = (*quoteForBaseStrategy)(nil)

func (s quoteForBaseStrategy) BaseForQuote() bool { return false }

func (s quoteForBaseStrategy) GetSqrtTargetPrice(nextTickSqrtPrice math.LegacyDec) math.LegacyDec {
	if nextTickSqrtPrice.GT(s.sqrtPriceLimit) {
		return s.sqrtPriceLimit
	}
	return nextTickSqrtPrice
}

func (s quoteForBaseStrategy) ComputeSwapWithinBucketOutGivenIn(sqrtPriceCurrent, sqrtPriceTarget math.LegacyDec, liquidity, amountQuoteInRemaining math.LegacyDec) (math.LegacyDec, math.LegacyDec, math.LegacyDec, math.LegacyDec) {
	amountQuoteIn := types.CalcAmountQuoteDelta(liquidity, sqrtPriceTarget, sqrtPriceCurrent, true)

	amountQuoteInAfterFee := amountQuoteInRemaining.Mul(math.LegacyOneDec().Sub(s.feeRate))

	var sqrtPriceNext math.LegacyDec
	if amountQuoteInAfterFee.GTE(amountQuoteIn) {
		sqrtPriceNext = sqrtPriceTarget
	} else {
		sqrtPriceNext = types.GetNextSqrtPriceFromAmountQuoteInRoundingDown(sqrtPriceCurrent, liquidity, amountQuoteInAfterFee)
	}

	hasReachedTarget := sqrtPriceTarget.Equal(sqrtPriceNext)

	if !hasReachedTarget {
		amountQuoteIn = types.CalcAmountQuoteDelta(liquidity, sqrtPriceNext, sqrtPriceCurrent, true)
	}

	amountBaseOut := types.CalcAmountBaseDelta(liquidity, sqrtPriceNext, sqrtPriceCurrent, false)

	feeChargeTotal := computeFeeChargePerSwapStepOutGivenIn(hasReachedTarget, amountQuoteIn, amountQuoteInRemaining, s.feeRate)
	return sqrtPriceNext, amountQuoteIn, amountBaseOut, feeChargeTotal
}

func (s quoteForBaseStrategy) ComputeSwapWithinBucketInGivenOut(sqrtPriceCurrent, sqrtPriceTarget math.LegacyDec, liquidity, amountBaseRemainingOut math.LegacyDec) (math.LegacyDec, math.LegacyDec, math.LegacyDec, math.LegacyDec) {
	amountBaseOut := types.CalcAmountBaseDelta(liquidity, sqrtPriceTarget, sqrtPriceCurrent, false)

	var sqrtPriceNext math.LegacyDec
	if amountBaseRemainingOut.GTE(amountBaseOut) {
		sqrtPriceNext = sqrtPriceTarget
	} else {
		sqrtPriceNext = types.GetNextSqrtPriceFromAmountBaseOutRoundingUp(sqrtPriceCurrent, liquidity, amountBaseRemainingOut)
	}

	hasReachedTarget := sqrtPriceTarget.Equal(sqrtPriceNext)

	if !hasReachedTarget {
		amountBaseOut = types.CalcAmountBaseDelta(liquidity, sqrtPriceNext, sqrtPriceCurrent, false)
	}

	amountQuoteIn := types.CalcAmountQuoteDelta(liquidity, sqrtPriceNext, sqrtPriceCurrent, true)
	feeChargeTotal := computeFeeChargeFromInAmount(amountQuoteIn, getFeeRateOverOneMinusFeeRate(s.feeRate))

	if amountBaseOut.GT(amountBaseRemainingOut) {
		amountBaseOut = amountBaseRemainingOut
	}

	return sqrtPriceNext, amountBaseOut, amountQuoteIn, feeChargeTotal
}

func (s quoteForBaseStrategy) NextTickIterator(ctx sdk.Context, poolId uint64, currentTickIndex int64) dbm.Iterator {
	storeAdapter := runtime.KVStoreAdapter(s.storeService.OpenKVStore(ctx))
	prefixBz := types.KeyTickPrefixByPoolId(poolId)
	prefixStore := prefix.NewStore(storeAdapter, prefixBz)
	startKey := types.TickIndexToBytes(currentTickIndex)
	iter := prefixStore.Iterator(startKey, nil)

	for ; iter.Valid(); iter.Next() {
		tick, err := types.TickIndexFromBytes(iter.Key())
		if err != nil {
			iter.Close()
			panic(fmt.Errorf("invalid tick index (%s): %v", string(iter.Key()), err))
		}

		if tick > currentTickIndex {
			break
		}
	}
	return iter
}

func (s quoteForBaseStrategy) GetLiquidityDeltaSign(deltaLiquidity math.LegacyDec) math.LegacyDec {
	return deltaLiquidity
}

func (s quoteForBaseStrategy) NextTickAfterCrossing(nextTick int64) int64 {
	return nextTick
}

func (s quoteForBaseStrategy) ValidateSqrtPrice(sqrtPrice math.LegacyDec, currentSqrtPrice math.LegacyDec) error {
	if sqrtPrice.LT(currentSqrtPrice) || sqrtPrice.GT(types.MaxSqrtPrice) {
		return types.ErrInvalidSqrtPrice
	}
	return nil
}
