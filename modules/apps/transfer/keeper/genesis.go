package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/apps/transfer/types"
)

// InitGenesis initializes the ibc-transfer state and binds to PortID.
func (k *Keeper) InitGenesis(ctx sdk.Context, state types.GenesisState) {
	_ = "STUB: not implemented"
	return
}

// Every denom will have only one total escrow amount, since any
// duplicate entry will fail validation in Validate of GenesisState

// ExportGenesis exports ibc-transfer module's portID and denom trace info into its genesis state.
func (k *Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	_ = "STUB: not implemented"
	return nil
}
