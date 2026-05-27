package tendermint

import (
	"time"

	cmtbytes "github.com/cometbft/cometbft/libs/bytes"

	commitmenttypes "github.com/cosmos/ibc-go/v11/modules/core/23-commitment/types"
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var _ exported.ConsensusState = (*ConsensusState)(nil)

// SentinelRoot is used as a stand-in root value for the consensus state set at the upgrade height
const SentinelRoot = "sentinel_root"

// NewConsensusState creates a new ConsensusState instance.
func NewConsensusState(
	timestamp time.Time, root commitmenttypes.MerkleRoot, nextValsHash cmtbytes.HexBytes,
) *ConsensusState {
	_ = "STUB: not implemented"
	return nil
}

// ClientType returns Tendermint
func (ConsensusState) ClientType() string { _ = "STUB: not implemented"; return "" }

// GetRoot returns the commitment Root for the specific
func (cs ConsensusState) GetRoot() exported.Root {
	_ = "STUB: not implemented"

	// GetTimestamp returns block time in nanoseconds of the header that created consensus state
	return *new(exported.Root)
}

func (cs ConsensusState) GetTimestamp() uint64 { _ = "STUB: not implemented"; return 0 }

// ValidateBasic defines a basic validation for the tendermint consensus state.
// NOTE: ProcessedTimestamp may be zero if this is an initial consensus state passed in by relayer
// as opposed to a consensus state constructed by the chain.
func (cs ConsensusState) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
