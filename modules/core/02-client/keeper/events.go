package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

// emitCreateClientEvent emits a create client event
func emitCreateClientEvent(ctx sdk.Context, clientID, clientType string, initialHeight exported.Height) {
	_ = "STUB: not implemented"
	return
}

// emitUpdateClientEvent emits an update client event
func emitUpdateClientEvent(ctx sdk.Context, clientID string, clientType string, consensusHeights []exported.Height, _ codec.BinaryCodec, _ exported.ClientMessage) {
	_ = "STUB: not implemented"
	return
}

// Deprecated: AttributeKeyConsensusHeight is deprecated and will be removed in a future release.
// Please use AttributeKeyConsensusHeights instead.

// emitUpgradeClientEvent emits an upgrade client event
func emitUpgradeClientEvent(ctx sdk.Context, clientID, clientType string, latestHeight exported.Height) {
	_ = "STUB: not implemented"
	return
}

// emitSubmitMisbehaviourEvent emits a client misbehaviour event
func emitSubmitMisbehaviourEvent(ctx sdk.Context, clientID string, clientType string) {
	_ = "STUB: not implemented"
	return
}

// emitRecoverClientEvent emits a recover client event
func emitRecoverClientEvent(ctx sdk.Context, clientID, clientType string) {
	_ = "STUB: not implemented"
	return
}

// emitScheduleIBCSoftwareUpgradeEvent emits a schedule IBC software upgrade event
func emitScheduleIBCSoftwareUpgradeEvent(ctx sdk.Context, title string, height int64) {
	_ = "STUB: not implemented"
	return
}

// EmitUpgradeChainEvent emits an upgrade chain event.
func EmitUpgradeChainEvent(ctx sdk.Context, height int64) { _ = "STUB: not implemented"; return }

// which store to query proof of consensus state from
