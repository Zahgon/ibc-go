//go:build cgo && !nolink_libwasmvm

package keeper

import (
	"cosmossdk.io/core/store"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/cosmos/ibc-go/modules/light-clients/08-wasm/v11/types"
)

// NewKeeperWithVM creates a new Keeper instance with the provided Wasm VM.
// This constructor function is meant to be used when the chain uses x/wasm
// and the same Wasm VM instance should be shared with it.
func NewKeeperWithVM(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	clientKeeper types.ClientKeeper,
	authority string,
	vm types.WasmEngine,
	queryRouter types.QueryRouter,
	opts ...Option,
) Keeper {
	_ = "STUB: not implemented"
	return *new(Keeper)
}

// set query plugins to ensure there is a non-nil query plugin
// regardless of what options the user provides

// NewKeeperWithConfig creates a new Keeper instance with the provided Wasm configuration.
// This constructor function is meant to be used when the chain does not use x/wasm
// and a Wasm VM needs to be instantiated using the provided parameters.
func NewKeeperWithConfig(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	clientKeeper types.ClientKeeper,
	authority string,
	wasmConfig types.WasmConfig,
	queryRouter types.QueryRouter,
	opts ...Option,
) Keeper {
	_ = "STUB: not implemented"
	return *new(Keeper)
}
