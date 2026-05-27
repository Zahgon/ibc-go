package ibc

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/core/keeper"
	"github.com/cosmos/ibc-go/v11/modules/core/types"
)

// InitGenesis initializes the ibc state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, gs *types.GenesisState) {
	_ = "STUB: not implemented"
	return
}

// ExportGenesis returns the ibc exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	_ = "STUB: not implemented"
	return nil
}
