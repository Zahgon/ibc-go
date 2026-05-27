package tendermint

import (
	"time"

	ics23 "github.com/cosmos/ics23/go"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/v2/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var _ exported.ClientState = (*ClientState)(nil)

// NewClientState creates a new ClientState instance
func NewClientState(
	chainID string, trustLevel Fraction,
	trustingPeriod, ubdPeriod, maxClockDrift time.Duration,
	latestHeight clienttypes.Height, specs []*ics23.ProofSpec,
	upgradePath []string,
) *ClientState {
	_ = "STUB: not implemented"
	return nil
}

// GetChainID returns the chain-id
func (cs ClientState) GetChainID() string {
	_ = "STUB: not implemented"

	// ClientType is tendermint.
	return ""
}

func (ClientState) ClientType() string { _ = "STUB: not implemented"; return "" }

// getTimestampAtHeight returns the timestamp in nanoseconds of the consensus state at the given height.
func (ClientState) getTimestampAtHeight(
	clientStore storetypes.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
) (uint64, error) {
	_ = "STUB: not implemented"
	// get consensus state at height from clientStore to check for expiry
	return 0, nil
}

// status returns the status of the tendermint client.
// The client may be:
// - Active: FrozenHeight is zero and client is not expired
// - Frozen: Frozen Height is not zero
// - Expired: the latest consensus state timestamp + trusting period <= current time
//
// A frozen client will become expired, so the Frozen status
// has higher precedence.
func (cs ClientState) status(
	ctx sdk.Context,
	clientStore storetypes.KVStore,
	cdc codec.BinaryCodec,
) exported.Status {
	_ = "STUB: not implemented"
	return *new(exported.Status)
}

// get latest consensus state from clientStore to check for expiry

// if the client state does not have an associated consensus state for its latest height
// then it must be expired

// IsExpired returns whether or not the client has passed the trusting period since the last
// update (in which case no headers are considered valid).
func (cs ClientState) IsExpired(latestTimestamp, now time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

// Validate performs a basic validation of the client state fields.
func (cs ClientState) Validate() error { _ = "STUB: not implemented"; return nil }

// NOTE: the value of cmttypes.MaxChainIDLen may change in the future.
// If this occurs, the code here must account for potential difference
// between the tendermint version being run by the counterparty chain
// and the tendermint version used by this light client.
// https://github.com/cosmos/ibc-go/issues/177

// the latest height revision number must match the chain id revision number

// UpgradePath may be empty, but if it isn't, each key must be non-empty

// ZeroCustomFields returns a ClientState that is a copy of the current ClientState
// with all client customizable fields zeroed out. All chain specific fields must
// remain unchanged. This client state will be used to verify chain upgrades when a
// chain breaks a light client verification parameter such as chainID.
func (cs ClientState) ZeroCustomFields() *ClientState {
	_ = "STUB: not implemented"
	// copy over all chain-specified fields
	// and leave custom fields empty
	return nil
}

// initialize checks that the initial consensus state is an 07-tendermint consensus state and
// sets the client state, consensus state and associated metadata in the provided client store.
func (cs ClientState) initialize(ctx sdk.Context, cdc codec.BinaryCodec, clientStore storetypes.KVStore, consState exported.ConsensusState) error {
	_ = "STUB: not implemented"
	return nil
}

// verifyMembership is a generic proof verification method which verifies a proof of the existence of a value at a given CommitmentPath at the specified height.
// The caller is expected to construct the full CommitmentPath from a CommitmentPrefix and a standardized path (as defined in ICS 24).
// If a zero proof height is passed in, it will fail to retrieve the associated consensus state.
func (cs ClientState) verifyMembership(
	ctx sdk.Context,
	clientStore storetypes.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
	delayTimePeriod uint64,
	delayBlockPeriod uint64,
	proof []byte,
	path exported.Path,
	value []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

// verifyNonMembership is a generic proof verification method which verifies the absence of a given CommitmentPath at a specified height.
// The caller is expected to construct the full CommitmentPath from a CommitmentPrefix and a standardized path (as defined in ICS 24).
// If a zero proof height is passed in, it will fail to retrieve the associated consensus state.
func (cs ClientState) verifyNonMembership(
	ctx sdk.Context,
	clientStore storetypes.KVStore,
	cdc codec.BinaryCodec,
	height exported.Height,
	delayTimePeriod uint64,
	delayBlockPeriod uint64,
	proof []byte,
	path exported.Path,
) error {
	_ = "STUB: not implemented"
	return nil
}

// verifyDelayPeriodPassed will ensure that at least delayTimePeriod amount of time and delayBlockPeriod number of blocks have passed
// since consensus state was submitted before allowing verification to continue.
func verifyDelayPeriodPassed(ctx sdk.Context, store storetypes.KVStore, proofHeight exported.Height, delayTimePeriod, delayBlockPeriod uint64) error {
	_ = "STUB: not implemented"
	return nil

	// check that executing chain's timestamp has passed consensusState's processed time + delay time period
}

// NOTE: delay time period is inclusive, so if currentTimestamp is validTime, then we return no error

// check that executing chain's height has passed consensusState's processed height + delay block period

// NOTE: delay block period is inclusive, so if currentHeight is validHeight, then we return no error
