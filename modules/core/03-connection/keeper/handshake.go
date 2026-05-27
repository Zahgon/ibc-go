package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/core/03-connection/types"
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

// ConnOpenInit initialises a connection attempt on chain A. The generated connection identifier
// is returned.
//
// NOTE: Msg validation verifies the supplied identifiers and ensures that the counterparty
// connection identifier is empty.
func (k *Keeper) ConnOpenInit(
	ctx sdk.Context,
	clientID string,
	counterparty types.Counterparty, // counterpartyPrefix, counterpartyClientIdentifier
	version *types.Version,
	delayPeriod uint64,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// connection defines chain A's ConnectionEnd

// ConnOpenTry relays notice of a connection attempt on chain A to chain B (this
// code is executed on chain B).
//
// NOTE:
//   - Here chain A acts as the counterparty
//   - Identifiers are checked on msg validation
func (k *Keeper) ConnOpenTry(
	ctx sdk.Context,
	counterparty types.Counterparty, // counterpartyConnectionIdentifier, counterpartyPrefix and counterpartyClientIdentifier
	delayPeriod uint64,
	clientID string, // clientID of chainA
	counterpartyVersions []*types.Version, // supported versions of chain A
	initProof []byte, // proof that chainA stored connectionEnd in state (on ConnOpenInit)
	proofHeight exported.Height, // height at which relayer constructs proof of A storing connectionEnd in state
) (string, error) {
	_ = "STUB: not implemented"
	// generate a new connection
	return "", nil
}

// expectedConnection defines Chain A's ConnectionEnd
// NOTE: chain A's counterparty is chain B (i.e where this code is executed)
// NOTE: chainA and chainB must have the same delay period

// chain B picks a version from Chain A's available versions that is compatible
// with Chain B's supported IBC versions. PickVersion will select the intersection
// of the supported versions and the counterparty versions.

// connection defines chain B's ConnectionEnd

// Check that ChainA committed expectedConnectionEnd to its state

// store connection in chainB state

// ConnOpenAck relays acceptance of a connection open attempt from chain B back
// to chain A (this code is executed on chain A).
//
// NOTE: Identifiers are checked on msg validation.
func (k *Keeper) ConnOpenAck(
	ctx sdk.Context,
	connectionID string,
	version *types.Version, // version that ChainB chose in ConnOpenTry
	counterpartyConnectionID string,
	tryProof []byte, // proof that connectionEnd was added to ChainB state in ConnOpenTry
	proofHeight exported.Height, // height that relayer constructed proofTry
) error {
	_ = "STUB: not implemented"
	// Retrieve connection
	return nil
}

// verify the previously set connection state

// ensure selected version is supported

// Ensure that ChainB stored expected connectionEnd in its state during ConnOpenTry

// Update connection state to Open

// ConnOpenConfirm confirms opening of a connection on chain A to chain B, after
// which the connection is open on both chains (this code is executed on chain B).
//
// NOTE: Identifiers are checked on msg validation.
func (k *Keeper) ConnOpenConfirm(
	ctx sdk.Context,
	connectionID string,
	ackProof []byte, // proof that connection opened on ChainA during ConnOpenAck
	proofHeight exported.Height, // height that relayer constructed proofAck
) error {
	_ = "STUB: not implemented"
	// Retrieve connection
	return nil
}

// Check that connection state on ChainB is on state: TRYOPEN

// Check that connection on ChainA is open

// Update ChainB's connection to Open
