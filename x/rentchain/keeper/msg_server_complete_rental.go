package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gopherine/rentchain/x/rentchain/types"
)

func (k msgServer) CompleteRental(goCtx context.Context, msg *types.MsgCompleteRental) (*types.MsgCompleteRentalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	index := GenerateIndex(msg.ItemId, msg.OwnerId)
	// Retrieve the rental agreement
	agreement, found := k.GetRentalAgreement(ctx, index)
	if !found {
		return nil, fmt.Errorf("no rental found for item %s", msg.ItemId)
	}
	fmt.Println("Before :", agreement)

	// Check if the rental is already completed
	if !agreement.IsActive {
		return nil, fmt.Errorf("rental for item %s is already completed", msg.ItemId)
	}

	// Mark the rental as completed
	agreement.IsActive = false

	k.SetRentalAgreement(ctx, agreement)

	agreement, found = k.GetRentalAgreement(ctx, index)
	if !found {
		return nil, fmt.Errorf("no rental found for item %s", msg.ItemId)
	}

	fmt.Println("Updated :", agreement)

	// TODO: events not appearing on the logs
	event := types.Event{
		Type: "rental_completed",
		Attributes: []types.EventAttribute{
			{Key: "item_id", Value: msg.ItemId},
			{Key: "owner_id", Value: msg.OwnerId},
		},
	}

	// Emit the event
	k.EmitStructuredEvent(ctx, event)
	return &types.MsgCompleteRentalResponse{}, nil
}
