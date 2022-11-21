package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	UsdzDenom = "USDZ"
)

func NewUsdz(amount int64) sdk.Coin {
	return sdk.NewCoin(UsdzDenom, sdk.NewInt(amount))
}
