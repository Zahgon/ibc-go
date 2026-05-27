package keeper

import (
	"context"

	"github.com/cosmos/ibc-go/v11/modules/apps/27-gmp/types"
)

// InitGenesis initializes the module state from a genesis state.
func (k *Keeper) InitGenesis(ctx context.Context, data *types.GenesisState) error {
	_ = "STUB: not implemented"
	return nil
}

// ExportGenesis exports the module state to a genesis state.
func (k *Keeper) ExportGenesis(ctx context.Context) (*types.GenesisState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
