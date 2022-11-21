package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/linnefromice/lending/x/lending/types"
)

func (k msgServer) AddPool(goCtx context.Context, msg *types.MsgAddPool) (*types.MsgAddPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAddPoolResponse{}, nil
}
