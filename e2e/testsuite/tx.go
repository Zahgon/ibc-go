package testsuite

import (
	"context"

	"github.com/cosmos/interchaintest/v11/ibc"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	govtypesv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	abci "github.com/cometbft/cometbft/abci/types"

	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
)

// BroadcastMessages broadcasts the provided messages to the given chain and signs them on behalf of the provided user.
// Once the broadcast response is returned, we wait for a few blocks to be created on the chain the message was broadcast to.
func (s *E2ETestSuite) BroadcastMessages(ctx context.Context, chain ibc.Chain, user ibc.Wallet, msgs ...sdk.Msg) sdk.TxResponse {
	_ = "STUB: not implemented"
	return *new(sdk.TxResponse)
}

// strip out any fields that may not be supported for the given chain version.

// use a codec with all the types our tests care about registered.
// BroadcastTx will deserialize the response and will not be able to otherwise.

// Retry the operation a few times if the user signing the transaction is a relayer. (See issue #3264)

// Retry five times, the value of 5 chosen is arbitrary.

// retryNtimes retries the provided function up to the provided number of attempts.
func (s *E2ETestSuite) retryNtimes(f func() (sdk.TxResponse, error), attempts int) (sdk.TxResponse, error) {
	_ = "STUB: not implemented"
	// Ignore account sequence mismatch errors.
	return *new(sdk.TxResponse), nil
}

// If the response's raw log doesn't contain any of the allowed prefixes we return, else, we retry.

// If the response's raw log doesn't contain any of the allowed prefixes we return, else, we retry.

// AssertTxFailure verifies that an sdk.TxResponse has failed.
func (s *E2ETestSuite) AssertTxFailure(resp sdk.TxResponse, expectedError *errorsmod.Error, alternativeError ...*errorsmod.Error) {
	_ = "STUB: not implemented"
	return
}

// In older versions, the codespace and abci codes were different. So in compatibility tests
// we can not make assertions on them.

// Verify that the error message contains the expected error message or one of the alternative error messages.

// AssertTxSuccess verifies that an sdk.TxResponse has succeeded.
func (s *E2ETestSuite) AssertTxSuccess(resp sdk.TxResponse) { _ = "STUB: not implemented"; return }

// addDebuggingInformation adds additional debugging information to the error message
// based on common types of errors that can occur.
func addDebuggingInformation(errorMsg string) string { _ = "STUB: not implemented"; return "" }

// ExecuteAndPassGovV1Proposal submits a v1 governance proposal using the provided user and message and uses all validators
// to vote yes on the proposal. It ensures the proposal successfully passes.
func (s *E2ETestSuite) ExecuteAndPassGovV1Proposal(ctx context.Context, msg sdk.Msg, chain ibc.Chain, user ibc.Wallet) {
	_ = "STUB: not implemented"
	return
}

// ExecuteGovV1Proposal submits a v1 governance proposal using the provided user and message and uses all validators
// to vote yes on the proposal.
func (s *E2ETestSuite) ExecuteGovV1Proposal(ctx context.Context, msg sdk.Msg, chain ibc.Chain, user ibc.Wallet) error {
	_ = "STUB: not implemented"
	return nil
}

// waitForGovV1ProposalToPass polls for the entire voting period to see if the proposal has passed.
// if the proposal has not passed within the duration of the voting period, an error is returned.
func (s *E2ETestSuite) waitForGovV1ProposalToPass(ctx context.Context, chain ibc.Chain, proposalID uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// poll for the query for the entire voting period to see if the proposal has passed.

// in the case of a failed proposal, we wrap the polling error with additional information about why the proposal failed.

// ExecuteAndPassGovV1Beta1Proposal submits the given v1beta1 governance proposal using the provided user and uses all validators to vote yes on the proposal.
// It ensures the proposal successfully passes.
func (s *E2ETestSuite) ExecuteAndPassGovV1Beta1Proposal(ctx context.Context, chain ibc.Chain, user ibc.Wallet, content govtypesv1beta1.Content) {
	_ = "STUB: not implemented"
	return
}

// ensure voting period has not passed before validators finished voting

// waitForGovV1Beta1ProposalToPass polls for the entire voting period to see if the proposal has passed.
// if the proposal has not passed within the duration of the voting period, an error is returned.
func (*E2ETestSuite) waitForGovV1Beta1ProposalToPass(ctx context.Context, chain ibc.Chain, proposalID uint64) error {
	_ = "STUB: not implemented"
	// poll for the query for the entire voting period to see if the proposal has passed.
	return nil
}

// ExecuteGovV1Beta1Proposal submits a v1beta1 governance proposal using the provided content.
func (s *E2ETestSuite) ExecuteGovV1Beta1Proposal(ctx context.Context, chain ibc.Chain, user ibc.Wallet, content govtypesv1beta1.Content) sdk.TxResponse {
	_ = "STUB: not implemented"
	return *new(sdk.TxResponse)
}

// Transfer broadcasts a MsgTransfer message.
func (s *E2ETestSuite) Transfer(ctx context.Context, chain ibc.Chain, user ibc.Wallet,
	portID, channelID string, token sdk.Coin, sender, receiver string,
	timeoutHeight clienttypes.Height, timeoutTimestamp uint64,
	memo string,
) sdk.TxResponse {
	_ = "STUB: not implemented"
	return *new(sdk.TxResponse)
}

// QueryTxsByEvents runs the QueryTxsByEvents command on the given chain.
// https://github.com/cosmos/cosmos-sdk/blob/65ab2530cc654fd9e252b124ed24cbaa18023b2b/x/auth/client/cli/query.go#L33
func (*E2ETestSuite) QueryTxsByEvents(
	ctx context.Context, chain ibc.Chain,
	page, limit int, queryReq, orderBy string,
) (*txtypes.GetTxsEventResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExtractValueFromEvents extracts the value of an attribute from a list of events.
// If the attribute is not found, the function returns an empty string and false.
// If the attribute is found, the function returns the value and true.
func (*E2ETestSuite) ExtractValueFromEvents(events []abci.Event, eventType, attrKey string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}
