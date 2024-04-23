package keeper_test

import (
	"fmt"
	"testing"

	keepertest "github.com/gopherine/rentchain/testutil/keeper"
	keeper "github.com/gopherine/rentchain/x/rentchain/keeper"
	"github.com/gopherine/rentchain/x/rentchain/types"
	"github.com/stretchr/testify/require"
)

func TestCreateRental(t *testing.T) {
	k, ctx := keepertest.RentchainKeeper(t) // setupKeeper needs to be implemented to prepare the environment
	srv := keeper.NewMsgServerImpl(k)

	msg := types.MsgCreateRental{
		Creator:   "creator",
		ItemId:    "111",
		OwnerId:   "cosmos1...",
		RenterId:  "cosmos2...",
		Price:     "10token",
		StartTime: "123456789",
		Duration:  "3600",
	}

	index := GenerateIndex(msg.ItemId, msg.OwnerId)
	// Create a new RentalAgreement struct
	agreement := types.RentalAgreement{
		Creator:   msg.Creator,
		Index:     index,
		ItemId:    msg.ItemId,
		OwnerId:   msg.OwnerId,
		RenterId:  msg.RenterId,
		Price:     msg.Price,
		StartTime: msg.StartTime,
		Duration:  msg.Duration,
		IsActive:  true,
	}

	_, err := srv.CreateRental(ctx, &msg)
	require.NoError(t, err, "CreateRental should not error")

	msg2 := types.MsgCompleteRental{
		Creator: msg.Creator,
		ItemId:  msg.ItemId,
		OwnerId: msg.OwnerId,
	}
	_, err = srv.CompleteRental(ctx, &msg2)
	require.NoError(t, err, "CreateRental should not error")

	// Check if the agreement was saved
	agreement, found := k.GetRentalAgreement(ctx, index)
	require.True(t, found, "Rental agreement should be found")
	require.Equal(t, msg.ItemId, agreement.ItemId, "Item ID should match")
}

// GenerateIndex generates a unique index for a rental agreement
func GenerateIndex(itemId, ownerId string) string {
	return fmt.Sprintf("%s-%s", ownerId, itemId)
}
