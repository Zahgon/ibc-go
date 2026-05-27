package gmp

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/apps/27-gmp/keeper"
	channeltypesv2 "github.com/cosmos/ibc-go/v11/modules/core/04-channel/v2/types"
	"github.com/cosmos/ibc-go/v11/modules/core/api"
)

var (
	_ api.IBCModule             = (*IBCModule)(nil)
	_ api.PacketDataUnmarshaler = (*IBCModule)(nil)
)

// IBCModule implements the ICS26 interface for transfer given the transfer keeper.
type IBCModule struct {
	keeper *keeper.Keeper
}

// NewIBCModule creates a new IBCModule given the keeper
func NewIBCModule(k *keeper.Keeper) *IBCModule { _ = "STUB: not implemented"; return nil }

func (*IBCModule) OnSendPacket(ctx sdk.Context, sourceChannel string, destinationChannel string, sequence uint64, payload channeltypesv2.Payload, signer sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}

func (im *IBCModule) OnRecvPacket(ctx sdk.Context, sourceClient, destinationClient string, sequence uint64, payload channeltypesv2.Payload, relayer sdk.AccAddress) channeltypesv2.RecvPacketResult {
	_ = "STUB: not implemented"
	return *new(channeltypesv2.RecvPacketResult)
}

// we are explicitly wrapping this emit event call in an anonymous function so that
// the packet data is evaluated after it has been assigned a value.

// TODO: implement telemetry

func (*IBCModule) OnTimeoutPacket(_ sdk.Context, _, _ string, _ uint64, _ channeltypesv2.Payload, _ sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}

func (*IBCModule) OnAcknowledgementPacket(_ sdk.Context, _, _ string, _ uint64, _ []byte, _ channeltypesv2.Payload, _ sdk.AccAddress) error {
	_ = "STUB: not implemented"

	// UnmarshalPacketData unmarshals GMP packet data from the payload.
	// This method implements the PacketDataUnmarshaler interface required for callbacks middleware support.
	return nil
}

func (*IBCModule) UnmarshalPacketData(payload channeltypesv2.Payload) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
