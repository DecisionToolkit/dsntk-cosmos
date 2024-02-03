package keeper

import (
	"decon/x/decon/types"
)

var _ types.QueryServer = Keeper{}
