package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

// emitChannelOpenInitEvent emits a channel open init event
func emitChannelOpenInitEvent(ctx sdk.Context, portID string, channelID string, channel types.Channel) {
	_ = "STUB: not implemented"
	return
}

// emitChannelOpenTryEvent emits a channel open try event
func emitChannelOpenTryEvent(ctx sdk.Context, portID string, channelID string, channel types.Channel) {
	_ = "STUB: not implemented"
	return
}

// emitChannelOpenAckEvent emits a channel open acknowledge event
func emitChannelOpenAckEvent(ctx sdk.Context, portID string, channelID string, channel types.Channel) {
	_ = "STUB: not implemented"
	return
}

// emitChannelOpenConfirmEvent emits a channel open confirm event
func emitChannelOpenConfirmEvent(ctx sdk.Context, portID string, channelID string, channel types.Channel) {
	_ = "STUB: not implemented"
	return
}

// emitChannelCloseInitEvent emits a channel close init event
func emitChannelCloseInitEvent(ctx sdk.Context, portID string, channelID string, channel types.Channel) {
	_ = "STUB: not implemented"
	return
}

// emitChannelCloseConfirmEvent emits a channel close confirm event
func emitChannelCloseConfirmEvent(ctx sdk.Context, portID string, channelID string, channel types.Channel) {
	_ = "STUB: not implemented"
	return
}

// emitSendPacketEvent emits an event with packet data along with other packet information for relayer
// to pick up and relay to other chain
func emitSendPacketEvent(ctx sdk.Context, packet types.Packet, channel types.Channel, timeoutHeight exported.Height) {
	_ = "STUB: not implemented"
	return
}

// we only support 1-hop packets now, and that is the most important hop for a relayer
// (is it going to a chain I am connected to)
// DEPRECATED

// emitRecvPacketEvent emits a receive packet event. It will be emitted both the first time a packet
// is received for a certain sequence and for all duplicate receives.
func emitRecvPacketEvent(ctx sdk.Context, packet types.Packet, channel types.Channel) {
	_ = "STUB: not implemented"
	return
}

// we only support 1-hop packets now, and that is the most important hop for a relayer
// (is it going to a chain I am connected to)
// DEPRECATED

// emitWriteAcknowledgementEvent emits an event that the relayer can query for
func emitWriteAcknowledgementEvent(ctx sdk.Context, packet types.Packet, channel types.Channel, acknowledgement []byte) {
	_ = "STUB: not implemented"
	return
}

// we only support 1-hop packets now, and that is the most important hop for a relayer
// (is it going to a chain I am connected to)
// DEPRECATED

// emitAcknowledgePacketEvent emits an acknowledge packet event. It will be emitted both the first time
// a packet is acknowledged for a certain sequence and for all duplicate acknowledgements.
func emitAcknowledgePacketEvent(ctx sdk.Context, packet types.Packet, channel types.Channel) {
	_ = "STUB: not implemented"
	return
}

// we only support 1-hop packets now, and that is the most important hop for a relayer
// (is it going to a chain I am connected to)
// DEPRECATED

// emitTimeoutPacketEvent emits a timeout packet event. It will be emitted both the first time a packet
// is timed out for a certain sequence and for all duplicate timeouts.
func emitTimeoutPacketEvent(ctx sdk.Context, packet types.Packet, channel types.Channel) {
	_ = "STUB: not implemented"
	return
}

// emitChannelClosedEvent emits a channel closed event.
func emitChannelClosedEvent(ctx sdk.Context, packet types.Packet, channel types.Channel) {
	_ = "STUB: not implemented"
	return
}
