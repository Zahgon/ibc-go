package v7

import (
	"github.com/cosmos/cosmos-sdk/codec"

	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
)

// MigrateGenesis accepts an exported IBC client genesis file and migrates it to:
//
// - Update solo machine client state protobuf definition (v2 to v3)
// - Remove all solo machine consensus states
// - Remove localhost client
func MigrateGenesis(clientGenState *clienttypes.GenesisState, cdc codec.ProtoCodecMarshaler) (*clienttypes.GenesisState, error) {
	_ = "STUB: not implemented"
	// To prune the client and consensus states, we will create new slices to fill up
	// with information we want to keep.
	return nil, nil
}

// remove localhost client state by not adding client state

// add all other client states

// iterate consensus states by client

// look for consensus states for the current client

// remove all consensus states for the solo machine and localhost
// do not add to new clientsConsensus

// ensure all consensus states added for other client types
