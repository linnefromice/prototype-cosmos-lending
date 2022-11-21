package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/linnefromice/prototype-cosmos-lending/testutil/keeper"
	"github.com/linnefromice/prototype-cosmos-lending/x/lending/keeper"
	"github.com/linnefromice/prototype-cosmos-lending/x/lending/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.LendingKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
