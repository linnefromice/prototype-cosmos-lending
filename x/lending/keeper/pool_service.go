package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/linnefromice/lending/x/lending/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func newPairPoolAddress(poolId uint64) sdk.AccAddress {
	key := append([]byte("pairpool"), sdk.Uint64ToBigEndian(poolId)...)
	return address.Module(types.ModuleName, key)
}

func (k Keeper) createModuleAccount(ctx sdk.Context, addr sdk.AccAddress) error {
	account := k.accountKeeper.NewAccount(
		ctx,
		authtypes.NewModuleAccount(
			authtypes.NewBaseAccountWithAddress(addr),
			addr.String(),
		),
	)
	k.accountKeeper.SetAccount(ctx, account)
	return nil
}

func (k Keeper) AddPool(ctx sdk.Context, msg *types.MsgAddPool) (types.PairPool, error) {
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	poolId := k.GetPairPoolCount(ctx) + 1

	// validations

	// execute
	poolAccAddress := newPairPoolAddress(poolId)
	k.createModuleAccount(ctx, poolAccAddress)

	pool := types.PairPool{
		Address:                    poolAccAddress.String(),
		PoolId:                     poolId,
		AssetLiquidity:             msg.Amount,
		AssetLpCoinDenom:           types.GetPairPoolAssetLpCoinDenom(msg.Amount.Denom),
		AssetTotalNormalDeposited:  msg.Amount.Amount.Uint64(),
		AssetTotalConlyDeposited:   0,
		AssetTotalBorrowed:         0,
		ShadowLiquidity:            NewUsdz(0),
		ShadowLpCoinDenom:          types.GetPairPoolShadowLpCoinDenom(msg.Amount.Denom, UsdzDenom),
		ShadowTotalNormalDeposited: 0,
		ShadowTotalConlyDeposited:  0,
		ShadowTotalBorrowed:        0,
		LastUpdated:                0, // TODO: current timestamp
	}

	if msg.Amount.IsPositive() {
		if err := k.bankKeeper.SendCoins(ctx, creator, poolAccAddress, sdk.NewCoins(msg.Amount)); err != nil {
			return types.PairPool{}, err
		}
	}

	k.AppendPairPool(ctx, pool)

	return pool, nil
}
