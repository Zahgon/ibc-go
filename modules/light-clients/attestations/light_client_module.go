package attestations

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var _ exported.LightClientModule = (*LightClientModule)(nil)

// LightClientModule implements the core IBC api.LightClientModule interface.
type LightClientModule struct {
	cdc           codec.BinaryCodec
	storeProvider clienttypes.StoreProvider
}

// NewLightClientModule creates and returns a new attestations LightClientModule.
func NewLightClientModule(cdc codec.BinaryCodec, storeProvider clienttypes.StoreProvider) LightClientModule {
	_ = "STUB: not implemented"
	return *new(LightClientModule)
}

// Initialize unmarshals the provided client and consensus states and performs basic validation.
func (l LightClientModule) Initialize(ctx sdk.Context, clientID string, clientStateBz, consensusStateBz []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyClientMessage obtains the client state associated with the client identifier and calls into the clientState.VerifyClientMessage method.
func (l LightClientModule) VerifyClientMessage(ctx sdk.Context, clientID string, clientMsg exported.ClientMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// CheckForMisbehaviour returns true if the provided client message contains conflicting timestamps
func (l LightClientModule) CheckForMisbehaviour(ctx sdk.Context, clientID string, clientMsg exported.ClientMessage) bool {
	_ = "STUB: not implemented"
	return false
}

// UpdateStateOnMisbehaviour freezes the client
func (l LightClientModule) UpdateStateOnMisbehaviour(ctx sdk.Context, clientID string, _ exported.ClientMessage) {
	_ = "STUB: not implemented"
	return
}

// UpdateState obtains the client state associated with the client identifier and calls into the clientState.UpdateState method.
func (l LightClientModule) UpdateState(ctx sdk.Context, clientID string, clientMsg exported.ClientMessage) []exported.Height {
	_ = "STUB: not implemented"
	return nil
}

// VerifyMembership obtains the client state associated with the client identifier and calls into the clientState.verifyMembership method.
func (l LightClientModule) VerifyMembership(
	ctx sdk.Context,
	clientID string,
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

// VerifyNonMembership obtains the client state associated with the client identifier and calls into the clientState.verifyNonMembership method.
func (l LightClientModule) VerifyNonMembership(
	ctx sdk.Context,
	clientID string,
	height exported.Height,
	delayTimePeriod uint64,
	delayBlockPeriod uint64,
	proof []byte,
	path exported.Path,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Status returns the status of the attestations client.
// The client may be:
// - Active: if `IsFrozen` is false.
// - Frozen: if `IsFrozen` is true.
// - Unknown: if the client state associated with the provided client identifier is not found.
func (l LightClientModule) Status(ctx sdk.Context, clientID string) exported.Status {
	_ = "STUB: not implemented"
	return *new(exported.Status)
}

// LatestHeight returns the latest height for the client state for the given client identifier.
// If no client is present for the provided client identifier a zero value height is returned.
// NOTE: RevisionNumber is always 0 for attestations client heights.
func (l LightClientModule) LatestHeight(ctx sdk.Context, clientID string) exported.Height {
	_ = "STUB: not implemented"
	return *new(exported.Height)
}

// TimestampAtHeight obtains the client state associated with the client identifier and returns the timestamp in nanoseconds of the consensus state at the given height.
func (l LightClientModule) TimestampAtHeight(ctx sdk.Context, clientID string, height exported.Height) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// RecoverClient is not supported in this version.
func (LightClientModule) RecoverClient(ctx sdk.Context, clientID, substituteClientID string) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyUpgradeAndUpdateState returns an error since attestations client does not support upgrades.
func (LightClientModule) VerifyUpgradeAndUpdateState(ctx sdk.Context, clientID string, newClient, newConsState, upgradeClientProof, upgradeConsensusStateProof []byte) error {
	_ = "STUB: not implemented"
	return nil
}
