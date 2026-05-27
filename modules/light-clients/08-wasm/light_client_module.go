package wasm

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	wasmkeeper "github.com/cosmos/ibc-go/modules/light-clients/08-wasm/v11/keeper"
	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var _ exported.LightClientModule = (*LightClientModule)(nil)

// LightClientModule implements the core IBC api.LightClientModule interface.
type LightClientModule struct {
	keeper        wasmkeeper.Keeper
	storeProvider clienttypes.StoreProvider
}

// NewLightClientModule creates and returns a new 08-wasm LightClientModule.
func NewLightClientModule(keeper wasmkeeper.Keeper, storeProvider clienttypes.StoreProvider) LightClientModule {
	_ = "STUB: not implemented"
	return *new(LightClientModule)
}

// Initialize unmarshals the provided client and consensus states and performs basic validation. It sets the client
// state and consensus state in the client store.
// It also initializes the wasm contract for the client.
func (l LightClientModule) Initialize(ctx sdk.Context, clientID string, clientStateBz, consensusStateBz []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Do not allow initialization of a client with a checksum that hasn't been previously stored via storeWasmCode.

// VerifyClientMessage obtains the client state associated with the client identifier, it then must verify the ClientMessage.
// A ClientMessage could be a Header, Misbehaviour, or batch update.
// It must handle each type of ClientMessage appropriately. Calls to CheckForMisbehaviour, UpdateState, and UpdateStateOnMisbehaviour
// will assume that the content of the ClientMessage has been verified and can be trusted. An error should be returned
// if the ClientMessage fails to verify.
func (l LightClientModule) VerifyClientMessage(ctx sdk.Context, clientID string, clientMsg exported.ClientMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// CheckForMisbehaviour obtains the client state associated with the client identifier, it detects misbehaviour in a submitted Header
// message and verifies the correctness of a submitted Misbehaviour ClientMessage.
func (l LightClientModule) CheckForMisbehaviour(ctx sdk.Context, clientID string, clientMsg exported.ClientMessage) bool {
	_ = "STUB: not implemented"
	return false
}

// UpdateStateOnMisbehaviour obtains the client state associated with the client identifier performs appropriate state changes on
// a client state given that misbehaviour has been detected and verified.
// Client state is updated in the store by the contract.
func (l LightClientModule) UpdateStateOnMisbehaviour(ctx sdk.Context, clientID string, clientMsg exported.ClientMessage) {
	_ = "STUB: not implemented"
	return
}

// UpdateState obtains the client state associated with the client identifier and calls into the appropriate
// contract endpoint. Client state and new consensus states are updated in the store by the contract.
func (l LightClientModule) UpdateState(ctx sdk.Context, clientID string, clientMsg exported.ClientMessage) []exported.Height {
	_ = "STUB: not implemented"
	return nil
}

// VerifyMembership obtains the client state associated with the client identifier and calls into the appropriate contract endpoint.
// VerifyMembership is a generic proof verification method which verifies a proof of the existence of a value at a given CommitmentPath at the specified height.
// The caller is expected to construct the full CommitmentPath from a CommitmentPrefix and a standardized path (as defined in ICS 24).
// If a zero proof height is passed in, it will fail to retrieve the associated consensus state.
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

// VerifyNonMembership obtains the client state associated with the client identifier and calls into the appropriate contract endpoint.
// VerifyNonMembership is a generic proof verification method which verifies the absence of a given CommitmentPath at a specified height.
// The caller is expected to construct the full CommitmentPath from a CommitmentPrefix and a standardized path (as defined in ICS 24).
// If a zero proof height is passed in, it will fail to retrieve the associated consensus state.
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

// Status obtains the client state associated with the client identifier and calls into the appropriate contract endpoint.
// It returns the status of the wasm client.
// The client may be:
// - Active: frozen height is zero and client is not expired
// - Frozen: frozen height is not zero
// - Expired: the latest consensus state timestamp + trusting period <= current time
// - Unauthorized: the client type is not registered as an allowed client type
//
// A frozen client will become expired, so the Frozen status
// has higher precedence.
func (l LightClientModule) Status(ctx sdk.Context, clientID string) exported.Status {
	_ = "STUB: not implemented"
	return *new(exported.Status)
}

// Return unauthorized if the checksum hasn't been previously stored via storeWasmCode.

// LatestHeight returns the latest height for the client state for the given client identifier.
// If no client is present for the provided client identifier a zero value height is returned.
func (l LightClientModule) LatestHeight(ctx sdk.Context, clientID string) exported.Height {
	_ = "STUB: not implemented"
	return *new(exported.Height)
}

// TimestampAtHeight obtains the client state associated with the client identifier and calls into the appropriate contract endpoint.
// It returns the timestamp in nanoseconds of the consensus state at the given height.
func (l LightClientModule) TimestampAtHeight(ctx sdk.Context, clientID string, height exported.Height) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// RecoverClient asserts that the substitute client is a wasm client. It obtains the client state associated with the
// subject client and calls into the appropriate contract endpoint.
// It will verify that a substitute client state is valid and update the subject client state.
// Note that this method is used only for recovery and will not allow changes to the checksum.
func (l LightClientModule) RecoverClient(ctx sdk.Context, clientID, substituteClientID string) error {
	_ = "STUB: not implemented"
	return nil
}

// check that checksums of subject client state and substitute client state match
// changing the checksum is only allowed through the migrate contract RPC endpoint

// VerifyUpgradeAndUpdateState obtains the client state associated with the client identifier and calls into the appropriate contract endpoint.
// The new client and consensus states will be unmarshaled and an error is returned if the new client state is not at a height greater
// than the existing client. On a successful verification, it expects the contract to update the new client state, consensus state, and any other client metadata.
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
