package keeper

import (
	"errors"
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	lptypes "github.com/sunriselayer/sunrise/x/liquiditypool/types"
	"github.com/sunriselayer/sunrise/x/swap/types"
)

func (k Keeper) SwapExactAmountInNoTakerFee(
	ctx sdk.Context,
	sender sdk.AccAddress,
	poolId uint64,
	tokenIn sdk.Coin,
	tokenOutDenom string,
	tokenOutMinAmount math.Int,
) (tokenOutAmount math.Int, err error) {
	pool, found := k.liquidityPoolKeeper.GetPool(ctx, poolId)
	if !found {
		return math.Int{}, lptypes.ErrPoolNotFound
	}

	// routeStep to the pool-specific SwapExactAmountIn implementation.
	tokenOutAmount, err = k.liquidityPoolKeeper.SwapExactAmountIn(ctx, sender, pool, tokenIn, tokenOutDenom)
	if err != nil {
		return math.Int{}, err
	}

	return tokenOutAmount, nil
}

func (k Keeper) MultihopEstimateOutGivenExactAmountInNoTakerFee(
	ctx sdk.Context,
	route []types.Route,
	tokenIn sdk.Coin,
) (tokenOutAmount math.Int, err error) {
	return k.multihopEstimateOutGivenExactAmountInInternal(ctx, route, tokenIn, false)
}

func (k Keeper) MultihopEstimateOutGivenExactAmountIn(
	ctx sdk.Context,
	route []types.Route,
	tokenIn sdk.Coin,
) (tokenOutAmount math.Int, err error) {
	return k.multihopEstimateOutGivenExactAmountInInternal(ctx, route, tokenIn, true)
}

func (k Keeper) multihopEstimateOutGivenExactAmountInInternal(
	ctx sdk.Context,
	route []types.Route,
	tokenIn sdk.Coin,
	applyTakerFee bool,
) (tokenOutAmount math.Int, err error) {
	// recover from panic
	defer func() {
		if r := recover(); r != nil {
			tokenOutAmount = math.Int{}
			err = fmt.Errorf("function MultihopEstimateOutGivenExactAmountIn failed due to internal reason: %v", r)
		}
	}()

	for _, routeStep := range route {
		switch strategy := routeStep.Strategy.(type) {
		case *types.Route_Pool:
			pool, found := k.liquidityPoolKeeper.GetPool(ctx, strategy.Pool.PoolId)
			if !found {
				return math.Int{}, lptypes.ErrPoolNotFound
			}

			actualTokenIn := tokenIn

			tokenOutAmount, err := k.liquidityPoolKeeper.CalculateResultExactAmountIn(ctx, pool, actualTokenIn, routeStep.DenomOut)
			if err != nil {
				return math.Int{}, err
			}

			if !tokenOutAmount.IsPositive() {
				return math.Int{}, errors.New("token amount must be positive")
			}

			// Chain output of current pool as the input for the next routed pool
			// We don't need to validate the denom,
			// as CalculateResultExactAmountIn is responsible for ensuring the denom exists in the pool.
			tokenIn = sdk.Coin{Denom: routeStep.DenomOut, Amount: tokenOutAmount}
		case *types.Route_Series:
			panic("not implemented strategy")
		case *types.Route_Parallel:
			panic("not implemented strategy")
		}
	}
	return tokenOutAmount, err
}

func (k Keeper) RouteExactAmountOut(ctx sdk.Context,
	sender sdk.AccAddress,
	route []types.Route,
	tokenInMaxAmount math.Int,
	tokenOut sdk.Coin,
) (tokenInAmount math.Int, err error) {
	defer func() {
		if r := recover(); r != nil {
			tokenInAmount = math.Int{}
			err = fmt.Errorf("function RouteExactAmountOut failed due to internal reason: %v", r)
		}
	}()

	var insExpected []math.Int
	insExpected, err = k.createMultihopExpectedSwapOuts(ctx, route, tokenOut)

	if err != nil {
		return math.Int{}, err
	}
	if len(insExpected) == 0 {
		return math.Int{}, nil
	}
	insExpected[0] = tokenInMaxAmount

	// Iterates through each routed pool and executes their respective swaps. Note that all of the work to get the return
	// value of this method is done when we calculate insExpected – this for loop primarily serves to execute the actual
	// swaps on each pool.
	for i, routeStep := range route {
		switch strategy := routeStep.Strategy.(type) {
		case *types.Route_Pool:
			pool, found := k.liquidityPoolKeeper.GetPool(ctx, strategy.Pool.PoolId)
			if !found {
				return math.Int{}, lptypes.ErrPoolNotFound
			}

			_tokenOut := tokenOut

			// If there is one pool left in the routeStep, set the expected output of the current swap
			// to the estimated input of the final pool.
			if i != len(route)-1 {
				_tokenOut = sdk.NewCoin(route[i+1].DenomIn, insExpected[i+1])
			}

			curTokenInAmount, swapErr := k.liquidityPoolKeeper.SwapExactAmountOut(ctx, sender, pool, _tokenOut, routeStep.DenomIn)
			if swapErr != nil {
				return math.Int{}, swapErr
			}

			tokenIn := sdk.NewCoin(routeStep.DenomIn, curTokenInAmount)

			if i == 0 {
				tokenInAmount = tokenIn.Amount
			}
		case *types.Route_Series:
			panic("not implemented strategy")
		case *types.Route_Parallel:
			panic("not implemented strategy")
		}
	}

	return tokenInAmount, nil
}

func (k Keeper) MultihopEstimateInGivenExactAmountOut(
	ctx sdk.Context,
	route []types.Route,
	tokenOut sdk.Coin,
) (tokenInAmount math.Int, err error) {
	var insExpected []math.Int

	// recover from panic
	defer func() {
		if r := recover(); r != nil {
			insExpected = []math.Int{}
			err = fmt.Errorf("function MultihopEstimateInGivenExactAmountOut failed due to internal reason: %v", r)
		}
	}()

	// Determine what the estimated input would be for each pool along the multi-hop route
	insExpected, err = k.createMultihopExpectedSwapOuts(ctx, route, tokenOut)
	if err != nil {
		return math.Int{}, err
	}
	if len(insExpected) == 0 {
		return math.Int{}, nil
	}

	return insExpected[0], nil
}

func (k Keeper) createMultihopExpectedSwapOuts(
	ctx sdk.Context,
	route []types.Route,
	tokenOut sdk.Coin,
) ([]math.Int, error) {
	insExpected := make([]math.Int, len(route))
	for i := len(route) - 1; i >= 0; i-- {
		routeStep := route[i]
		switch strategy := routeStep.Strategy.(type) {
		case *types.Route_Pool:
			pool, found := k.liquidityPoolKeeper.GetPool(ctx, strategy.Pool.PoolId)
			if !found {
				return nil, lptypes.ErrPoolNotFound
			}

			tokenIn, err := k.liquidityPoolKeeper.CalculateResultExactAmountOut(ctx, pool, tokenOut, routeStep.DenomIn)
			if err != nil {
				return nil, err
			}

			insExpected[i] = tokenIn
			tokenOut = sdk.NewCoin(routeStep.DenomIn, tokenIn)
		case *types.Route_Series:
			panic("not implemented strategy")
		case *types.Route_Parallel:
			panic("not implemented strategy")
		}
	}

	return insExpected, nil
}
