package lending_test

import (
	"testing"

	keepertest "github.com/linnefromice/lending/testutil/keeper"
	"github.com/linnefromice/lending/testutil/nullify"
	"github.com/linnefromice/lending/x/lending"
	"github.com/linnefromice/lending/x/lending/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LendingKeeper(t)
	lending.InitGenesis(ctx, *k, genesisState)
	got := lending.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
