package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/linnefromice/lending/testutil/keeper"
	"github.com/linnefromice/lending/x/lending"
	"github.com/linnefromice/lending/x/lending/keeper"
	"github.com/linnefromice/lending/x/lending/types"
	"github.com/stretchr/testify/require"
)

const (
	alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
)

func setupMsgServerForPool(t *testing.T) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.LendingKeeper(t)
	lending.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}

func TestAddPool(t *testing.T) {
	msgServer, keeper, goCtx := setupMsgServerForPool(t)
	ctx := sdk.UnwrapSDKContext(goCtx)
	resposnse, err := msgServer.AddPool(goCtx, &types.MsgAddPool{
		Creator: alice,
		Amount:  sdk.NewCoin("TestUSD", sdk.NewInt(100*1000000)),
		Active:  true,
	})

	require.Nil(t, err)
	require.Equal(t, uint64(1), resposnse.PoolId)
	pools := keeper.GetAllPairPool(ctx)
	require.Equal(t, 1, len(pools))
	count := keeper.GetPairPoolCount(ctx)
	require.Equal(t, uint64(1), count)
}
