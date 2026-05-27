package transfer

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/apps/transfer/keeper"
	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v11/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var (
	_ porttypes.IBCModule             = (*IBCModule)(nil)
	_ porttypes.PacketDataUnmarshaler = (*IBCModule)(nil)
)

// IBCModule implements the ICS26 interface for transfer given the transfer keeper.
type IBCModule struct {
	keeper *keeper.Keeper
}

// NewIBCModule creates a new IBCModule given the keeper
func NewIBCModule(k *keeper.Keeper) *IBCModule { _ = "STUB: not implemented"; return nil }

// ValidateTransferChannelParams does validation of a newly created transfer channel. A transfer
// channel must be UNORDERED, use the correct port (by default 'transfer'), and use the current
// supported version. Only 2^32 channels are allowed to be created.
func ValidateTransferChannelParams(
	ctx sdk.Context,
	transferkeeper keeper.Keeper,
	order channeltypes.Order,
	portID string,
	channelID string,
) error {
	_ = "STUB: not implemented"
	// NOTE: for escrow address security only 2^32 channels are allowed to be created
	// Issue: https://github.com/cosmos/cosmos-sdk/issues/7737
	return nil
}

// Require portID is the portID transfer module is bound to

// OnChanOpenInit implements the IBCModule interface
func (im IBCModule) OnChanOpenInit(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	counterparty channeltypes.Counterparty,
	version string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// default to latest supported version

// OnChanOpenTry implements the IBCModule interface.
func (im IBCModule) OnChanOpenTry(
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
func (IBCModule) OnChanOpenAck(
	ctx sdk.Context,
	portID,
	channelID string,
	_ string,
	counterpartyVersion string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// OnChanOpenConfirm implements the IBCModule interface
func (IBCModule) OnChanOpenConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	_ = "STUB: not implemented"

	// OnChanCloseInit implements the IBCModule interface
	return nil
}

func (IBCModule) OnChanCloseInit(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	_ = "STUB: not implemented"
	// Disallow user-initiated channel closing for transfer channels
	return nil
}

// OnChanCloseConfirm implements the IBCModule interface
func (IBCModule) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	_ = "STUB: not implemented"

	// OnRecvPacket implements the IBCModule interface. A successful acknowledgement
	// is returned if the packet data is successfully decoded and the receive application
	// logic returns without error.
	return nil
}

func (im IBCModule) OnRecvPacket(
	ctx sdk.Context,
	channelVersion string,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
	_ = "STUB: not implemented"
	return *new(ibcexported.Acknowledgement)
}

// we are explicitly wrapping this emit event call in an anonymous function so that
// the packet data is evaluated after it has been assigned a value.

// NOTE: this needs to set the ackErr variable and not do if ackErr := ... because the ackErr variable is used in the defer function

// NOTE: acknowledgement will be written synchronously during IBC handler execution.

// OnAcknowledgementPacket implements the IBCModule interface
func (im IBCModule) OnAcknowledgementPacket(
	ctx sdk.Context,
	channelVersion string,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	_ = "STUB: not implemented"
	return nil
}

// OnTimeoutPacket implements the IBCModule interface
func (im IBCModule) OnTimeoutPacket(
	ctx sdk.Context,
	channelVersion string,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	_ = "STUB: not implemented"
	return nil
}

// refund tokens

// UnmarshalPacketData attempts to unmarshal the provided packet data bytes
// into a FungibleTokenPacketData. This function implements the optional
// PacketDataUnmarshaler interface required for ADR 008 support.
func (im IBCModule) UnmarshalPacketData(ctx sdk.Context, portID string, channelID string, bz []byte) (any, string, error) {
	_ = "STUB: not implemented"
	return *new(any), "", nil
}

// SetICS4Wrapper sets the ICS4Wrapper. This function may be used after
// the module's initialization to set the middleware which is above this
// module in the IBC application stack.
func (im IBCModule) SetICS4Wrapper(wrapper porttypes.ICS4Wrapper) {
	_ = "STUB: not implemented"
	return
}
