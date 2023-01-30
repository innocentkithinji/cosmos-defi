package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/loans module sentinel errors
var (
	ErrDeadline       = sdkerrors.Register(ModuleName, 3, "deadline")
	ErrWrongLoanState = sdkerrors.Register(ModuleName, 2, "wrong loan state")
)
