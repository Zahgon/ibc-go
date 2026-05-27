package ibctesting

import (
	"testing"

	"github.com/cosmos/gogoproto/proto"

	"github.com/cosmos/cosmos-sdk/codec"

	abci "github.com/cometbft/cometbft/abci/types"
	cmttypes "github.com/cometbft/cometbft/types"
)

// ApplyValSetChanges takes in cmttypes.ValidatorSet and []abci.ValidatorUpdate and will return a new cmttypes.ValidatorSet which has the
// provided validator updates applied to the provided validator set.
func ApplyValSetChanges(tb testing.TB, valSet *cmttypes.ValidatorSet, valUpdates []abci.ValidatorUpdate) *cmttypes.ValidatorSet {
	_ = "STUB: not implemented"
	return nil
}

// must copy since validator set will mutate with UpdateWithChangeSet

// VoteAndCheckProposalStatus votes on a gov proposal, checks if the proposal has passed, and returns an error if it has not with the failure reason.
func VoteAndCheckProposalStatus(endpoint *Endpoint, proposalID uint64) error {
	_ = "STUB: not implemented"
	// vote on proposal
	return nil
}

// fast forward the chain context to end the voting period

// check if proposal passed or failed on msg execution
// we need to grab the context again since the previous context is no longer valid as the chain header time has been incremented

// GenerateString generates a random string of the given length in bytes
func GenerateString(length uint) string { _ = "STUB: not implemented"; return "" }

// UnmarshalMsgResponses parse out msg responses from a transaction result
func UnmarshalMsgResponses(cdc codec.Codec, data []byte, msgs ...proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// RequireErrorIsOrContains verifies that the passed error is either a target error or contains its error message.
func RequireErrorIsOrContains(t *testing.T, err, targetError error, msgAndArgs ...any) {
	_ = "STUB: not implemented"
	return
}
