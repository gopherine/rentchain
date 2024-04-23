package assets

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/gopherine/rentchain/api/rentchain/assets"
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
					RpcMethod: "AssetAll",
					Use:       "list-asset",
					Short:     "List all asset",
				},
				{
					RpcMethod:      "Asset",
					Use:            "show-asset [id]",
					Short:          "Shows a asset by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
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
					RpcMethod:      "CreateAsset",
					Use:            "create-asset [owner] [name] [description] [details] [pricePerUnit] [unit] [tags]",
					Short:          "Create asset",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "owner"}, {ProtoField: "name"}, {ProtoField: "description"}, {ProtoField: "details"}, {ProtoField: "pricePerUnit"}, {ProtoField: "unit"}, {ProtoField: "tags"}},
				},
				{
					RpcMethod:      "UpdateAsset",
					Use:            "update-asset [id] [name] [description] [details] [pricePerUnit] [unit] [tags]",
					Short:          "Update asset",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "name"}, {ProtoField: "description"}, {ProtoField: "details"}, {ProtoField: "pricePerUnit"}, {ProtoField: "unit"}, {ProtoField: "tags"}},
				},
				{
					RpcMethod:      "DeleteAsset",
					Use:            "delete-asset [id]",
					Short:          "Delete asset",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
