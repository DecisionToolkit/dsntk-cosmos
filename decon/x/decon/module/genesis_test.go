package decon_test

import (
	"testing"

	keepertest "decon/testutil/keeper"
	"decon/testutil/nullify"
	"decon/x/decon/module"
	"decon/x/decon/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DeconKeeper(t)
	decon.InitGenesis(ctx, k, genesisState)
	got := decon.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
