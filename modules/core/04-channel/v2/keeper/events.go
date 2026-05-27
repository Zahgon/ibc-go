package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/core/04-channel/v2/types"
)

// emitSendPacketEvents emits events for the SendPacket handler.
func emitSendPacketEvents(ctx sdk.Context, packet types.Packet) { _ = "STUB: not implemented"; return }

// emitRecvPacketEvents emits events for the RecvPacket handler.
func emitRecvPacketEvents(ctx sdk.Context, packet types.Packet) { _ = "STUB: not implemented"; return }

// emitWriteAcknowledgementEvents emits events for WriteAcknowledgement.
func emitWriteAcknowledgementEvents(ctx sdk.Context, packet types.Packet, ack types.Acknowledgement) {
	_ = "STUB: not implemented"
	return
}

// emitAcknowledgePacketEvents emits events for the AcknowledgePacket handler.
func emitAcknowledgePacketEvents(ctx sdk.Context, packet types.Packet) {
	_ = "STUB: not implemented"
	return
}

// emitTimeoutPacketEvents emits events for the TimeoutPacket handler.
func emitTimeoutPacketEvents(ctx sdk.Context, packet types.Packet) {
	_ = "STUB: not implemented"
	return
}
