package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/linnefromice/lending/x/lending/types"
)

// GetPairPoolCount get the total number of pairPool
func (k Keeper) GetPairPoolCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PairPoolCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetPairPoolCount set the total number of pairPool
func (k Keeper) SetPairPoolCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PairPoolCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendPairPool appends a pairPool in the store with a new id and update the count
func (k Keeper) AppendPairPool(
	ctx sdk.Context,
	pairPool types.PairPool,
) uint64 {
	// Create the pairPool
	count := k.GetPairPoolCount(ctx)

	// Set the ID of the appended value
	pairPool.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PairPoolKey))
	appendedValue := k.cdc.MustMarshal(&pairPool)
	store.Set(GetPairPoolIDBytes(pairPool.Id), appendedValue)

	// Update pairPool count
	k.SetPairPoolCount(ctx, count+1)

	return count
}

// SetPairPool set a specific pairPool in the store
func (k Keeper) SetPairPool(ctx sdk.Context, pairPool types.PairPool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PairPoolKey))
	b := k.cdc.MustMarshal(&pairPool)
	store.Set(GetPairPoolIDBytes(pairPool.Id), b)
}

// GetPairPool returns a pairPool from its id
func (k Keeper) GetPairPool(ctx sdk.Context, id uint64) (val types.PairPool, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PairPoolKey))
	b := store.Get(GetPairPoolIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePairPool removes a pairPool from the store
func (k Keeper) RemovePairPool(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PairPoolKey))
	store.Delete(GetPairPoolIDBytes(id))
}

// GetAllPairPool returns all pairPool
func (k Keeper) GetAllPairPool(ctx sdk.Context) (list []types.PairPool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PairPoolKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PairPool
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetPairPoolIDBytes returns the byte representation of the ID
func GetPairPoolIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetPairPoolIDFromBytes returns ID in uint64 format from a byte array
func GetPairPoolIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
