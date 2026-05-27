package attestations

import (
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var _ exported.ConsensusState = (*ConsensusState)(nil)

// ClientType returns Attestations type.
func (ConsensusState) ClientType() string { _ = "STUB: not implemented"; return "" }

// GetTimestamp is deprecated and will panic if called.
func (ConsensusState) GetTimestamp() uint64 { _ = "STUB: not implemented"; return 0 }

// ValidateBasic defines basic validation for the attestations consensus state.
func (cs ConsensusState) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
