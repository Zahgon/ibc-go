package solomachine

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/v2/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

// CheckForMisbehaviour returns true for type Misbehaviour (passed VerifyClientMessage check), otherwise returns false
func (ClientState) CheckForMisbehaviour(_ sdk.Context, _ codec.BinaryCodec, _ storetypes.KVStore, clientMsg exported.ClientMessage) bool {
	_ = "STUB: not implemented"
	return false
}

func (cs ClientState) verifyMisbehaviour(cdc codec.BinaryCodec, misbehaviour *Misbehaviour) error {
	_ = "STUB: not implemented"
	// NOTE: a check that the misbehaviour message data are not equal is done by
	// misbehaviour.ValidateBasic which is called by the 02-client keeper.
	// verify first signature
	return nil
}

// verify second signature

// verifySignatureAndData verifies that the currently registered public key has signed
// over the provided data and that the data is valid. The data is valid if it can be
// unmarshaled into the specified data type.
func (cs ClientState) verifySignatureAndData(cdc codec.BinaryCodec, misbehaviour *Misbehaviour, sigAndData *SignatureAndData) error {
	_ = "STUB: not implemented"
	// do not check misbehaviour timestamp since we want to allow processing of past misbehaviour
	return nil
}
