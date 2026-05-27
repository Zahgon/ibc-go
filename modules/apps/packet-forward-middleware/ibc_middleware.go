package packetforward

import (
	"time"

	coreaddress "cosmossdk.io/core/address"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/apps/packet-forward-middleware/keeper"
	transfertypes "github.com/cosmos/ibc-go/v11/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v11/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var (
	_ porttypes.Middleware              = &IBCMiddleware{}
	_ porttypes.PacketUnmarshalerModule = &IBCMiddleware{}
)

// IBCMiddleware implements the ICS26 callbacks for the forward middleware given the
// forward keeper and the underlying application.
type IBCMiddleware struct {
	app    porttypes.PacketUnmarshalerModule
	keeper *keeper.Keeper

	retriesOnTimeout uint8
	forwardTimeout   time.Duration
}

// NewIBCMiddleware creates a new IBCMiddleware given the keeper and underlying application.
func NewIBCMiddleware(k *keeper.Keeper, retriesOnTimeout uint8, forwardTimeout time.Duration) *IBCMiddleware {
	_ = "STUB: not implemented"
	return nil
}

// OnChanOpenInit implements the IBCModule interface.
func (im *IBCMiddleware) OnChanOpenInit(ctx sdk.Context, order channeltypes.Order, connectionHops []string, portID string, channelID string, counterparty channeltypes.Counterparty, version string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// OnChanOpenTry implements the IBCModule interface.
func (im *IBCMiddleware) OnChanOpenTry(ctx sdk.Context, order channeltypes.Order, connectionHops []string, portID, channelID string, counterparty channeltypes.Counterparty, counterpartyVersion string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// OnChanOpenAck implements the IBCModule interface.
func (im *IBCMiddleware) OnChanOpenAck(ctx sdk.Context, portID, channelID string, counterpartyChannelID string, counterpartyVersion string) error {
	_ = "STUB: not implemented"
	return nil
}

// OnChanOpenConfirm implements the IBCModule interface.
func (im *IBCMiddleware) OnChanOpenConfirm(ctx sdk.Context, portID, channelID string) error {
	_ = "STUB: not implemented"
	return nil
}

// OnChanCloseInit implements the IBCModule interface.
func (im *IBCMiddleware) OnChanCloseInit(ctx sdk.Context, portID, channelID string) error {
	_ = "STUB: not implemented"
	return nil
}

// OnChanCloseConfirm implements the IBCModule interface.
func (im *IBCMiddleware) OnChanCloseConfirm(ctx sdk.Context, portID, channelID string) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalPacketData implements PacketDataUnmarshaler.
func (im *IBCMiddleware) UnmarshalPacketData(ctx sdk.Context, portID string, channelID string, bz []byte) (any, string, error) {
	_ = "STUB: not implemented"
	return *new(any), "", nil
}

func getDenomForThisChain(port, channel, counterpartyPort, counterpartyChannel string, denom transfertypes.Denom) string {
	_ = "STUB: not implemented"
	return ""
}

// unwind denom

// denom is now unwound back to native denom

// denom is still IBC denom

// append port and channel from this chain to denom

// GetReceiver returns the receiver address for a given channel and original sender.
// it overrides the receiver address to be a hash of the channel/origSender so that
// the receiver address is deterministic and can be used to identify the sender on the
// initial chain.
func GetReceiver(addressCodec coreaddress.Codec, channel string, originalSender string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// OnRecvPacket checks the memo field on this packet and if the metadata inside's root key indicates this packet
// should be handled by the swap middleware it attempts to perform a swap. If the swap is successful
// the underlying application's OnRecvPacket callback is invoked, an ack error is returned otherwise.
func (im *IBCMiddleware) OnRecvPacket(ctx sdk.Context, channelVersion string, packet channeltypes.Packet, relayer sdk.AccAddress) ibcexported.Acknowledgement {
	_ = "STUB: not implemented"
	return *new(ibcexported.Acknowledgement)
}

// not a packet that should be forwarded

// override the receiver so that senders cannot move funds through arbitrary addresses.

// if this packet's token denom is already the base denom for some native token on this chain,
// we do not need to do any further composition of the denom before forwarding the packet

// returning nil ack will prevent WriteAcknowledgement from occurring for forwarded packet.
// This is intentional so that the acknowledgement will be written later based on the ack/timeout of the forwarded packet.

// receiveFunds receives funds from the packet into the override receiver
// address and returns an error if the funds cannot be received.
func (im *IBCMiddleware) receiveFunds(ctx sdk.Context, channelVersion string, packet channeltypes.Packet, data transfertypes.FungibleTokenPacketData, overrideReceiver string, relayer sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}

// Memo explicitly emptied.

// Override data.

// OnAcknowledgementPacket implements the IBCModule interface.
func (im *IBCMiddleware) OnAcknowledgementPacket(ctx sdk.Context, channelVersion string, packet channeltypes.Packet, acknowledgement []byte, relayer sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}

// this is a forwarded packet, so override handling to avoid refund from being processed.

// OnTimeoutPacket implements the IBCModule interface.
func (im *IBCMiddleware) OnTimeoutPacket(ctx sdk.Context, channelVersion string, packet channeltypes.Packet, relayer sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}

// this is a forwarded packet, so override handling to avoid refund from being processed on this chain.
// WriteAcknowledgement with proxied ack to return success/fail to previous chain.

// timeout should be retried. In order to do that, we need to handle this timeout to refund on this chain first.

// SendPacket implements the ICS4 Wrapper interface.
func (im *IBCMiddleware) SendPacket(ctx sdk.Context, sourcePort, sourceChannel string, timeoutHeight clienttypes.Height, timeoutTimestamp uint64, data []byte) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// WriteAcknowledgement implements the ICS4 Wrapper interface.
func (im *IBCMiddleware) WriteAcknowledgement(ctx sdk.Context, packet ibcexported.PacketI, ack ibcexported.Acknowledgement) error {
	_ = "STUB: not implemented"
	return nil
}

func (im *IBCMiddleware) GetAppVersion(ctx sdk.Context, portID, channelID string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
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
