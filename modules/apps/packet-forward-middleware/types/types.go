package types

import (
	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
)

func (ifp *InFlightPacket) ChannelPacket() channeltypes.Packet {
	_ = "STUB: not implemented"
	return *new(channeltypes.Packet)
}
