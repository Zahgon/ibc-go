package solomachine

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/v2/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

// CheckSubstituteAndUpdateState verifies that the subject is allowed to be updated by
// a governance proposal and that the substitute client is a solo machine.
// It will update the consensus state to the substitute's consensus state and
// the sequence to the substitute's current sequence. An error is returned if
// the client has been disallowed to be updated by a governance proposal,
// the substitute is not a solo machine, or the current public key equals
// the new public key.
func (cs *ClientState) CheckSubstituteAndUpdateState(
	ctx sdk.Context, cdc codec.BinaryCodec, subjectClientStore,
	_ storetypes.KVStore, substituteClient exported.ClientState,
) error {
	_ = "STUB: not implemented"
	return nil
}

// update to substitute parameters
