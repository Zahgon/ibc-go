package ibccallbacks

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/apps/callbacks/types"
	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v11/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var (
	_ porttypes.Middleware            = (*IBCMiddleware)(nil)
	_ porttypes.PacketDataUnmarshaler = (*IBCMiddleware)(nil)
)

// IBCMiddleware implements the ICS26 callbacks for the ibc-callbacks middleware given
// the underlying application.
type IBCMiddleware struct {
	app         porttypes.PacketUnmarshalerModule
	ics4Wrapper porttypes.ICS4Wrapper

	contractKeeper types.ContractKeeper

	// maxCallbackGas defines the maximum amount of gas that a callback actor can ask the
	// relayer to pay for. If a callback fails due to insufficient gas, the entire tx
	// is reverted if the relayer hadn't provided the minimum(userDefinedGas, maxCallbackGas).
	// If the actor hasn't defined a gas limit, then it is assumed to be the maxCallbackGas.
	maxCallbackGas uint64
}

// NewIBCMiddleware creates a new IBCMiddleware given the keeper and underlying application.
// The underlying application must implement the required callback interfaces.
func NewIBCMiddleware(
	contractKeeper types.ContractKeeper, maxCallbackGas uint64,
) *IBCMiddleware {
	_ = "STUB: not implemented"
	return nil
}

// SetICS4Wrapper sets the ICS4Wrapper. This function may be used after the
// middleware's creation to set the middleware which is above this module in
// the IBC application stack.
func (im *IBCMiddleware) SetICS4Wrapper(wrapper porttypes.ICS4Wrapper) {
	_ = "STUB: not implemented"
	return
}

// SetUnderlyingApplication sets the underlying IBC module. This function may be used after
// the middleware's creation to set the ibc module which is below this middleware.
func (im *IBCMiddleware) SetUnderlyingApplication(app porttypes.IBCModule) {
	_ = "STUB: not implemented"
	return
}

// the underlying application must implement the PacketUnmarshalerModule interface

// GetICS4Wrapper returns the ICS4Wrapper.
func (im *IBCMiddleware) GetICS4Wrapper() porttypes.ICS4Wrapper {
	_ = "STUB: not implemented"
	return *

	// SendPacket implements source callbacks for sending packets.
	// It defers to the underlying application and then calls the contract callback.
	// If the contract callback returns an error, panics, or runs out of gas, then
	// the packet send is rejected.
	new(porttypes.ICS4Wrapper)
}

