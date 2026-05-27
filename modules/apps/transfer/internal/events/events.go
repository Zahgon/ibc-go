package events

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
	ibcexported "github.com/cosmos/ibc-go/v11/modules/core/exported"
)

// EmitTransferEvent emits a ibc transfer event on successful transfers.
func EmitTransferEvent(ctx sdk.Context, sender, receiver string, token types.Token, memo string) {
	_ = "STUB: not implemented"
	return
}

// EmitOnRecvPacketEvent emits a fungible token packet event in the OnRecvPacket callback
func EmitOnRecvPacketEvent(ctx sdk.Context, packetData types.InternalTransferRepresentation, ack ibcexported.Acknowledgement, ackErr error) {
	_ = "STUB: not implemented"
	return
}

// EmitOnAcknowledgementPacketEvent emits a fungible token packet event in the OnAcknowledgementPacket callback
func EmitOnAcknowledgementPacketEvent(ctx sdk.Context, packetData types.InternalTransferRepresentation, ack channeltypes.Acknowledgement) {
	_ = "STUB: not implemented"
	return
}

// EmitOnTimeoutEvent emits a fungible token packet event in the OnTimeoutPacket callback
func EmitOnTimeoutEvent(ctx sdk.Context, packetData types.InternalTransferRepresentation) {
	_ = "STUB: not implemented"
	return
}

// EmitDenomEvent emits a denomination event in the OnRecv callback.
func EmitDenomEvent(ctx sdk.Context, token types.Token) { _ = "STUB: not implemented"; return }

// mustMarshalJSON json marshals the given type and panics on failure.
func mustMarshalJSON(v any) string { _ = "STUB: not implemented"; return "" }
