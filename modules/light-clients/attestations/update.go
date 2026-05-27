package attestations

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/v2/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

// VerifyClientMessage introspects the provided ClientMessage and checks its validity.
// An AttestationProof is considered valid if it has valid signatures from unique attestors meeting quorum.
func (cs *ClientState) VerifyClientMessage(ctx sdk.Context, cdc codec.BinaryCodec, clientStore storetypes.KVStore, clientMsg exported.ClientMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateState updates the consensus state to a new height and timestamp.
// A list containing the updated consensus height is returned.
// Since client message is validated in VerifyClientMessage, we don't validate much here, and panics on anything unexpected.
func (cs *ClientState) UpdateState(ctx sdk.Context, cdc codec.BinaryCodec, clientStore storetypes.KVStore, clientMsg exported.ClientMessage) []exported.Height {
	_ = "STUB: not implemented"
	return nil
}
