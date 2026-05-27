package host

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/apps/27-interchain-accounts/host/keeper"
	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v11/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var (
	_ porttypes.IBCModule             = (*IBCModule)(nil)
	_ porttypes.PacketDataUnmarshaler = (*IBCModule)(nil)
)

// IBCModule implements the ICS26 interface for interchain accounts host chains
type IBCModule struct {
	keeper *keeper.Keeper
}

// NewIBCModule creates a new IBCModule given the associated keeper
func NewIBCModule(k *keeper.Keeper) *IBCModule { _ = "STUB: not implemented"; return nil }

// OnChanOpenInit implements the IBCModule interface
func (*IBCModule) OnChanOpenInit(
	_ sdk.Context,
	_ channeltypes.Order,
	_ []string,
	_ string,
	_ string,
	_ channeltypes.Counterparty,
	_ string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// OnChanOpenTry implements the IBCModule interface
func (im *IBCModule) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID,
	channelID string,
	counterparty channeltypes.Counterparty,
	counterpartyVersion string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// OnChanOpenAck implements the IBCModule interface
func (*IBCModule) OnChanOpenAck(
	_ sdk.Context,
	_,
	_ string,
	_ string,
	_ string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// OnChanOpenConfirm implements the IBCModule interface
func (im *IBCModule) OnChanOpenConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// OnChanCloseInit implements the IBCModule interface
func (*IBCModule) OnChanCloseInit(
	_ sdk.Context,
	_ string,
	_ string,
) error {
	_ = "STUB: not implemented"
	// Disallow user-initiated channel closing for interchain account channels
	return nil
}

// OnChanCloseConfirm implements the IBCModule interface
func (im *IBCModule) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// OnRecvPacket implements the IBCModule interface
func (im *IBCModule) OnRecvPacket(
	ctx sdk.Context,
	_ string,
	packet channeltypes.Packet,
	_ sdk.AccAddress,
) ibcexported.Acknowledgement {
	_ = "STUB: not implemented"
	return *new(ibcexported.Acknowledgement)
}

// Emit an event indicating a successful or failed acknowledgement.

// NOTE: acknowledgement will be written synchronously during IBC handler execution.

// OnAcknowledgementPacket implements the IBCModule interface
func (*IBCModule) OnAcknowledgementPacket(
	_ sdk.Context,
	_ string,
	_ channeltypes.Packet,
	_ []byte,
	_ sdk.AccAddress,
) error {
	_ = "STUB: not implemented"
	return nil
}

// OnTimeoutPacket implements the IBCModule interface
func (*IBCModule) OnTimeoutPacket(
	_ sdk.Context,
	_ string,
	_ channeltypes.Packet,
	_ sdk.AccAddress,
) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalPacketData attempts to unmarshal the provided packet data bytes
// into an InterchainAccountPacketData. This function implements the optional
// PacketDataUnmarshaler interface required for ADR 008 support.
func (im *IBCModule) UnmarshalPacketData(ctx sdk.Context, portID string, channelID string, bz []byte) (any, string, error) {
	_ = "STUB: not implemented"
	return *new(any), "", nil
}

// SetICS4Wrapper sets the ICS4Wrapper for the IBCModule.
func (im *IBCModule) SetICS4Wrapper(wrapper porttypes.ICS4Wrapper) {
	_ = "STUB: not implemented"
	return
}
