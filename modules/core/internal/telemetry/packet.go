package telemetry

import (
	metrics "github.com/hashicorp/go-metrics"

	"github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
)

func ReportRecvPacket(packet types.Packet) { _ = "STUB: not implemented"; return }

func ReportTimeoutPacket(packet types.Packet, timeoutType string) {
	_ = "STUB: not implemented"
	return
}

func ReportAcknowledgePacket(packet types.Packet) { _ = "STUB: not implemented"; return }

func addPacketLabels(packet types.Packet) []metrics.Label { _ = "STUB: not implemented"; return nil }
