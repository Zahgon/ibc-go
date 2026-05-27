package types

import (
	commitmenttypes "github.com/cosmos/ibc-go/v11/modules/core/23-commitment/types"
)

// MaxMerklePrefixLength defines the maximum length of the counterparty prefix in bytes. (This is an arbitrarily chosen value)
const MaxMerklePrefixLength = 256

// NewConnectionEnd creates a new ConnectionEnd instance.
func NewConnectionEnd(state State, clientID string, counterparty Counterparty, versions []*Version, delayPeriod uint64) ConnectionEnd {
	_ = "STUB: not implemented"
	return *new(ConnectionEnd)
}

// ValidateBasic implements the Connection interface.
// NOTE: the protocol supports that the connection and client IDs match the
// counterparty's.
func (c ConnectionEnd) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewCounterparty creates a new Counterparty instance.
func NewCounterparty(clientID, connectionID string, prefix commitmenttypes.MerklePrefix) Counterparty {
	_ = "STUB: not implemented"
	return *new(Counterparty)
}

// ValidateBasic performs a basic validation check of the identifiers and prefix
func (c Counterparty) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewIdentifiedConnection creates a new IdentifiedConnection instance
func NewIdentifiedConnection(connectionID string, conn ConnectionEnd) IdentifiedConnection {
	_ = "STUB: not implemented"
	return *new(IdentifiedConnection)
}

// ValidateBasic performs a basic validation of the connection identifier and connection fields.
func (ic IdentifiedConnection) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
