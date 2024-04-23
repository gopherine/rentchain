package rentchain

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/gopherine/rentchain/api/rentchain/rentchain"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "RentalAgreementAll",
					Use:       "list-rental-agreement",
					Short:     "List all RentalAgreement",
				},
				{
					RpcMethod:      "RentalAgreement",
					Use:            "show-rental-agreement [id]",
					Short:          "Shows a RentalAgreement",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateRental",
					Use:            "create-rental [item-id] [owner-id] [renter-id] [price] [start-time] [duration]",
					Short:          "Send a CreateRental tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "itemId"}, {ProtoField: "ownerId"}, {ProtoField: "renterId"}, {ProtoField: "price"}, {ProtoField: "startTime"}, {ProtoField: "duration"}},
				},
				{
					RpcMethod:      "CompleteRental",
					Use:            "complete-rental [item-id] [owner-id]",
					Short:          "Send a CompleteRental tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "itemId"}, {ProtoField: "ownerId"}},
				},
				{
					RpcMethod:      "CreateRentalAgreement",
					Use:            "create-rental-agreement [index] [itemId] [ownerId] [renterId] [price] [startTime] [duration] [isActive]",
					Short:          "Create a new RentalAgreement",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "itemId"}, {ProtoField: "ownerId"}, {ProtoField: "renterId"}, {ProtoField: "price"}, {ProtoField: "startTime"}, {ProtoField: "duration"}, {ProtoField: "isActive"}},
				},
				{
					RpcMethod:      "UpdateRentalAgreement",
					Use:            "update-rental-agreement [index] [itemId] [ownerId] [renterId] [price] [startTime] [duration] [isActive]",
					Short:          "Update RentalAgreement",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "itemId"}, {ProtoField: "ownerId"}, {ProtoField: "renterId"}, {ProtoField: "price"}, {ProtoField: "startTime"}, {ProtoField: "duration"}, {ProtoField: "isActive"}},
				},
				{
					RpcMethod:      "DeleteRentalAgreement",
					Use:            "delete-rental-agreement [index]",
					Short:          "Delete RentalAgreement",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
