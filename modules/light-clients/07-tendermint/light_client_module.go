package tendermint

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

// NewLightClientModule creates and returns a new 07-tendermint LightClientModule.
func NewLightClientModule(cdc codec.BinaryCodec, storeProvider clienttypes.StoreProvider) LightClientModule {
	_ = "STUB: not implemented"
	return *new(LightClientModule)
}

// Initialize unmarshals the provided client and consensus states and performs basic validation. It calls into the
// clientState.initialize method.
func (l LightClientModule) Initialize(ctx sdk.Context, clientID string, clientStateBz, consensusStateBz []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyClientMessage obtains the client state associated with the client identifier and calls into the clientState.VerifyClientMessage method.
func (l LightClientModule) VerifyClientMessage(ctx sdk.Context, clientID string, clientMsg exported.ClientMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// CheckForMisbehaviour obtains the client state associated with the client identifier and calls into the clientState.CheckForMisbehaviour method.
func (l LightClientModule) CheckForMisbehaviour(ctx sdk.Context, clientID string, clientMsg exported.ClientMessage) bool {
	_ = "STUB: not implemented"
	return false
}

// UpdateStateOnMisbehaviour obtains the client state associated with the client identifier and calls into the clientState.UpdateStateOnMisbehaviour method.
func (l LightClientModule) UpdateStateOnMisbehaviour(ctx sdk.Context, clientID string, clientMsg exported.ClientMessage) {
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

// Status obtains the client state associated with the client identifier and calls into the clientState.status method.
func (l LightClientModule) Status(ctx sdk.Context, clientID string) exported.Status {
	_ = "STUB: not implemented"
	return *new(exported.Status)
}

// LatestHeight returns the latest height for the client state for the given client identifier.
// If no client is present for the provided client identifier a zero value height is returned.
func (l LightClientModule) LatestHeight(ctx sdk.Context, clientID string) exported.Height {
	_ = "STUB: not implemented"
	return *new(exported.Height)
}

// TimestampAtHeight obtains the client state associated with the client identifier and calls into the clientState.getTimestampAtHeight method.
func (l LightClientModule) TimestampAtHeight(
	ctx sdk.Context,
	clientID string,
	height exported.Height,
) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// RecoverClient asserts that the substitute client is a tendermint client. It obtains the client state associated with the
// subject client and calls into the subjectClientState.CheckSubstituteAndUpdateState method.
func (l LightClientModule) RecoverClient(ctx sdk.Context, clientID, substituteClientID string) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyUpgradeAndUpdateState obtains the client state associated with the client identifier and calls into the clientState.VerifyUpgradeAndUpdateState method.
// The new client and consensus states will be unmarshaled and an error is returned if the new client state is not at a height greater
// than the existing client.
func (l LightClientModule) VerifyUpgradeAndUpdateState(
	ctx sdk.Context,
	clientID string,
	newClient []byte,
	newConsState []byte,
	upgradeClientProof,
	upgradeConsensusStateProof []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

// last height of current counterparty chain must be client's latest height
