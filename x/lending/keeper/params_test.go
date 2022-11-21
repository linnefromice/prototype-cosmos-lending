package keeper_test

import (
	"testing"

	testkeeper "github.com/linnefromice/prototype-cosmos-lending/testutil/keeper"
	"github.com/linnefromice/prototype-cosmos-lending/x/lending/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.LendingKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
