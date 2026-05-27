package tendermint

import (
	"time"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	cmttypes "github.com/cometbft/cometbft/types"

	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var _ exported.ClientMessage = (*Misbehaviour)(nil)

// FrozenHeight is same for all misbehaviour
var FrozenHeight = clienttypes.NewHeight(0, 1)

// NewMisbehaviour creates a new Misbehaviour instance.
func NewMisbehaviour(clientID string, header1, header2 *Header) *Misbehaviour {
	_ = "STUB: not implemented"
	return nil
}

// ClientType is Tendermint light client
func (Misbehaviour) ClientType() string { _ = "STUB: not implemented"; return "" }

// GetTime returns the timestamp at which misbehaviour occurred. It uses the
// maximum value from both headers to prevent producing an invalid header outside
// of the misbehaviour age range.
func (m Misbehaviour) GetTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// ValidateBasic implements Misbehaviour interface
func (m Misbehaviour) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// ValidateBasic on both validators

// Ensure that Height1 is greater than or equal to Height2

// validCommit checks if the given commit is a valid commit from the passed-in validatorset
func validCommit(chainID string, blockID cmttypes.BlockID, commit *cmtproto.Commit, valSet *cmtproto.ValidatorSet) error {
	_ = "STUB: not implemented"
	return nil
}
