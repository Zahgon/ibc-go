package keeper

import (
	snapshot "github.com/cosmos/cosmos-sdk/store/v2/snapshots/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/v2/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ snapshot.ExtensionSnapshotter = &WasmSnapshotter{}

// SnapshotFormat defines the default snapshot extension encoding format.
// SnapshotFormat 1 is gzipped wasm byte code for each item payload. No protobuf envelope, no metadata.
const SnapshotFormat = 1

// WasmSnapshotter implements the snapshot.ExtensionSnapshotter interface and is used to
// import and export state maintained within the wasmvm cache.
// NOTE: The following ExtensionSnapshotter has been adapted from CosmWasm's x/wasm:
// https://github.com/CosmWasm/wasmd/blob/v0.43.0/x/wasm/keeper/snapshotter.go
type WasmSnapshotter struct {
	cms    storetypes.MultiStore
	keeper *Keeper
}

// NewWasmSnapshotter creates and returns a new snapshot.ExtensionSnapshotter implementation for the 08-wasm module.
func NewWasmSnapshotter(cms storetypes.MultiStore, keeper *Keeper) snapshot.ExtensionSnapshotter {
	_ = "STUB: not implemented"
	return *new(snapshot.ExtensionSnapshotter)
}

// SnapshotName implements the snapshot.ExtensionSnapshotter interface.
// A unique name should be provided such that the implementation can be identified by the manager.
func (*WasmSnapshotter) SnapshotName() string { _ = "STUB: not implemented"; return "" }

// SnapshotFormat implements the snapshot.ExtensionSnapshotter interface.
// This is the default format used for encoding payloads when taking a snapshot.
func (*WasmSnapshotter) SnapshotFormat() uint32 { _ = "STUB: not implemented"; return 0 }

// SupportedFormats implements the snapshot.ExtensionSnapshotter interface.
// This defines a list of supported formats the snapshotter extension can restore from.
func (*WasmSnapshotter) SupportedFormats() []uint32 { _ = "STUB: not implemented"; return nil }

// SnapshotExtension implements the snapshot.ExntensionSnapshotter interface.
// SnapshotExtension is used to write data payloads into the underlying protobuf stream from the 08-wasm module.
func (ws *WasmSnapshotter) SnapshotExtension(height uint64, payloadWriter snapshot.ExtensionPayloadWriter) error {
	_ = "STUB: not implemented"
	return nil
}

// RestoreExtension implements the snapshot.ExtensionSnapshotter interface.
// RestoreExtension is used to read data from an existing extension state snapshot into the 08-wasm module.
// The payload reader returns io.EOF when it has reached the end of the extension state snapshot.
func (ws *WasmSnapshotter) RestoreExtension(height uint64, format uint32, payloadReader snapshot.ExtensionPayloadReader) error {
	_ = "STUB: not implemented"
	return nil
}

func restoreV1(ctx sdk.Context, k *Keeper, compressedCode []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (ws *WasmSnapshotter) processAllItems(
	height uint64,
	payloadReader snapshot.ExtensionPayloadReader,
	cb func(sdk.Context, *Keeper, []byte) error,
) error {
	_ = "STUB: not implemented"
	return nil
}
