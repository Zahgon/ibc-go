package types

import (
	wasmvmtypes "github.com/CosmWasm/wasmvm/v3/types"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/v2/types"
)

// GetClientState retrieves the client state from the store using the provided KVStore and codec.
// It returns the unmarshaled ClientState and a boolean indicating if the state was found.
func GetClientState(store storetypes.KVStore, cdc codec.BinaryCodec) (*ClientState, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Checksum is a type alias used for wasm byte code checksums.
type Checksum = wasmvmtypes.Checksum

// CreateChecksum creates a sha256 checksum from the given wasm code, it forwards the
// call to the wasmvm package. The code is checked for the following conditions:
// - code length is zero.
// - code length is less than 4 bytes (magic number length).
// - code does not start with the wasm magic number.
func CreateChecksum(code []byte) (Checksum, error) {
	_ = "STUB: not implemented"
	return *new(Checksum), nil
}
