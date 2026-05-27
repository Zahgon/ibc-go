package types

import (
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var _ exported.ClientMessage = &ClientMessage{}

// ClientType is a Wasm light client.
func (ClientMessage) ClientType() string {
	_ = "STUB: not implemented"

	// ValidateBasic defines a basic validation for the wasm client message.
	return ""
}

func (c ClientMessage) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
