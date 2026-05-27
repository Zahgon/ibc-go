package v7

import (
	"github.com/cosmos/cosmos-sdk/codec"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
)

// MigrateGenesis accepts an exported IBC client genesis file and migrates it to:
//
// - Update solo machine client state protobuf definition (v2 to v3)
// - Remove all solo machine consensus states
// - Remove any localhost clients
func MigrateGenesis(appState genutiltypes.AppMap, cdc codec.ProtoCodecMarshaler) (genutiltypes.AppMap, error) {
	_ = "STUB: not implemented"
	return *new(genutiltypes.AppMap), nil
}

// ensure legacy solo machines types are registered

// unmarshal old ibc genesis state

// delete old genesis state

// set new ibc genesis state
