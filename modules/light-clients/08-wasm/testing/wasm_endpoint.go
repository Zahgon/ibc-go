package testing

import (
	ibctesting "github.com/cosmos/ibc-go/v11/testing"
)

// WasmEndpoint is a wrapper around the ibctesting pkg Endpoint struct.
// It will override any functions which require special handling for the wasm client.
type WasmEndpoint struct {
	*ibctesting.Endpoint
}

// NewWasmEndpoint returns a wasm endpoint with the default ibctesting pkg
// Endpoint embedded.
func NewWasmEndpoint(chain *ibctesting.TestChain) *WasmEndpoint {
	_ = "STUB: not implemented"
	return nil
}

// CreateClient creates an wasm client on a mock cometbft chain.
// The client and consensus states are represented by byte slices
// and the starting height is 1.
func (ep *WasmEndpoint) CreateClient() error { _ = "STUB: not implemented"; return nil }
