package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "blog/testutil/keeper"
	"blog/x/blog/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.BlogKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
