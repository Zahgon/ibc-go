package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/apps/packet-forward-middleware/types"
)

// InitGenesis
func (k *Keeper) InitGenesis(ctx sdk.Context, state types.GenesisState) {
	_ = "STUB: not implemented"
	// Initialize store refund path for forwarded packets in genesis state that have not yet been acked.
	return
}

// ExportGenesis
func (k *Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	_ = "STUB: not implemented"
	return nil
}
