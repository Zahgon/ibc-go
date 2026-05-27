package solomachine

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var _ exported.ConsensusState = (*ConsensusState)(nil)

// ClientType returns Solo Machine type.
func (ConsensusState) ClientType() string { _ = "STUB: not implemented"; return "" }

// GetTimestamp returns zero.
func (cs ConsensusState) GetTimestamp() uint64 { _ = "STUB: not implemented"; return 0 }

// GetPubKey unmarshals the public key into a cryptotypes.PubKey type.
// An error is returned if the public key is nil or the cached value
// is not a PubKey.
func (cs ConsensusState) GetPubKey() (cryptotypes.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(cryptotypes.PubKey), nil
}

// ValidateBasic defines basic validation for the solo machine consensus state.
func (cs ConsensusState) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
