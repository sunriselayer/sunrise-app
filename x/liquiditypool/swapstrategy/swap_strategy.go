package swapstrategy

import (
	"fmt"

	"cosmossdk.io/core/store"
	"cosmossdk.io/math"
	dbm "github.com/cometbft/cometbft-db"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sunriselayer/sunrise/x/liquiditypool/types"
)

type SwapStrategy interface {
	GetSqrtTargetPrice(nextTickSqrtPrice math.LegacyDec) math.LegacyDec
	ComputeSwapWithinBucketOutGivenIn(sqrtPriceCurrent, sqrtPriceTarget math.LegacyDec, liquidity, amountRemainingIn math.LegacyDec) (sqrtPriceNext math.LegacyDec, amountInConsumed, amountOutComputed, feeChargeTotal math.LegacyDec)
	ComputeSwapWithinBucketInGivenOut(sqrtPriceCurrent, sqrtPriceTarget math.LegacyDec, liquidity, amountRemainingOut math.LegacyDec) (sqrtPriceNext math.LegacyDec, amountOutConsumed, amountInComputed, feeChargeTotal math.LegacyDec)
	NextTickIterator(ctx sdk.Context, poolId uint64, tickIndex int64) dbm.Iterator
	SetLiquidityDeltaSign(liquidityDelta math.LegacyDec) math.LegacyDec
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
