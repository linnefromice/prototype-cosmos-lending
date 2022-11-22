package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	// "github.com/cosmos/cosmos-sdk/types/address"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

func (k Keeper) DepositToPairPool(ctx sdk.Context, msg *types.MsgDeposit) error {
	// sender, _ := sdk.AccAddressFromBech32(msg.Creator)
	pool, found := k.GetPairPool(ctx, msg.PoolId)
	if !found {
		return sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.PoolId))
	}

	// validation

	// execute
	// moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)

	// TODO: collect sended
	// amount := sdk.NewInt(int64(msg.Amount))
	// var coin sdk.Coin
	// if msg.IsShadow {
	// 	coin = sdk.NewCoin(pool.ShadowLiquidity.Denom, amount)
	// } else {
	// 	coin = sdk.NewCoin(pool.AssetLiquidity.Denom, amount)
	// }

	// if err := k.bankKeeper.SendCoins(ctx, sender, moduleAddr, sdk.NewCoins(coin)); err != nil {
	// 	return err
	// }

	if msg.IsShadow {
		// TODO: Send lp coin
		// lpCoin := sdk.NewCoin(pool.GetShadowLpCoinDenom(), amount)
		// err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(lpCoin))
		// if err != nil {
		// 	return err
		// }
		// err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, sdk.NewCoins(lpCoin))
		// if err != nil {
		// 	return err
		// }

		// Update stats
		if msg.IsConly {
			pool.ShadowTotalConlyDeposited += msg.Amount
		} else {
			pool.ShadowTotalNormalDeposited += msg.Amount
		}
	} else {
		// TODO: Send lp coin
		// lpCoin := sdk.NewCoin(pool.GetAssetLpCoinDenom(), amount)
		// err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(lpCoin))
		// if err != nil {
		// 	return err
		// }
		// err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, sdk.NewCoins(lpCoin))
		// if err != nil {
		// 	return err
		// }

		// Update stats
		if msg.IsConly {
			pool.AssetTotalConlyDeposited += msg.Amount
		} else {
			pool.AssetTotalNormalDeposited += msg.Amount
		}
	}

	// Update state
	k.SetPairPool(ctx, pool)

	return nil
}

func (k Keeper) WithdrawFromPairPool(ctx sdk.Context, msg *types.MsgWithdraw) error {
	sender, _ := sdk.AccAddressFromBech32(msg.Creator)

	pool, found := k.GetPairPool(ctx, msg.PoolId)
	if !found {
		return sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.PoolId))
	}

	// validation

	// execute
	amount := sdk.NewInt(int64(msg.Amount))
	var coin sdk.Coin
	if msg.IsShadow {
		coin = sdk.NewCoin(pool.ShadowLiquidity.Denom, amount)
	} else {
		coin = sdk.NewCoin(pool.AssetLiquidity.Denom, amount)
	}
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, sdk.NewCoins(coin))
	if err != nil {
		return err
	}

	// TODO: consider conly
	if msg.IsShadow {
		pool.ShadowTotalNormalDeposited -= msg.Amount
	} else {
		pool.AssetTotalNormalDeposited -= msg.Amount
	}

	// Update state
	k.SetPairPool(ctx, pool)

	return nil
}

func (k Keeper) BorrowFromPairPool(ctx sdk.Context, msg *types.MsgBorrow) error {
	sender, _ := sdk.AccAddressFromBech32(msg.Creator)

	pool, found := k.GetPairPool(ctx, msg.PoolId)
	if !found {
		return sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.PoolId))
	}

	// validation

	// execute
	amount := sdk.NewInt(int64(msg.Amount))
	var coin sdk.Coin
	if msg.IsShadow {
		coin = sdk.NewCoin(pool.ShadowLiquidity.Denom(), amount)
	} else {
		coin = sdk.NewCoin(pool.AssetLiquidity.Denom(), amount)
	}
	err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(coin))
	if err != nil {
		return err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, sdk.NewCoins(coin))
	if err != nil {
		return err
	}

	if msg.IsShadow {
		pool.ShadowTotalBorrowed += msg.Amount
	} else {
		pool.AssetTotalBorrowed += msg.Amount
	}

	// Update state
	k.SetPairPool(ctx, pool)

	return nil
}

func (k Keeper) RepayToPairPool(ctx sdk.Context, msg *types.MsgRepay) error {
	sender, _ := sdk.AccAddressFromBech32(msg.Creator)

	pool, found := k.GetPairPool(ctx, msg.PoolId)
	if !found {
		return sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.PoolId))
	}

	// validation

	// execute
	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	amount := sdk.NewInt(int64(msg.Amount))
	var coin sdk.Coin
	if msg.IsShadow {
		coin = sdk.NewCoin(pool.ShadowLiquidity.Denom(), amount)
	} else {
		coin = sdk.NewCoin(pool.AssetLiquidity.Denom(), amount)
	}
	if err := k.bankKeeper.SendCoins(ctx, sender, moduleAddr, sdk.NewCoins(coin)); err != nil {
		return err
	}

	if msg.IsShadow {
		pool.ShadowTotalBorrowed -= msg.Amount
	} else {
		pool.AssetTotalBorrowed -= msg.Amount
	}

	// Update state
	k.SetPairPool(ctx, pool)

	return nil
}
