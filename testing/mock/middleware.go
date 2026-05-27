package mock

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v11/modules/core/05-port/types"
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

const (
	MockBlockUpgrade = "mockblockupgrade"
)

var _ porttypes.Middleware = (*BlockUpgradeMiddleware)(nil)

// BlockUpgradeMiddleware does not implement the UpgradeableModule interface
type BlockUpgradeMiddleware struct {
	appModule *AppModule
	IBCApp    *IBCApp // base application of an IBC middleware stack
}

// NewIBCModule creates a new IBCModule given the underlying mock IBC application and scopedKeeper.
func NewBlockUpgradeMiddleware(appModule *AppModule, app *IBCApp) BlockUpgradeMiddleware {
	_ = "STUB: not implemented"
	return *new(BlockUpgradeMiddleware)
}

// OnChanOpenInit implements the IBCModule interface.
func (im BlockUpgradeMiddleware) OnChanOpenInit(
	ctx sdk.Context, order channeltypes.Order, connectionHops []string, portID string,
	channelID string, counterparty channeltypes.Counterparty, version string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// OnChanOpenTry implements the IBCModule interface.
func (im BlockUpgradeMiddleware) OnChanOpenTry(
	ctx sdk.Context, order channeltypes.Order, connectionHops []string, portID string,
	channelID string, counterparty channeltypes.Counterparty, counterpartyVersion string,
) (version string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// OnChanOpenAck implements the IBCModule interface.
func (im BlockUpgradeMiddleware) OnChanOpenAck(ctx sdk.Context, portID string, channelID string, counterpartyChannelID string, counterpartyVersion string) error {
	_ = "STUB: not implemented"
	return nil
}

// OnChanOpenConfirm implements the IBCModule interface.
func (im BlockUpgradeMiddleware) OnChanOpenConfirm(ctx sdk.Context, portID, channelID string) error {
	_ = "STUB: not implemented"
	return nil
}

// OnChanCloseInit implements the IBCModule interface.
func (im BlockUpgradeMiddleware) OnChanCloseInit(ctx sdk.Context, portID, channelID string) error {
	_ = "STUB: not implemented"
	return nil
}

// OnChanCloseConfirm implements the IBCModule interface.
func (im BlockUpgradeMiddleware) OnChanCloseConfirm(ctx sdk.Context, portID, channelID string) error {
	_ = "STUB: not implemented"
	return nil
}

// OnRecvPacket implements the IBCModule interface.
func (im BlockUpgradeMiddleware) OnRecvPacket(ctx sdk.Context, channelVersion string, packet channeltypes.Packet, relayer sdk.AccAddress) exported.Acknowledgement {
	_ = "STUB: not implemented"
	return *new(exported.Acknowledgement)
}

// OnAcknowledgementPacket implements the IBCModule interface.
func (im BlockUpgradeMiddleware) OnAcknowledgementPacket(ctx sdk.Context, channelVersion string, packet channeltypes.Packet, acknowledgement []byte, relayer sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}

// OnTimeoutPacket implements the IBCModule interface.
func (im BlockUpgradeMiddleware) OnTimeoutPacket(ctx sdk.Context, channelVersion string, packet channeltypes.Packet, relayer sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}

// SendPacket implements the ICS4 Wrapper interface
func (BlockUpgradeMiddleware) SendPacket(
	ctx sdk.Context,
	sourcePort string,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
	data []byte,
) (uint64, error) {
	_ = "STUB: not implemented"

	// WriteAcknowledgement implements the ICS4 Wrapper interface
	return 0, nil
}

func (BlockUpgradeMiddleware) WriteAcknowledgement(
	ctx sdk.Context,
	packet exported.PacketI,
	ack exported.Acknowledgement,
) error {
	_ = "STUB: not implemented"

	// GetAppVersion returns the application version of the underlying application
	return nil
}

func (BlockUpgradeMiddleware) GetAppVersion(ctx sdk.Context, portID, channelID string) (string, bool) {
	_ = "STUB: not implemented"
	return "",

		// SetICS4Wrapper sets the ICS4Wrapper. This function may be used after
		// the module's initialization to set the middleware which is above this
		// module in the IBC application stack.
		false
}

func (BlockUpgradeMiddleware) SetICS4Wrapper(wrapper porttypes.ICS4Wrapper) {
	_ = "STUB: not implemented"

	// SetUnderlyingApplication sets the underlying application of the middleware.
	return
}

func (BlockUpgradeMiddleware) SetUnderlyingApplication(app porttypes.IBCModule) {
	_ = "STUB: not implemented"
	return
}
