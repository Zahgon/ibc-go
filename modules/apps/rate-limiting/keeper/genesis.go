package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/apps/rate-limiting/types"
)

// InitGenesis initializes the rate-limiting module's state from a provided genesis state.
func (k *Keeper) InitGenesis(ctx sdk.Context, state types.GenesisState) {
	_ = "STUB: not implemented"
	// Set rate limits, blacklists, and whitelists
	return
}

// Set pending sequence numbers - validating that they're in right format of {channelId}/{sequenceNumber}

// If the hour epoch has been initialized already (epoch number != 0), validate and then use it

// If the hour epoch has not been initialized yet, set it so that the epoch number matches
// the current hour and the start time is precisely on the hour
//nolint:gosec

// ExportGenesis returns the rate-limiting module's exported genesis.
func (k *Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	_ = "STUB: not implemented"
	return nil
}
