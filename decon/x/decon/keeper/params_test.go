package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "decon/testutil/keeper"
	"decon/x/decon/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.DeconKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
