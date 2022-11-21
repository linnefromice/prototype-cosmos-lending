package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/linnefromice/lending/testutil/keeper"
	"github.com/linnefromice/lending/x/lending/types"
	"github.com/stretchr/testify/require"
)

func TestPoolServiceAddPool(t *testing.T) {
	_, keeper, goCtx := keepertest.SetupMsgServerForPool(t)
	ctx := sdk.UnwrapSDKContext(goCtx)
	pool, err := keeper.AddPool(ctx, &types.MsgAddPool{
		Creator: alice,
		Amount:  sdk.NewCoin("TestUSD", sdk.NewInt(100*1000000)),
		Active:  true,
	})

	require.Nil(t, err)
	require.Equal(t, &types.PairPool{
		Address:                    alice,
		PoolId:                     1,
		AssetLiquidity:             sdk.NewCoin("TestUSD", sdk.NewInt(100*1000000)),
		AssetLpCoinDenom:           "TestUSD",
		AssetTotalNormalDeposited:  0,
		AssetTotalConlyDeposited:   0,
		AssetTotalBorrowed:         0,
		ShadowLiquidity:            sdk.NewCoin("TestUSD", sdk.NewInt(100*1000000)),
		ShadowLpCoinDenom:          "TestUSD",
		ShadowTotalNormalDeposited: 0,
		ShadowTotalConlyDeposited:  0,
		ShadowTotalBorrowed:        0,
		LastUpdated:                0,
	}, &pool)
}
