package tendermint

import (
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/v2/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

// CheckForMisbehaviour detects duplicate height misbehaviour and BFT time violation misbehaviour
// in a submitted Header message and verifies the correctness of a submitted Misbehaviour ClientMessage
func (ClientState) CheckForMisbehaviour(ctx sdk.Context, cdc codec.BinaryCodec, clientStore storetypes.KVStore, msg exported.ClientMessage) bool {
	_ = "STUB: not implemented"
	return false
}

// Check if the Client store already has a consensus state for the header's height
// If the consensus state exists, and it matches the header then we return early
// since header has already been submitted in a previous UpdateClient.

// This header has already been submitted and the necessary state is already stored
// in client store, thus we can return early without further validation.

// A consensus state already exists for this height, but it does not match the provided header.
// The assumption is that Header has already been validated. Thus we can return true as misbehaviour is present

// Check that consensus state timestamps are monotonic

// if previous consensus state exists, check consensus state time is greater than previous consensus state time
// if previous consensus state is not before current consensus state return true

// if next consensus state exists, check consensus state time is less than next consensus state time
// if next consensus state is not after current consensus state return true

// if heights are equal check that this is valid misbehaviour of a fork
// otherwise if heights are unequal check that this is valid misbehavior of BFT time violation

// Ensure that Commit Hashes are different

// Header1 is at greater height than Header2, therefore Header1 time must be less than or equal to
// Header2 time in order to be valid misbehaviour (violation of monotonic time).

// verifyMisbehaviour determines whether or not two conflicting
// headers at the same height would have convinced the light client.
//
// NOTE: consensusState1 is the trusted consensus state that corresponds to the TrustedHeight
// of misbehaviour.Header1
// Similarly, consensusState2 is the trusted consensus state that corresponds
// to misbehaviour.Header2
// Misbehaviour sets frozen height to {0, 1} since it is only used as a boolean value (zero or non-zero).
func (cs *ClientState) verifyMisbehaviour(ctx sdk.Context, clientStore storetypes.KVStore, cdc codec.BinaryCodec, misbehaviour *Misbehaviour) error {
	_ = "STUB: not implemented"
	// Regardless of the type of misbehaviour, ensure that both headers are valid and would have been accepted by light-client
	return nil
}

// Retrieve trusted consensus states for each Header in misbehaviour

// Check the validity of the two conflicting headers against their respective
// trusted consensus states
// NOTE: header height and commitment root assertions are checked in
// misbehaviour.ValidateBasic by the client keeper and msg.ValidateBasic
// by the base application.

// checkMisbehaviourHeader checks that a Header in Misbehaviour is valid misbehaviour given
// a trusted ConsensusState
func checkMisbehaviourHeader(
	clientState *ClientState, consState *ConsensusState, header *Header, currentTimestamp time.Time,
) error {
	_ = "STUB: not implemented"
	return nil
}

// check the trusted fields for the header against ConsensusState

// assert that the age of the trusted consensus state is not older than the trusting period

// If chainID is in revision format, then set revision number of chainID with the revision number
// of the misbehaviour header
// NOTE: misbehaviour verification is not supported for chains which upgrade to a new chainID without
// strictly following the chainID revision format

// - ValidatorSet must have TrustLevel similarity with trusted FromValidatorSet
// - ValidatorSets on both headers are valid given the last trusted ValidatorSet
