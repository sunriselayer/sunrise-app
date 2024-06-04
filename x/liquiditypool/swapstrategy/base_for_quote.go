package swapstrategy

import (
	"fmt"

	"cosmossdk.io/store/prefix"
	dbm "github.com/cometbft/cometbft-db"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"cosmossdk.io/math"
	"github.com/sunriselayer/sunrise/x/liquiditypool/types"

	"cosmossdk.io/core/store"
)

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

func (s baseForQuoteStrategy) SetLiquidityDeltaSign(deltaLiquidity math.LegacyDec) math.LegacyDec {
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
