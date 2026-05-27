package events

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/apps/27-gmp/types"
)

// EmitSendCall emits a GMP send call event.
func EmitSendCall(
	ctx sdk.Context,
	packetData types.GMPPacketData,
	sourceClient,
	destinationClient,
	sourcePort,
	destinationPort string,
	sequence uint64,
) {
	_ = "STUB: not implemented"
	return
}

// EmitOnRecvPacketEvent emits a GMP packet event in the OnRecvPacket callback.
func EmitOnRecvPacketEvent(
	ctx sdk.Context,
	packetData types.GMPPacketData,
	sourceClient,
	destinationClient,
	sourcePort,
	destinationPort string,
	sequence uint64,
	ackErr error,
) {
	_ = "STUB: not implemented"
	return
}

func packetAttributes(
	packetData types.GMPPacketData,
	sourceClient,
	destinationClient,
	sourcePort,
	destinationPort string,
	sequence uint64,
) []sdk.Attribute {
	_ = "STUB: not implemented"
	return nil
}
