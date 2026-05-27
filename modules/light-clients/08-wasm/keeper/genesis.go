package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/modules/light-clients/08-wasm/v11/types"
)

// InitGenesis initializes the 08-wasm module's state from a provided genesis
// state.
func (k *Keeper) InitGenesis(ctx sdk.Context, gs types.GenesisState) error {
	_ = "STUB: not implemented"
	return nil
}

// ExportGenesis returns the 08-wasm module's exported genesis. This includes the code
// for all contracts previously stored.
func (k *Keeper) ExportGenesis(ctx sdk.Context) types.GenesisState {
	_ = "STUB: not implemented"
	return *new(types.GenesisState)
}

// Grab code from wasmVM and add to genesis state.
