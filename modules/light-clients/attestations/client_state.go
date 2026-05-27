package attestations

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/v2/types"

	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var _ exported.ClientState = (*ClientState)(nil)

var nonMembershipCommitment = make([]byte, 32)

// NewClientState creates a new ClientState instance.
func NewClientState(attestorAddresses []string, minRequiredSigs uint32, latestHeight uint64) *ClientState {
	_ = "STUB: not implemented"
	return nil
}

// ClientType is Attestations.
func (ClientState) ClientType() string { _ = "STUB: not implemented"; return "" }

// Validate performs basic validation of the client state fields.
func (cs ClientState) Validate() error { _ = "STUB: not implemented"; return nil }

// verifyMembership is a generic proof verification method which verifies a proof of the existence of a value at a given CommitmentPath at the specified height.
func (cs *ClientState) verifyMembership(
	clientStore storetypes.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
	proof []byte,
	path exported.Path,
	value []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

// verifyNonMembership verifies a proof of the absence of a value at a given CommitmentPath at the specified height.
func (cs *ClientState) verifyNonMembership(
	clientStore storetypes.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
	proof []byte,
	path exported.Path,
) error {
	_ = "STUB: not implemented"
	return nil
}

// sets the client state to the store
func setClientState(store storetypes.KVStore, cdc codec.BinaryCodec, clientState exported.ClientState) {
	_ = "STUB: not implemented"
	return
}
