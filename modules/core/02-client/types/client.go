package types

import (
	"sort"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var (
	_ codectypes.UnpackInterfacesMessage = (*IdentifiedClientState)(nil)
	_ codectypes.UnpackInterfacesMessage = (*ConsensusStateWithHeight)(nil)
)

// NewIdentifiedClientState creates a new IdentifiedClientState instance
func NewIdentifiedClientState(clientID string, clientState exported.ClientState) IdentifiedClientState {
	_ = "STUB: not implemented"
	return *new(IdentifiedClientState)
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (ics IdentifiedClientState) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

var _ sort.Interface = (*IdentifiedClientStates)(nil)

// IdentifiedClientStates defines a slice of ClientConsensusStates that supports the sort interface
type IdentifiedClientStates []IdentifiedClientState

// Len implements sort.Interface
func (ics IdentifiedClientStates) Len() int {
	_ = "STUB: not implemented"

	// Less implements sort.Interface
	return 0
}

func (ics IdentifiedClientStates) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Swap implements sort.Interface
func (ics IdentifiedClientStates) Swap(i, j int) { _ = "STUB: not implemented"; return }

// Sort is a helper function to sort the set of IdentifiedClientStates in place
func (ics IdentifiedClientStates) Sort() IdentifiedClientStates {
	_ = "STUB: not implemented"
	return *new(IdentifiedClientStates)
}

// NewConsensusStateWithHeight creates a new ConsensusStateWithHeight instance
func NewConsensusStateWithHeight(height Height, consensusState exported.ConsensusState) ConsensusStateWithHeight {
	_ = "STUB: not implemented"
	return *new(ConsensusStateWithHeight)
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (cswh ConsensusStateWithHeight) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateClientType validates the client type. It cannot be blank or empty. It must be a valid
// client identifier when used with '0' or the maximum uint64 as the sequence.
func ValidateClientType(clientType string) error { _ = "STUB: not implemented"; return nil }

// IsValidClientID will check client type format and if the sequence is a uint64
