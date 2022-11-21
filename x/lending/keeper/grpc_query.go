package keeper

import (
	"github.com/linnefromice/prototype-cosmos-lending/x/lending/types"
)

var _ types.QueryServer = Keeper{}
