package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/gopherine/rentchain/x/rentchain/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RentalAgreementAll(ctx context.Context, req *types.QueryAllRentalAgreementRequest) (*types.QueryAllRentalAgreementResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var rentalAgreements []types.RentalAgreement

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	rentalAgreementStore := prefix.NewStore(store, types.KeyPrefix(types.RentalAgreementKeyPrefix))

	pageRes, err := query.Paginate(rentalAgreementStore, req.Pagination, func(key []byte, value []byte) error {
		var rentalAgreement types.RentalAgreement
		if err := k.cdc.Unmarshal(value, &rentalAgreement); err != nil {
			return err
		}

		rentalAgreements = append(rentalAgreements, rentalAgreement)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRentalAgreementResponse{RentalAgreement: rentalAgreements, Pagination: pageRes}, nil
}

func (k Keeper) RentalAgreement(ctx context.Context, req *types.QueryGetRentalAgreementRequest) (*types.QueryGetRentalAgreementResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetRentalAgreement(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetRentalAgreementResponse{RentalAgreement: val}, nil
}
