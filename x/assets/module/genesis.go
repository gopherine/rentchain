package assets

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/gopherine/rentchain/x/assets/keeper"
	"github.com/gopherine/rentchain/x/assets/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the asset
	for _, elem := range genState.AssetList {
		k.SetAsset(ctx, elem)
	}

	// Set asset count
	k.SetAssetCount(ctx, genState.AssetCount)
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.AssetList = k.GetAllAsset(ctx)
	genesis.AssetCount = k.GetAssetCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
