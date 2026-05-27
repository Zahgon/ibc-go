package solomachine

import (
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var _ exported.ClientMessage = (*Misbehaviour)(nil)

// ClientType is a Solo Machine light client.
func (Misbehaviour) ClientType() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic implements Misbehaviour interface.
func (m Misbehaviour) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// misbehaviour signatures cannot be identical.

// message data signed cannot be identical if both paths are the same.

// ValidateBasic ensures that the signature and data fields are non-empty.
func (sd SignatureAndData) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
