package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/linnefromice/lending/x/lending/types"
)

func (k msgServer) Borrow(goCtx context.Context, msg *types.MsgBorrow) (*types.MsgBorrowResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.Keeper.BorrowFromPairPool(ctx, msg)
	if err != nil {
		panic(err)
	}

	return &types.MsgBorrowResponse{}, nil
}
