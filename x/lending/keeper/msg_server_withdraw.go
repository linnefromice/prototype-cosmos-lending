package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/linnefromice/lending/x/lending/types"
)

func (k msgServer) Withdraw(goCtx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.Keeper.WithdrawFromPairPool(ctx, msg)
	if err != nil {
		panic(err)
	}

	return &types.MsgWithdrawResponse{}, nil
}
