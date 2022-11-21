package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/linnefromice/lending/x/lending/types"
)

func (k Keeper) AddPool(ctx sdk.Context, msg *types.MsgAddPool) (types.PairPool, error) {
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	poolId := k.GetPairPoolCount(ctx) + 1

	// validations

	pool := types.PairPool{
		Address:                    creator.String(), // temp
		PoolId:                     poolId,
		AssetLiquidity:             msg.Amount,            // temp
		AssetLpCoinDenom:           msg.Amount.GetDenom(), // temp
		AssetTotalNormalDeposited:  0,
		AssetTotalConlyDeposited:   0,
		AssetTotalBorrowed:         0,
		ShadowLiquidity:            msg.Amount,            // temp
		ShadowLpCoinDenom:          msg.Amount.GetDenom(), // temp
		ShadowTotalNormalDeposited: 0,
		ShadowTotalConlyDeposited:  0,
		ShadowTotalBorrowed:        0,
		LastUpdated:                0, // temp
	}

	k.AppendPairPool(ctx, pool)

	return pool, nil
}
