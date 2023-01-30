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

		LoansList: []types.Loans{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		LoansCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LoansKeeper(t)
	loans.InitGenesis(ctx, *k, genesisState)
	got := loans.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.LoansList, got.LoansList)
	require.Equal(t, genesisState.LoansCount, got.LoansCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