func (im *IBCMiddleware) SendPacket(
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

// packet is created without destination information present, GetSourceCallbackData does not use these.

// SendPacket is not blocked if the packet does not opt-in to callbacks

// if the packet does opt-in to callbacks but the callback data is malformed,
// then the packet send is rejected.

// contract keeper is allowed to reject the packet send.

// OnAcknowledgementPacket implements source callbacks for acknowledgement packets.
// It defers to the underlying application and then calls the contract callback.
// If the contract callback runs out of gas and may be retried with a higher gas limit then the state changes are
// reverted via a panic.
func (im *IBCMiddleware) OnAcknowledgementPacket(
	ctx sdk.Context,
	channelVersion string,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	_ = "STUB: not implemented"
	// we first call the underlying app to handle the acknowledgement
	return nil
}

// OnAcknowledgementPacket is not blocked if the packet does not opt-in to callbacks

// if the packet does opt-in to callbacks but the callback data is malformed,
// then the packet acknowledgement is rejected.
// This should never occur, since this error is already checked on `SendPacket`

// callback execution errors are not allowed to block the packet lifecycle, they are only used in event emissions

// OnTimeoutPacket implements timeout source callbacks for the ibc-callbacks middleware.
// It defers to the underlying application and then calls the contract callback.
// If the contract callback runs out of gas and may be retried with a higher gas limit then the state changes are
// reverted via a panic.
func (im *IBCMiddleware) OnTimeoutPacket(ctx sdk.Context, channelVersion string, packet channeltypes.Packet, relayer sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}

// OnTimeoutPacket is not blocked if the packet does not opt-in to callbacks

// if the packet does opt-in to callbacks but the callback data is malformed,
// then the packet timeout is rejected.
// This should never occur, since this error is already checked on `SendPacket`

// callback execution errors are not allowed to block the packet lifecycle, they are only used in event emissions

// OnRecvPacket implements the ReceivePacket destination callbacks for the ibc-callbacks middleware during
// synchronous packet acknowledgement.
// It defers to the underlying application and then calls the contract callback.
// If the contract callback runs out of gas and may be retried with a higher gas limit then the state changes are
// reverted via a panic.
func (im *IBCMiddleware) OnRecvPacket(ctx sdk.Context, channelVersion string, packet channeltypes.Packet, relayer sdk.AccAddress) ibcexported.Acknowledgement {
	_ = "STUB: not implemented"
	return *new(ibcexported.Acknowledgement)
}

// if ack is nil (asynchronous acknowledgements), then the callback will be handled in WriteAcknowledgement
// if ack is not successful, all state changes are reverted. If a packet cannot be received, then there is
// no need to execute a callback on the receiving chain.

// OnRecvPacket is not blocked if the packet does not opt-in to callbacks

// if the packet does opt-in to callbacks but the callback data is malformed,
// then the packet receive is rejected.

// callback execution errors in RecvPacket are allowed to write an error acknowledgement
// in this case, the receive logic of the underlying app is reverted
// and the error acknowledgement is processed on the sending chain
// Thus the sending application MUST be capable of processing the standard channel acknowledgement

// WriteAcknowledgement implements the ReceivePacket destination callbacks for the ibc-callbacks middleware
// during asynchronous packet acknowledgement.
// It defers to the underlying application and then calls the contract callback.
// If the contract callback runs out of gas and may be retried with a higher gas limit then the state changes are
// reverted via a panic.
func (im *IBCMiddleware) WriteAcknowledgement(
	ctx sdk.Context,
	packet ibcexported.PacketI,
	ack ibcexported.Acknowledgement,
) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteAcknowledgement is not blocked if the packet does not opt-in to callbacks

// This should never occur, since this error is already checked on `OnRecvPacket`

// callback execution errors are not allowed to block the packet lifecycle, they are only used in event emissions

// OnChanOpenInit defers to the underlying application
func (im *IBCMiddleware) OnChanOpenInit(
	ctx sdk.Context,
	channelOrdering channeltypes.Order,
	connectionHops []string,
	portID,
	channelID string,
	counterparty channeltypes.Counterparty,
	version string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// OnChanOpenTry defers to the underlying application
func (im *IBCMiddleware) OnChanOpenTry(
	ctx sdk.Context,
	channelOrdering channeltypes.Order,
	connectionHops []string, portID,
	channelID string,
	counterparty channeltypes.Counterparty,
	counterpartyVersion string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// OnChanOpenAck defers to the underlying application
func (im *IBCMiddleware) OnChanOpenAck(
	ctx sdk.Context,
	portID,
	channelID,
	counterpartyChannelID,
	counterpartyVersion string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// OnChanOpenConfirm defers to the underlying application
func (im *IBCMiddleware) OnChanOpenConfirm(ctx sdk.Context, portID, channelID string) error {
	_ = "STUB: not implemented"
	return nil
}

// OnChanCloseInit defers to the underlying application
func (im *IBCMiddleware) OnChanCloseInit(ctx sdk.Context, portID, channelID string) error {
	_ = "STUB: not implemented"
	return nil
}

// OnChanCloseConfirm defers to the underlying application
func (im *IBCMiddleware) OnChanCloseConfirm(ctx sdk.Context, portID, channelID string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetAppVersion implements the ICS4Wrapper interface. Callbacks has no version,
// so the call is deferred to the underlying application.
func (im *IBCMiddleware) GetAppVersion(ctx sdk.Context, portID, channelID string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// UnmarshalPacketData defers to the underlying app to unmarshal the packet data.
// This function implements the optional PacketDataUnmarshaler interface.
func (im *IBCMiddleware) UnmarshalPacketData(ctx sdk.Context, portID string, channelID string, bz []byte) (any, string, error) {
	_ = "STUB: not implemented"
	return *new(any), "", nil
}
