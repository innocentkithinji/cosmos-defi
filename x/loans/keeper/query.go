package keeper

import (
	"loans/x/loans/types"
)

var _ types.QueryServer = Keeper{}
