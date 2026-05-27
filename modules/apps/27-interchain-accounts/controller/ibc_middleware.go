package controller

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/apps/27-interchain-accounts/controller/keeper"
	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v11/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var (
	_ porttypes.Middleware            = (*IBCMiddleware)(nil)
	_ porttypes.PacketDataUnmarshaler = (*IBCMiddleware)(nil)
)

// IBCMiddleware implements the ICS26 callbacks for the controller middleware given the
// ICA controller keeper and the underlying application.
type IBCMiddleware struct {
	app    porttypes.IBCModule
	keeper *keeper.Keeper
}

// NewIBCMiddleware creates a new IBCMiddleware given the associated keeper.
// The underlying application is set to nil and authentication is assumed to
// be performed by a Cosmos SDK module that sends messages to controller message server.
func NewIBCMiddleware(k *keeper.Keeper) *IBCMiddleware { _ = "STUB: not implemented"; return nil }

// NewIBCMiddlewareWithAuth creates a new IBCMiddleware given the associated keeper and underlying application
func NewIBCMiddlewareWithAuth(app porttypes.IBCModule, k *keeper.Keeper) *IBCMiddleware {
	_ = "STUB: not implemented"
	return nil
}

// SetUnderlyingApplication sets the underlying application for the middleware.
func (im *IBCMiddleware) SetUnderlyingApplication(app porttypes.IBCModule) {
	_ = "STUB: not implemented"
	return
}

// SetICS4Wrapper sets the ICS4Wrapper for the middleware.
func (im *IBCMiddleware) SetICS4Wrapper(ics4Wrapper porttypes.ICS4Wrapper) {
	_ = "STUB: not implemented"
	return
}

// OnChanOpenInit implements the IBCMiddleware interface
//
// Interchain Accounts is implemented to act as middleware for connected authentication modules on
// the controller side. The connected modules may not change the controller side portID or
// version. They will be allowed to perform custom logic without changing
// the parameters stored within a channel struct.
func (im *IBCMiddleware) OnChanOpenInit(
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

// call underlying app's OnChanOpenInit callback with the passed in version
// the version returned is discarded as the ica-auth module does not have permission to edit the version string.
// ics27 will always return the version string containing the Metadata struct which is created during the `RegisterInterchainAccount` call.

// OnChanOpenTry implements the IBCMiddleware interface
func (*IBCMiddleware) OnChanOpenTry(
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

// OnChanOpenAck implements the IBCMiddleware interface
//
// Interchain Accounts is implemented to act as middleware for connected authentication modules on
// the controller side. The connected modules may not change the portID or
// version. They will be allowed to perform custom logic without changing
// the parameters stored within a channel struct.
func (im *IBCMiddleware) OnChanOpenAck(
	ctx sdk.Context,
	portID,
	channelID string,
	counterpartyChannelID string,
	counterpartyVersion string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// call underlying app's OnChanOpenAck callback with the counterparty app version.

// OnChanOpenConfirm implements the IBCMiddleware interface
func (*IBCMiddleware) OnChanOpenConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// OnChanCloseInit implements the IBCMiddleware interface
func (*IBCMiddleware) OnChanCloseInit(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	_ = "STUB: not implemented"
	// Disallow user-initiated channel closing for interchain account channels
	return nil
}

// OnChanCloseConfirm implements the IBCMiddleware interface
func (im *IBCMiddleware) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// OnRecvPacket implements the IBCMiddleware interface
func (*IBCMiddleware) OnRecvPacket(
	ctx sdk.Context,
	_ string,
	packet channeltypes.Packet,
	_ sdk.AccAddress,
) ibcexported.Acknowledgement {
	_ = "STUB: not implemented"
	return *new(ibcexported.Acknowledgement)
}

// OnAcknowledgementPacket implements the IBCMiddleware interface
func (im *IBCMiddleware) OnAcknowledgementPacket(
	ctx sdk.Context,
	channelVersion string,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	_ = "STUB: not implemented"
	return nil
}

// call underlying app's OnAcknowledgementPacket callback.

// OnTimeoutPacket implements the IBCMiddleware interface
func (im *IBCMiddleware) OnTimeoutPacket(
	ctx sdk.Context,
	channelVersion string,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	_ = "STUB: not implemented"
	return nil
}

// SendPacket implements the ICS4 Wrapper interface
func (*IBCMiddleware) SendPacket(
	ctx sdk.Context,
	sourcePort string,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
	data []byte,
) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// WriteAcknowledgement implements the ICS4 Wrapper interface
func (*IBCMiddleware) WriteAcknowledgement(
	ctx sdk.Context,
	packet ibcexported.PacketI,
	ack ibcexported.Acknowledgement,
) error {
	_ = "STUB: not implemented"
	return nil
}

// GetAppVersion returns the interchain accounts metadata.
func (im *IBCMiddleware) GetAppVersion(ctx sdk.Context, portID, channelID string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// UnmarshalPacketData attempts to unmarshal the provided packet data bytes
// into an InterchainAccountPacketData. This function implements the optional
// PacketDataUnmarshaler interface required for ADR 008 support.
func (im *IBCMiddleware) UnmarshalPacketData(ctx sdk.Context, portID string, channelID string, bz []byte) (any, string, error) {
	_ = "STUB: not implemented"
	return *new(any), "", nil
}
