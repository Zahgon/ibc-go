package types

import (
	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var _ exported.ClientState = (*ClientState)(nil)

// NewClientState creates a new ClientState instance.
func NewClientState(data []byte, checksum []byte, height clienttypes.Height) *ClientState {
	_ = "STUB: not implemented"
	return nil
}

// ClientType is Wasm light client.
func (ClientState) ClientType() string {
	_ = "STUB: not implemented"

	// Validate performs a basic validation of the client state fields.
	return ""
}

func (cs ClientState) Validate() error { _ = "STUB: not implemented"; return nil }
