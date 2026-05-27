package solomachine

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

// SentinelHeaderPath defines a placeholder path value used for headers in solomachine client updates
const SentinelHeaderPath = "solomachine:header"

var _ exported.ClientMessage = (*Header)(nil)

// ClientType defines that the Header is a Solo Machine.
func (Header) ClientType() string { _ = "STUB: not implemented"; return "" }

// GetPubKey unmarshals the new public key into a cryptotypes.PubKey type.
// An error is returned if the new public key is nil or the cached value
// is not a PubKey.
func (h Header) GetPubKey() (cryptotypes.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(cryptotypes.PubKey), nil
}

// ValidateBasic ensures that the timestamp, signature and public key have all
// been initialized.
func (h Header) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
