package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/gopherine/rentchain/testutil/keeper"
	"github.com/gopherine/rentchain/x/assets/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.AssetsKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
