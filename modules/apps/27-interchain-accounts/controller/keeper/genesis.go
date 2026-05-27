package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	genesistypes "github.com/cosmos/ibc-go/v11/modules/apps/27-interchain-accounts/genesis/types"
)

// InitGenesis initializes the interchain accounts controller application state from a provided genesis state
func InitGenesis(ctx sdk.Context, keeper Keeper, state genesistypes.ControllerGenesisState) {
	_ = "STUB: not implemented"
	return
}

// ExportGenesis returns the interchain accounts controller exported genesis
func ExportGenesis(ctx sdk.Context, keeper Keeper) genesistypes.ControllerGenesisState {
	_ = "STUB: not implemented"
	return *new(genesistypes.ControllerGenesisState)
}
