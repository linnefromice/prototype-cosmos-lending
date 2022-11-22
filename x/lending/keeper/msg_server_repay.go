package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/linnefromice/lending/x/lending/types"
)

func (k msgServer) Repay(goCtx context.Context, msg *types.MsgRepay) (*types.MsgRepayResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRepayResponse{}, nil
}
