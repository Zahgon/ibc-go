package localhost

import (
	corestore "cosmossdk.io/core/store"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

const (
	// ModuleName defines the 09-localhost light client module name
	ModuleName = "09-localhost"
)

// SentinelProof defines the 09-localhost sentinel proof.
// Submission of nil or empty proofs is disallowed in core IBC messaging.
// This serves as a placeholder value for relayers to leverage as the proof field in various message types.
// Localhost client state verification will fail if the sentintel proof value is not provided.
var SentinelProof = []byte{0x01}

var _ exported.LightClientModule = (*LightClientModule)(nil)

// LightClientModule implements the core IBC api.LightClientModule interface.
type LightClientModule struct {
	cdc          codec.BinaryCodec
	storeService corestore.KVStoreService
}

// NewLightClientModule creates and returns a new 09-localhost LightClientModule.
func NewLightClientModule(cdc codec.BinaryCodec, storeService corestore.KVStoreService) *LightClientModule {
	_ = "STUB: not implemented"
	return nil
}

// Initialize returns an error because it is stateless.
func (LightClientModule) Initialize(_ sdk.Context, _ string, _, _ []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyClientMessage is unsupported by the 09-localhost client type and returns an error.
func (LightClientModule) VerifyClientMessage(_ sdk.Context, _ string, _ exported.ClientMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// CheckForMisbehaviour is unsupported by the 09-localhost client type and performs a no-op, returning false.
func (LightClientModule) CheckForMisbehaviour(_ sdk.Context, _ string, _ exported.ClientMessage) bool {
	_ = "STUB: not implemented"

	// UpdateStateOnMisbehaviour is unsupported by the 09-localhost client type and performs a no-op.
	return false
}

func (LightClientModule) UpdateStateOnMisbehaviour(_ sdk.Context, _ string, _ exported.ClientMessage) {
	_ = "STUB: not implemented"
	// no-op

	// UpdateState performs a no-op and returns the context height in the updated heights return value.
	return
}

func (LightClientModule) UpdateState(ctx sdk.Context, _ string, _ exported.ClientMessage) []exported.Height {
	_ = "STUB: not implemented"
	return nil
}

// VerifyMembership is a generic proof verification method which verifies the existence of a given key and value within the IBC store.
// The caller is expected to construct the full CommitmentPath from a CommitmentPrefix and a standardized path (as defined in ICS 24).
// The caller must provide the full IBC store.
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

// ensure the proof provided is the expected sentinel localhost client proof

// The commitment prefix (eg: "ibc") is omitted when operating on the core IBC store

// VerifyNonMembership is a generic proof verification method which verifies the absence of a given CommitmentPath within the IBC store.
// The caller is expected to construct the full CommitmentPath from a CommitmentPrefix and a standardized path (as defined in ICS 24).
// The caller must provide the full IBC store.
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

// ensure the proof provided is the expected sentinel localhost client proof

// The commitment prefix (eg: "ibc") is omitted when operating on the core IBC store

// Status always returns Active. The 09-localhost status cannot be changed.
func (LightClientModule) Status(_ sdk.Context, _ string) exported.Status {
	_ = "STUB: not implemented"
	return *

	// LatestHeight returns the context height.
	new(exported.Status)
}

func (LightClientModule) LatestHeight(ctx sdk.Context, _ string) exported.Height {
	_ = "STUB: not implemented"
	return *new(exported.Height)
}

// TimestampAtHeight returns the current block time retrieved from the application context. The localhost client does not store consensus states and thus
// cannot provide a timestamp for the provided height.
func (LightClientModule) TimestampAtHeight(ctx sdk.Context, _ string, _ exported.Height) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// RecoverClient returns an error. The localhost cannot be modified by proposals.
func (LightClientModule) RecoverClient(_ sdk.Context, _, _ string) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyUpgradeAndUpdateState returns an error since localhost cannot be upgraded.
func (LightClientModule) VerifyUpgradeAndUpdateState(_ sdk.Context, _ string, _, _, _, _ []byte) error {
	_ = "STUB: not implemented"
	return nil
}
