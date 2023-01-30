package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "loans/testutil/keeper"
	"loans/x/loans/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.LoansKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
