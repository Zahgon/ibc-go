package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	genesistypes "github.com/cosmos/ibc-go/v11/modules/apps/27-interchain-accounts/genesis/types"
)

// InitGenesis initializes the interchain accounts host application state from a provided genesis state
func InitGenesis(ctx sdk.Context, keeper Keeper, state genesistypes.HostGenesisState) {
	_ = "STUB: not implemented"
	return
}

// ExportGenesis returns the interchain accounts host exported genesis
func ExportGenesis(ctx sdk.Context, keeper Keeper) genesistypes.HostGenesisState {
	_ = "STUB: not implemented"
	return *new(genesistypes.HostGenesisState)
}
