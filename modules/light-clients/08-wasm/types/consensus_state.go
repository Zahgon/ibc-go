package types

import (
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var _ exported.ConsensusState = (*ConsensusState)(nil)

// NewConsensusState creates a new ConsensusState instance.
func NewConsensusState(data []byte) *ConsensusState { _ = "STUB: not implemented"; return nil }

// ClientType returns Wasm type.
func (ConsensusState) ClientType() string {
	_ = "STUB: not implemented"

	// GetTimestamp returns block time in nanoseconds of the header that created consensus state.
	return ""
}

func (ConsensusState) GetTimestamp() uint64 {
	_ = "STUB: not implemented"

	// ValidateBasic defines a basic validation for the wasm client consensus state.
	return 0
}

func (cs ConsensusState) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
