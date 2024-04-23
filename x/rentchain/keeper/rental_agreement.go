package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/gopherine/rentchain/x/rentchain/types"
)

// SetRentalAgreement set a specific rentalAgreement in the store from its index
func (k Keeper) SetRentalAgreement(ctx context.Context, rentalAgreement types.RentalAgreement) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RentalAgreementKeyPrefix))
	b := k.cdc.MustMarshal(&rentalAgreement)
	store.Set(types.RentalAgreementKey(
		rentalAgreement.Index,
	), b)
}

// GetRentalAgreement returns a rentalAgreement from its index
func (k Keeper) GetRentalAgreement(
	ctx context.Context,
	index string,

) (val types.RentalAgreement, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RentalAgreementKeyPrefix))

	b := store.Get(types.RentalAgreementKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRentalAgreement removes a rentalAgreement from the store
func (k Keeper) RemoveRentalAgreement(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RentalAgreementKeyPrefix))
	store.Delete(types.RentalAgreementKey(
		index,
	))
}

// GetAllRentalAgreement returns all rentalAgreement
func (k Keeper) GetAllRentalAgreement(ctx context.Context) (list []types.RentalAgreement) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RentalAgreementKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RentalAgreement
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
