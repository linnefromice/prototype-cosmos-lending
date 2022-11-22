package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/linnefromice/lending/x/lending/types"
)

func (k msgServer) Deposit(goCtx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	res, err := k.Keeper.DepositToPairPool(ctx, msg)
	if err != nil {
		panic(err)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.PoolDepositedEventType,
			sdk.NewAttribute(types.PoolEventId, fmt.Sprint(res.Id)),
			sdk.NewAttribute(types.PoolEventAmount, fmt.Sprint(res.Id))),
	)

	return &types.MsgDepositResponse{}, nil
}
