package loans_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "loans/testutil/keeper"
	"loans/testutil/nullify"
	"loans/x/loans"
	"loans/x/loans/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LoansKeeper(t)
	loans.InitGenesis(ctx, *k, genesisState)
	got := loans.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
