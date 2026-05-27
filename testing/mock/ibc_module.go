package mock

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v11/modules/core/05-port/types"
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var (
	_ porttypes.IBCModule             = (*IBCModule)(nil)
	_ porttypes.PacketDataUnmarshaler = (*IBCModule)(nil)
)

// applicationCallbackError is a custom error type that will be unique for testing purposes.
type applicationCallbackError struct{}

func (applicationCallbackError) Error() string { _ = "STUB: not implemented"; return "" }

// IBCModule implements the ICS26 callbacks for testing/mock.
type IBCModule struct {
	appModule *AppModule
	IBCApp    *IBCApp // base application of an IBC middleware stack
}

// NewIBCModule creates a new IBCModule given the underlying mock IBC application and scopedKeeper.
func NewIBCModule(appModule *AppModule, app *IBCApp) *IBCModule {
	_ = "STUB: not implemented"
	return nil
}

// OnChanOpenInit implements the IBCModule interface.
func (im IBCModule) OnChanOpenInit(
	ctx sdk.Context, order channeltypes.Order, connectionHops []string, portID string,
	channelID string, counterparty channeltypes.Counterparty, version string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// OnChanOpenTry implements the IBCModule interface.
func (im IBCModule) OnChanOpenTry(
	ctx sdk.Context, order channeltypes.Order, connectionHops []string, portID string,
	channelID string, counterparty channeltypes.Counterparty, counterpartyVersion string,
) (version string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// OnChanOpenAck implements the IBCModule interface.
func (im IBCModule) OnChanOpenAck(ctx sdk.Context, portID string, channelID string, counterpartyChannelID string, counterpartyVersion string) error {
	_ = "STUB: not implemented"
	return nil
}

// OnChanOpenConfirm implements the IBCModule interface.
func (im IBCModule) OnChanOpenConfirm(ctx sdk.Context, portID, channelID string) error {
	_ = "STUB: not implemented"
	return nil
}

// OnChanCloseInit implements the IBCModule interface.
func (im IBCModule) OnChanCloseInit(ctx sdk.Context, portID, channelID string) error {
	_ = "STUB: not implemented"
	return nil
}

// OnChanCloseConfirm implements the IBCModule interface.
func (im IBCModule) OnChanCloseConfirm(ctx sdk.Context, portID, channelID string) error {
	_ = "STUB: not implemented"
	return nil
}

// OnRecvPacket implements the IBCModule interface.
func (im IBCModule) OnRecvPacket(ctx sdk.Context, channelVersion string, packet channeltypes.Packet, relayer sdk.AccAddress) exported.Acknowledgement {
	_ = "STUB: not implemented"
	return *new(exported.Acknowledgement)
}

// OnAcknowledgementPacket implements the IBCModule interface.
func (im IBCModule) OnAcknowledgementPacket(ctx sdk.Context, channelVersion string, packet channeltypes.Packet, acknowledgement []byte, relayer sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}

// OnTimeoutPacket implements the IBCModule interface.
func (im IBCModule) OnTimeoutPacket(ctx sdk.Context, channelVersion string, packet channeltypes.Packet, relayer sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalPacketData returns the MockPacketData. This function implements the optional
// PacketDataUnmarshaler interface required for ADR 008 support.
func (IBCModule) UnmarshalPacketData(ctx sdk.Context, portID string, channelID string, bz []byte) (any, string, error) {
	_ = "STUB: not implemented"
	return *new(any), "", nil
}

func (*IBCModule) SetICS4Wrapper(wrapper porttypes.ICS4Wrapper) { _ = "STUB: not implemented"; return }
