package solomachine

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/v2/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

// VerifyClientMessage introspects the provided ClientMessage and checks its validity
// A Solomachine Header is considered valid if the currently registered public key has signed over the new public key with the correct sequence
// A Solomachine Misbehaviour is considered valid if duplicate signatures of the current public key are found on two different messages at a given sequence
func (cs *ClientState) VerifyClientMessage(ctx sdk.Context, cdc codec.BinaryCodec, clientStore storetypes.KVStore, clientMsg exported.ClientMessage) error {
	_ = "STUB: not implemented"
	return nil
}

func (cs *ClientState) verifyHeader(cdc codec.BinaryCodec, header *Header) error {
	_ = "STUB: not implemented"
	// assert update timestamp is not less than current consensus state timestamp
	return nil
}

// assert currently registered public key signed over the new public key with correct sequence

// UpdateState updates the consensus state to the new public key and an incremented sequence.
// A list containing the updated consensus height is returned.
// If the provided clientMsg is not of type Header, the handler will no-op and return an empty slice.
func (cs *ClientState) UpdateState(ctx sdk.Context, cdc codec.BinaryCodec, clientStore storetypes.KVStore, clientMsg exported.ClientMessage) []exported.Height {
	_ = "STUB: not implemented"
	return nil
}

// clientMsg is invalid Misbehaviour, no update necessary

// create new solomachine ConsensusState
