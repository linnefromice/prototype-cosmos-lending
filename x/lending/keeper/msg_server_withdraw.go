package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/linnefromice/lending/x/lending/types"
)

func (k msgServer) Withdraw(goCtx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	res, err := k.Keeper.WithdrawFromPairPool(ctx, msg)
	if err != nil {
		panic(err)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.PoolWithdrawedEventType,
			sdk.NewAttribute(types.PoolEventId, fmt.Sprint(res.Id)),
			sdk.NewAttribute(types.PoolEventAmount, fmt.Sprint(res.Id))),
	)

	return &types.MsgWithdrawResponse{}, nil
}
