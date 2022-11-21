package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/linnefromice/lending/testutil/keeper"
	"github.com/linnefromice/lending/testutil/nullify"
	"github.com/linnefromice/lending/x/lending/keeper"
	"github.com/linnefromice/lending/x/lending/types"
	"github.com/stretchr/testify/require"
)

func createNPairPool(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PairPool {
	items := make([]types.PairPool, n)
	for i := range items {
		items[i].Id = keeper.AppendPairPool(ctx, items[i])
	}
	return items
}

func TestPairPoolGet(t *testing.T) {
	keeper, ctx := keepertest.LendingKeeper(t)
	items := createNPairPool(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetPairPool(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestPairPoolRemove(t *testing.T) {
	keeper, ctx := keepertest.LendingKeeper(t)
	items := createNPairPool(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePairPool(ctx, item.Id)
		_, found := keeper.GetPairPool(ctx, item.Id)
		require.False(t, found)
	}
}

func TestPairPoolGetAll(t *testing.T) {
	keeper, ctx := keepertest.LendingKeeper(t)
	items := createNPairPool(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPairPool(ctx)),
	)
}

func TestPairPoolCount(t *testing.T) {
	keeper, ctx := keepertest.LendingKeeper(t)
	items := createNPairPool(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetPairPoolCount(ctx))
}
