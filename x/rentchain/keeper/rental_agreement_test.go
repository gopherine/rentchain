package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "github.com/gopherine/rentchain/testutil/keeper"
	"github.com/gopherine/rentchain/testutil/nullify"
	"github.com/gopherine/rentchain/x/rentchain/keeper"
	"github.com/gopherine/rentchain/x/rentchain/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNRentalAgreement(keeper keeper.Keeper, ctx context.Context, n int) []types.RentalAgreement {
	items := make([]types.RentalAgreement, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetRentalAgreement(ctx, items[i])
	}
	return items
}

func TestRentalAgreementGet(t *testing.T) {
	keeper, ctx := keepertest.RentchainKeeper(t)
	items := createNRentalAgreement(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetRentalAgreement(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestRentalAgreementRemove(t *testing.T) {
	keeper, ctx := keepertest.RentchainKeeper(t)
	items := createNRentalAgreement(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRentalAgreement(ctx,
			item.Index,
		)
		_, found := keeper.GetRentalAgreement(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestRentalAgreementGetAll(t *testing.T) {
	keeper, ctx := keepertest.RentchainKeeper(t)
	items := createNRentalAgreement(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRentalAgreement(ctx)),
	)
}
