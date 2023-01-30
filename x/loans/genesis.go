package loans

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"loans/x/loans/keeper"
	"loans/x/loans/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the loans
	for _, elem := range genState.LoansList {
		k.SetLoans(ctx, elem)
	}

	// Set loans count
	k.SetLoansCount(ctx, genState.LoansCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.LoansList = k.GetAllLoans(ctx)
	genesis.LoansCount = k.GetLoansCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
