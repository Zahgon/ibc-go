package telemetry

import (
	"github.com/cosmos/ibc-go/v11/modules/core/04-channel/v2/types"
)

func ReportRecvPacket(packet types.Packet) { _ = "STUB: not implemented"; return }

func ReportTimeoutPacket(packet types.Packet) { _ = "STUB: not implemented"; return }

func ReportAcknowledgePacket(packet types.Packet) { _ = "STUB: not implemented"; return }
