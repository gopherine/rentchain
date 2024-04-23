package rentchain_test

import (
	"testing"

	keepertest "github.com/gopherine/rentchain/testutil/keeper"
	"github.com/gopherine/rentchain/testutil/nullify"
	rentchain "github.com/gopherine/rentchain/x/rentchain/module"
	"github.com/gopherine/rentchain/x/rentchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		RentalAgreementList: []types.RentalAgreement{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RentchainKeeper(t)
	rentchain.InitGenesis(ctx, k, genesisState)
	got := rentchain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.RentalAgreementList, got.RentalAgreementList)
	// this line is used by starport scaffolding # genesis/test/assert
}
