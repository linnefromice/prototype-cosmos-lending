package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/linnefromice/lending/x/lending/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PairPoolAll(c context.Context, req *types.QueryAllPairPoolRequest) (*types.QueryAllPairPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var pairPools []types.PairPool
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	pairPoolStore := prefix.NewStore(store, types.KeyPrefix(types.PairPoolKey))

	pageRes, err := query.Paginate(pairPoolStore, req.Pagination, func(key []byte, value []byte) error {
		var pairPool types.PairPool
		if err := k.cdc.Unmarshal(value, &pairPool); err != nil {
			return err
		}

		pairPools = append(pairPools, pairPool)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPairPoolResponse{PairPool: pairPools, Pagination: pageRes}, nil
}

func (k Keeper) PairPool(c context.Context, req *types.QueryGetPairPoolRequest) (*types.QueryGetPairPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	pairPool, found := k.GetPairPool(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetPairPoolResponse{PairPool: pairPool}, nil
}
