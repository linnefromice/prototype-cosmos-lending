package keeper

import (
	"github.com/linnefromice/lending/x/lending/types"
)

var _ types.QueryServer = Keeper{}
