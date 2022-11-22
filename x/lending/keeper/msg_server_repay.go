package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/linnefromice/lending/x/lending/types"
)

func (k msgServer) Repay(goCtx context.Context, msg *types.MsgRepay) (*types.MsgRepayResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.Keeper.RepayToPairPool(ctx, msg)
	if err != nil {
		panic(err)
	}

	return &types.MsgRepayResponse{}, nil
}
