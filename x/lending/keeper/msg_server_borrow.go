package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/linnefromice/lending/x/lending/types"
)

func (k msgServer) Borrow(goCtx context.Context, msg *types.MsgBorrow) (*types.MsgBorrowResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	res, err := k.Keeper.BorrowFromPairPool(ctx, msg)
	if err != nil {
		panic(err)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.PoolBorrowedEventType,
			sdk.NewAttribute(types.PoolEventId, fmt.Sprint(res.Id)),
			sdk.NewAttribute(types.PoolEventAmount, fmt.Sprint(res.Id))),
	)

	return &types.MsgBorrowResponse{}, nil
}
