package solomachine

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/v2/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"

	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var _ exported.ClientState = (*ClientState)(nil)

// NewClientState creates a new ClientState instance.
func NewClientState(latestSequence uint64, consensusState *ConsensusState) *ClientState {
	_ = "STUB: not implemented"
	return nil
}

// ClientType is Solo Machine.
func (ClientState) ClientType() string { _ = "STUB: not implemented"; return "" }

// Validate performs basic validation of the client state fields.
func (cs ClientState) Validate() error { _ = "STUB: not implemented"; return nil }

// verifyMembership is a generic proof verification method which verifies a proof of the existence of a value at a given CommitmentPath at the latest sequence.
// The caller is expected to construct the full CommitmentPath from a CommitmentPrefix and a standardized path (as defined in ICS 24).
func (cs *ClientState) verifyMembership(
	clientStore storetypes.KVStore,
	cdc codec.BinaryCodec,
	proof []byte,
	path exported.Path,
	value []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

// in a multistore context: index 0 is the key for the IBC store in the multistore, index 1 is the key in the IBC store

// verifyNonMembership is a generic proof verification method which verifies the absence of a given CommitmentPath at the latest sequence.
// The caller is expected to construct the full CommitmentPath from a CommitmentPrefix and a standardized path (as defined in ICS 24).
func (cs *ClientState) verifyNonMembership(
	clientStore storetypes.KVStore,
	cdc codec.BinaryCodec,
	proof []byte,
	path exported.Path,
) error {
	_ = "STUB: not implemented"
	return nil
}

// in a multistore context: index 0 is the key for the IBC store in the multistore, index 1 is the key in the IBC store

// produceVerificationArgs performs the basic checks on the arguments that are
// shared between the verification functions and returns the public key of the
// consensus state, the unmarshalled proof representing the signature and timestamp.
func produceVerificationArgs(
	cdc codec.BinaryCodec,
	cs *ClientState,
	proof []byte,
) (cryptotypes.PubKey, signing.SignatureData, uint64, uint64, error) {
	_ = "STUB: not implemented"
	return *new(cryptotypes.PubKey), *new(signing.SignatureData), 0, 0, nil
}

// sets the client state to the store
func setClientState(store storetypes.KVStore, cdc codec.BinaryCodec, clientState exported.ClientState) {
	_ = "STUB: not implemented"
	return
}
