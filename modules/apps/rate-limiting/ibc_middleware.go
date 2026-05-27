package ratelimiting

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/apps/rate-limiting/keeper"
	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v11/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var (
	_ porttypes.Middleware              = (*IBCMiddleware)(nil)
	_ porttypes.PacketUnmarshalerModule = (*IBCMiddleware)(nil)
)

// IBCMiddleware implements the ICS26 callbacks for the rate-limiting middleware.
type IBCMiddleware struct {
	app    porttypes.PacketUnmarshalerModule
	keeper *keeper.Keeper
}

// NewIBCMiddleware creates a new IBCMiddleware given the keeper, underlying application, and channel keeper.
func NewIBCMiddleware(k *keeper.Keeper) *IBCMiddleware { _ = "STUB: not implemented"; return nil }

// OnChanOpenInit implements the IBCMiddleware interface. Call underlying app's OnChanOpenInit.
func (im *IBCMiddleware) OnChanOpenInit(ctx sdk.Context, order channeltypes.Order, connectionHops []string, portID string, channelID string, counterparty channeltypes.Counterparty, version string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// OnChanOpenTry implements the IBCMiddleware interface. Call underlying app's OnChanOpenTry.
func (im *IBCMiddleware) OnChanOpenTry(ctx sdk.Context, order channeltypes.Order, connectionHops []string, portID, channelID string, counterparty channeltypes.Counterparty, counterpartyVersion string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// OnChanOpenAck implements the IBCMiddleware interface. Call underlying app's OnChanOpenAck.
func (im *IBCMiddleware) OnChanOpenAck(ctx sdk.Context, portID, channelID string, counterpartyChannelID string, counterpartyVersion string) error {
	_ = "STUB: not implemented"
	return nil
}

// OnChanOpenConfirm implements the IBCMiddleware interface. Call underlying app's OnChanOpenConfirm.
func (im *IBCMiddleware) OnChanOpenConfirm(ctx sdk.Context, portID, channelID string) error {
	_ = "STUB: not implemented"
	return nil
}

// OnChanCloseInit implements the IBCMiddleware interface. Call underlying app's OnChanCloseInit.
func (im *IBCMiddleware) OnChanCloseInit(ctx sdk.Context, portID, channelID string) error {
	_ = "STUB: not implemented"
	return nil
}

// OnChanCloseConfirm implements the IBCMiddleware interface. Call underlying app's OnChanCloseConfirm.
func (im *IBCMiddleware) OnChanCloseConfirm(ctx sdk.Context, portID, channelID string) error {
	_ = "STUB: not implemented"
	return nil
}

// OnRecvPacket implements the IBCMiddleware interface.
// Rate limits the incoming packet. If the packet is allowed, call underlying app's OnRecvPacket.
func (im *IBCMiddleware) OnRecvPacket(ctx sdk.Context, channelVersion string, packet channeltypes.Packet, relayer sdk.AccAddress) ibcexported.Acknowledgement {
	_ = "STUB: not implemented"
	return *new(ibcexported.Acknowledgement)
}

// If the packet was not rate-limited, pass it down to the underlying app's OnRecvPacket callback

// OnAcknowledgementPacket implements the IBCMiddleware interface.
// If the acknowledgement was an error, revert the outflow amount.
// Then, call underlying app's OnAcknowledgementPacket.
func (im *IBCMiddleware) OnAcknowledgementPacket(ctx sdk.Context, channelVersion string, packet channeltypes.Packet, acknowledgement []byte, relayer sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}

// OnTimeoutPacket implements the IBCMiddleware interface.
// Revert the outflow amount. Then, call underlying app's OnTimeoutPacket.
func (im *IBCMiddleware) OnTimeoutPacket(ctx sdk.Context, channelVersion string, packet channeltypes.Packet, relayer sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}

// SendPacket implements the ICS4 Wrapper interface.
// It calls the keeper's SendRateLimitedPacket function first to check the rate limit.
// If the packet is allowed, it then calls the underlying ICS4Wrapper SendPacket.
func (im *IBCMiddleware) SendPacket(ctx sdk.Context, sourcePort string, sourceChannel string, timeoutHeight clienttypes.Height, timeoutTimestamp uint64, data []byte) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// WriteAcknowledgement implements the ICS4 Wrapper interface.
// It calls the underlying ICS4Wrapper.
func (im *IBCMiddleware) WriteAcknowledgement(ctx sdk.Context, packet ibcexported.PacketI, ack ibcexported.Acknowledgement) error {
	_ = "STUB: not implemented"
	return nil
}

// GetAppVersion implements the ICS4 Wrapper interface.
// It calls the underlying ICS4Wrapper.
func (im *IBCMiddleware) GetAppVersion(ctx sdk.Context, portID, channelID string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// UnmarshalPacketData implements the PacketDataUnmarshaler interface.
// It defers to the underlying app to unmarshal the packet data.
func (im *IBCMiddleware) UnmarshalPacketData(ctx sdk.Context, portID string, channelID string, bz []byte) (any, string, error) {
	_ = "STUB: not implemented"
	return *new(any), "", nil
}

func (im *IBCMiddleware) SetICS4Wrapper(wrapper porttypes.ICS4Wrapper) {
	_ = "STUB: not implemented"
	return
}

func (im *IBCMiddleware) SetUnderlyingApplication(app porttypes.IBCModule) {
	_ = "STUB: not implemented"
	return
}

// the underlying application must implement the PacketUnmarshalerModule interface
