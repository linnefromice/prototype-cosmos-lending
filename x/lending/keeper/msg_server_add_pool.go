package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/linnefromice/lending/x/lending/types"
)

func (k msgServer) AddPool(goCtx context.Context, msg *types.MsgAddPool) (*types.MsgAddPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pool, err := k.Keeper.AddPool(ctx, msg)
	if err != nil {
		panic(err)
	}

	return &types.MsgAddPoolResponse{
		PoolId: pool.PoolId,
	}, nil
}
