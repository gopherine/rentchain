package assets_test

import (
	"testing"

	keepertest "github.com/gopherine/rentchain/testutil/keeper"
	"github.com/gopherine/rentchain/testutil/nullify"
	assets "github.com/gopherine/rentchain/x/assets/module"
	"github.com/gopherine/rentchain/x/assets/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AssetsKeeper(t)
	assets.InitGenesis(ctx, k, genesisState)
	got := assets.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
