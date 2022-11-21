package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	// "github.com/cosmos/cosmos-sdk/types/address"
	"github.com/linnefromice/lending/x/lending/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// func newPairPoolAddress(poolId uint64) sdk.AccAddress {
// 	key := append([]byte("pairpool"), sdk.Uint64ToBigEndian(poolId)...)
// 	return address.Module(types.ModuleName, key)
// }

// func (k Keeper) createModuleAccount(ctx sdk.Context, addr sdk.AccAddress) error {
// 	account := k.accountKeeper.NewAccount(
// 		ctx,
// 		authtypes.NewModuleAccount(
// 			authtypes.NewBaseAccountWithAddress(addr),
// 			addr.String(),
// 		),
// 	)
// 	k.accountKeeper.SetAccount(ctx, account)
// 	return nil
// }

func (k Keeper) AddPool(ctx sdk.Context, msg *types.MsgAddPool) (types.PairPool, error) {
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	poolId := k.GetPairPoolCount(ctx) + 1

	// validations

	// execute
	// poolAccAddress := newPairPoolAddress(poolId)
	// k.createModuleAccount(ctx, poolAccAddress)

	pool := types.PairPool{
		Address:                    creator.String(), // temp
		PoolId:                     poolId,
		AssetLiquidity:             msg.Amount,
		AssetTotalNormalDeposited:  msg.Amount.Amount.Uint64(),
		AssetTotalConlyDeposited:   0,
		AssetTotalBorrowed:         0,
		ShadowLiquidity:            NewUsdz(0),
		ShadowTotalNormalDeposited: 0,
		ShadowTotalConlyDeposited:  0,
		ShadowTotalBorrowed:        0,
		LastUpdated:                0, // TODO: current timestamp
	}

	if msg.Amount.IsPositive() {

		moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
		// NOTE: create module account if not exist
		//   TODO: not work well (module account is not created?)
		if !k.accountKeeper.HasAccount(ctx, moduleAddr) {
			moduleAcc := authtypes.NewEmptyModuleAccount(types.ModuleName, authtypes.Minter)
			k.accountKeeper.SetModuleAccount(ctx, moduleAcc)
			moduleAddr = moduleAcc.GetAddress()
		}
		if err := k.bankKeeper.SendCoins(ctx, creator, moduleAddr, sdk.NewCoins(msg.Amount)); err != nil {
			return types.PairPool{}, err
		}

		lpCoin := sdk.NewCoin(pool.GetAssetLpCoinDenom(), msg.Amount.Amount)
		err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(lpCoin))
		if err != nil {
			return types.PairPool{}, err
		}
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, sdk.NewCoins(lpCoin))
		if err != nil {
			return types.PairPool{}, err
		}
	}

	k.AppendPairPool(ctx, pool)

	return pool, nil
}
