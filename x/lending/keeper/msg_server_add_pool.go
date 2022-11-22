package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/linnefromice/lending/x/lending/types"
)

func (k msgServer) AddPool(goCtx context.Context, msg *types.MsgAddPool) (*types.MsgAddPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pool, err := k.Keeper.AddPool(ctx, msg)
	if err != nil {
		panic(err)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.PoolAddedEventType,
			sdk.NewAttribute(types.PoolEventId, fmt.Sprint(pool.Id)),
			sdk.NewAttribute(types.PoolAddedEventAssetDenom, pool.AssetLiquidity.Denom),
			sdk.NewAttribute(types.PoolAddedEventShadowDenom, pool.ShadowLiquidity.Denom),
		),
	)

	return &types.MsgAddPoolResponse{
		PoolId: pool.PoolId,
	}, nil
}
