package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/gopherine/rentchain/x/assets/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AssetAll(ctx context.Context, req *types.QueryAllAssetRequest) (*types.QueryAllAssetResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var assets []types.Asset

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	assetStore := prefix.NewStore(store, types.KeyPrefix(types.AssetKey))

	pageRes, err := query.Paginate(assetStore, req.Pagination, func(key []byte, value []byte) error {
		var asset types.Asset
		if err := k.cdc.Unmarshal(value, &asset); err != nil {
			return err
		}

		assets = append(assets, asset)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAssetResponse{Asset: assets, Pagination: pageRes}, nil
}

func (k Keeper) Asset(ctx context.Context, req *types.QueryGetAssetRequest) (*types.QueryGetAssetResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	asset, found := k.GetAsset(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetAssetResponse{Asset: asset}, nil
}
