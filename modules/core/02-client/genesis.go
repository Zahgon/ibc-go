package client

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/core/02-client/keeper"
	"github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
)

// InitGenesis initializes the ibc client submodule's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k *keeper.Keeper, gs types.GenesisState) {
	_ = "STUB: not implemented"
	return
}

// Set all client metadata first. This will allow client keeper to overwrite client and consensus state keys
// if clients accidentally write to ClientKeeper reserved keys.

// ExportGenesis returns the ibc client submodule's exported genesis.
// NOTE: the export process is not optimized, it will iterate three
// times over the 02-client sub-store.
func ExportGenesis(ctx sdk.Context, k *keeper.Keeper) types.GenesisState {
	_ = "STUB: not implemented"
	return *new(types.GenesisState)
}

// Warning: CreateLocalhost is deprecated
