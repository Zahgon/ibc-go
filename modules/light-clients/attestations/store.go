package attestations

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/v2/types"

	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

// getClientState retrieves the client state from the store using the provided KVStore and codec.
// It returns the unmarshaled ClientState and a boolean indicating if the state was found.
func getClientState(store storetypes.KVStore, cdc codec.BinaryCodec) (*ClientState, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// setConsensusState stores the consensus state at the given height.
func setConsensusState(clientStore storetypes.KVStore, cdc codec.BinaryCodec, consensusState *ConsensusState, height exported.Height) {
	_ = "STUB: not implemented"
	return
}

// getConsensusState retrieves the consensus state from the client prefixed store.
// If the ConsensusState does not exist in state for the provided height a nil value and false boolean flag is returned
func getConsensusState(store storetypes.KVStore, cdc codec.BinaryCodec, height exported.Height) (*ConsensusState, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
