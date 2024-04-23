package keeper

import (
	"context"
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gopherine/rentchain/x/rentchain/types"
)

// GenerateIndex generates a unique index for a rental agreement
func GenerateIndex(itemId, ownerId string) string {
	return fmt.Sprintf("%s-%s", ownerId, itemId)
}

func (k msgServer) CreateRental(goCtx context.Context, msg *types.MsgCreateRental) (*types.MsgCreateRentalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	index := GenerateIndex(msg.ItemId, msg.OwnerId)

	// Logging input for debugging
	fmt.Printf("Received CreateRental: %+v\n", msg)

	// Check if the rental already exists
	if _, found := k.GetRentalAgreement(ctx, index); found {
		return nil, fmt.Errorf("item %s is already rented", msg.ItemId)
	}

	// Validate inputs
	if msg.OwnerId == "" || msg.RenterId == "" || msg.ItemId == "" {
		return nil, errors.New("owner, renter, and item ID must be provided")
	}

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

	k.SetRentalAgreement(ctx, agreement)

	// Check if the agreement is actually saved
	_, found := k.GetRentalAgreement(ctx, index)
	if !found {
		fmt.Println("Failed to save the rental agreement")
		return nil, fmt.Errorf("failed to save the rental agreement for item %s", msg.ItemId)
	}

	fmt.Println("Rental agreement created successfully")

	// TODO: events not appearing on the logs
	event := types.Event{
		Type: "rental_created",
		Attributes: []types.EventAttribute{
			{Key: "item_id", Value: msg.ItemId},
			{Key: "owner_id", Value: msg.OwnerId},
		},
	}

	// Emit the event
	k.EmitStructuredEvent(ctx, event)

	return &types.MsgCreateRentalResponse{}, nil
}
