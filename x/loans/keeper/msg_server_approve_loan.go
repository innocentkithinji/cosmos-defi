package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"loans/x/loans/types"
)

func (k msgServer) ApproveLoan(goCtx context.Context, msg *types.MsgApproveLoan) (*types.MsgApproveLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgApproveLoanResponse{}, nil
}
