package tendermint

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/v2/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

// VerifyClientMessage checks if the clientMessage is of type Header or Misbehaviour and verifies the message
func (cs *ClientState) VerifyClientMessage(
	ctx sdk.Context, cdc codec.BinaryCodec, clientStore storetypes.KVStore,
	clientMsg exported.ClientMessage,
) error {
	_ = "STUB: not implemented"
	return nil
}

// verifyHeader returns an error if:
// - the client or header provided are not parseable to tendermint types
// - the header is invalid
// - header height is less than or equal to the trusted header height
// - header revision is not equal to trusted header revision
// - header valset commit verification fails
// - header timestamp is past the trusting period in relation to the consensus state
// - header timestamp is less than or equal to the consensus state timestamp
func (cs *ClientState) verifyHeader(
	ctx sdk.Context, clientStore storetypes.KVStore, cdc codec.BinaryCodec,
	header *Header,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Retrieve trusted consensus states for each Header in misbehaviour

// UpdateClient only accepts updates with a header at the same revision
// as the trusted consensus state

// assert header height is newer than consensus state

// Construct a trusted header using the fields in consensus state
// Only Height, Time, and NextValidatorsHash are necessary for verification
// NOTE: updates must be within the same revision

// Verify next header with the passed-in trustedVals
// - asserts trusting period not passed
// - assert header timestamp is not past the trusting period
// - assert header timestamp is past latest stored consensus state timestamp
// - assert that a TrustLevel proportion of TrustedValidators signed new Commit

// UpdateState may be used to either create a consensus state for:
// - a future height greater than the latest client state height
// - a past height that was skipped during bisection
// If we are updating to a past height, a consensus state is created for that height to be persisted in client store
// If we are updating to a future height, the consensus state is created and the client state is updated to reflect
// the new latest height
// A list containing the updated consensus height is returned.
// UpdateState must only be used to update within a single revision, thus header revision number and trusted height's revision
// number must be the same. To update to a new revision, use a separate upgrade path
// UpdateState will prune the oldest consensus state if it is expired.
// If the provided clientMsg is not of type of Header then the handler will noop and empty slice is returned.
func (cs *ClientState) UpdateState(ctx sdk.Context, cdc codec.BinaryCodec, clientStore storetypes.KVStore, clientMsg exported.ClientMessage) []exported.Height {
	_ = "STUB: not implemented"
	return nil
}

// clientMsg is invalid Misbehaviour, no update necessary

// performance: do not prune in checkTx
// simulation must prune for accurate gas estimation

// check for duplicate update

// perform no-op

// set client state, consensus state and associated metadata

// pruneOldestConsensusState will retrieve the earliest consensus state for this clientID and check if it is expired. If it is,
// that consensus state will be pruned from store along with all associated metadata. This will prevent the client store from
// becoming bloated with expired consensus states that can no longer be used for updates and packet verification.
func (cs *ClientState) pruneOldestConsensusState(ctx sdk.Context, cdc codec.BinaryCodec, clientStore storetypes.KVStore) {
	_ = "STUB: not implemented"
	// Check the earliest consensus state to see if it is expired, if so then set the prune height
	// so that we can delete consensus state and all associated metadata.
	return
}

// this error should never occur

// if pruneHeight is set, delete consensus state and metadata

// UpdateStateOnMisbehaviour updates state upon misbehaviour, freezing the ClientState. This method should only be called when misbehaviour is detected
// as it does not perform any misbehaviour checks.
func (cs *ClientState) UpdateStateOnMisbehaviour(ctx sdk.Context, cdc codec.BinaryCodec, clientStore storetypes.KVStore, _ exported.ClientMessage) {
	_ = "STUB: not implemented"
	return
}

// checkTrustedHeader checks that consensus state matches trusted fields of Header
func checkTrustedHeader(header *Header, consState *ConsensusState) error {
	_ = "STUB: not implemented"
	return nil
}

// assert that trustedVals is NextValidators of last trusted header
// to do this, we check that trustedVals.Hash() == consState.NextValidatorsHash
