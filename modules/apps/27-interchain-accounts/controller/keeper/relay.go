package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	icatypes "github.com/cosmos/ibc-go/v11/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
)

// SendTx takes pre-built packet data containing messages to be executed on the host chain from an authentication module and attempts to send the packet.
// The packet sequence for the outgoing packet is returned as a result. An appropriate
// absolute timeoutTimestamp must be provided. If the packet is timed out, the channel will be closed.
// In the case of channel closure, a new channel may be reopened to reconnect to the host chain.
//
// Deprecated: this is a legacy API that is only intended to function correctly in workflows where an underlying application has been set.
// Prior to v6.x.x of ibc-go, the controller module was only functional as middleware, with authentication performed
// by the underlying application. For a full summary of the changes in v6.x.x, please see ADR009.
// This API will be removed in later releases.
func (k *Keeper) SendTx(ctx sdk.Context, connectionID, portID string,
	icaPacketData icatypes.InterchainAccountPacketData, timeoutTimestamp uint64,
) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (k *Keeper) sendTx(ctx sdk.Context, connectionID, portID string,
	icaPacketData icatypes.InterchainAccountPacketData, timeoutTimestamp uint64,
) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// OnTimeoutPacket removes the active channel associated with the provided packet, the underlying channel end is closed
// due to the semantics of ORDERED channels
func (*Keeper) OnTimeoutPacket(ctx sdk.Context, packet channeltypes.Packet) error {
	_ = "STUB: not implemented"
	return nil
}
