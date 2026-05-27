package v2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/apps/transfer/keeper"
	channeltypesv2 "github.com/cosmos/ibc-go/v11/modules/core/04-channel/v2/types"
	"github.com/cosmos/ibc-go/v11/modules/core/api"
)

var _ api.IBCModule = (*IBCModule)(nil)

// NewIBCModule creates a new IBCModule given the keeper
func NewIBCModule(k *keeper.Keeper) IBCModule { _ = "STUB: not implemented"; return *new(IBCModule) }

type IBCModule struct {
	keeper *keeper.Keeper
}

func (im IBCModule) OnSendPacket(ctx sdk.Context, sourceChannel string, destinationChannel string, sequence uint64, payload channeltypesv2.Payload, signer sdk.AccAddress) error {
	_ = "STUB: not implemented"
	// Enforce that the source and destination portIDs are the same and equal to the transfer portID
	// Enforce that the source and destination clientIDs are also in the clientID format that transfer expects: {clientid}-{sequence}
	// This is necessary for IBC v2 since the portIDs (and thus the application-application connection) is not prenegotiated
	// by the channel handshake
	// This restriction can be removed in a future where the trace hop on receive commits to **both** the source and destination portIDs
	// rather than just the destination port
	return nil
}

// Enforce that the base denom does not contain any slashes
// Since IBC v2 packets will no longer have channel identifiers, we cannot rely
// on the channel format to easily divide the trace from the base denomination in ICS20 v1 packets
// The simplest way to prevent any potential issues from arising is to simply disallow any slashes in the base denomination
// This prevents such denominations from being sent with IBCV v2 packets, however we can still support them in IBC v1 packets
// If we enforce that IBC v2 packets are sent with ICS20 v2 and above versions that separate the trace from the base denomination
// in the packet data, then we can remove this restriction.

func (im IBCModule) OnRecvPacket(ctx sdk.Context, sourceChannel string, destinationChannel string, sequence uint64, payload channeltypesv2.Payload, relayer sdk.AccAddress) channeltypesv2.RecvPacketResult {
	_ = "STUB: not implemented"
	// Enforce that the source and destination portIDs are the same and equal to the transfer portID
	// Enforce that the source and destination clientIDs are also in the clientID format that transfer expects: {clientid}-{sequence}
	// This is necessary for IBC v2 since the portIDs (and thus the application-application connection) is not prenegotiated
	// by the channel handshake
	// This restriction can be removed in a future where the trace hop on receive commits to **both** the source and destination portIDs
	// rather than just the destination port
	return *new(channeltypesv2.RecvPacketResult)
}

// we are explicitly wrapping this emit event call in an anonymous function so that
// the packet data is evaluated after it has been assigned a value.

// NOTE: acknowledgement will be written synchronously during IBC handler execution.

func (im IBCModule) OnTimeoutPacket(ctx sdk.Context, sourceChannel string, destinationChannel string, sequence uint64, payload channeltypesv2.Payload, relayer sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}

// refund tokens

func (im IBCModule) OnAcknowledgementPacket(ctx sdk.Context, sourceChannel string, destinationChannel string, sequence uint64, acknowledgement []byte, payload channeltypesv2.Payload, relayer sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}

// construct an error acknowledgement if the acknowledgement bytes are the sentinel error acknowledgement so we can use the shared transfer logic

// the specific error does not matter

// UnmarshalPacketData unmarshals the ICS20 packet data based on the version and encoding
// it implements the PacketDataUnmarshaler interface
func (IBCModule) UnmarshalPacketData(payload channeltypesv2.Payload) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
