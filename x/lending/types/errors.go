package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/lending module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 9999, "sample error")
)
