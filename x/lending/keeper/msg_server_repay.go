package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/linnefromice/lending/x/lending/types"
)

func (k msgServer) Repay(goCtx context.Context, msg *types.MsgRepay) (*types.MsgRepayResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	res, err := k.Keeper.RepayToPairPool(ctx, msg)
	if err != nil {
		panic(err)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.PoolBorrowedEventType,
			sdk.NewAttribute(types.PoolEventId, fmt.Sprint(res.Id)),
			sdk.NewAttribute(types.PoolEventAmount, fmt.Sprint(res.Id))),
	)

	return &types.MsgRepayResponse{}, nil
}
