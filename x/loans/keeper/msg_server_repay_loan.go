package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"loans/x/loans/types"
)

func (k msgServer) RepayLoan(goCtx context.Context, msg *types.MsgRepayLoan) (*types.MsgRepayLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRepayLoanResponse{}, nil
}
