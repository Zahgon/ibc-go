package clientv2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/core/02-client/v2/keeper"
	"github.com/cosmos/ibc-go/v11/modules/core/02-client/v2/types"
)

// InitGenesis initializes the ibc client/v2 submodule's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k *keeper.Keeper, gs types.GenesisState) {
	_ = "STUB: not implemented"
	return
}

// ExportGenesis returns the ibc client/v2 submodule's exported genesis.
func ExportGenesis(ctx sdk.Context, k *keeper.Keeper) types.GenesisState {
	_ = "STUB: not implemented"
	return *new(types.GenesisState)
}
