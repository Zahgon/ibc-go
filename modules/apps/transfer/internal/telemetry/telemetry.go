package telemetry

import (
	"github.com/cosmos/ibc-go/v11/modules/apps/transfer/types"
)

func ReportTransfer(sourcePort, sourceChannel, destinationPort, destinationChannel string, token types.Token) {
	_ = "STUB: not implemented"
	return
}

func ReportOnRecvPacket(sourcePort, sourceChannel, destinationPort, destinationChannel string, token types.Token) {
	_ = "STUB: not implemented"
	return
}

// Modify trace as Recv does.

// Transfer amount has already been parsed in caller.
