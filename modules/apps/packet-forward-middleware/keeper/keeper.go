package keeper

import (
	"time"

	"github.com/hashicorp/go-metrics"

	"cosmossdk.io/core/address"
	corestore "cosmossdk.io/core/store"
	"cosmossdk.io/log/v2"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/apps/packet-forward-middleware/types"
	transfertypes "github.com/cosmos/ibc-go/v11/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v11/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var (
	// DefaultTransferPacketTimeoutHeight is the timeout height following IBC defaults
	DefaultTransferPacketTimeoutHeight = clienttypes.NewHeight(0, 0)

	// DefaultForwardTransferPacketTimeoutTimestamp is the timeout timestamp following IBC defaults
	DefaultForwardTransferPacketTimeoutTimestamp = time.Duration(10) * time.Minute
)

// Keeper defines the packet forward middleware keeper
type Keeper struct {
	storeService corestore.KVStoreService
	cdc          codec.BinaryCodec
	addressCodec address.Codec

	transferKeeper types.TransferKeeper
	channelKeeper  types.ChannelKeeper
	bankKeeper     types.BankKeeper
	ics4Wrapper    porttypes.ICS4Wrapper

	// the address capable of executing a MsgUpdateParams message. Typically, this
	// should be the x/gov module account.
	authority string
}

// NewKeeper creates a new forward Keeper instance
func NewKeeper(cdc codec.BinaryCodec, addressCodec address.Codec, storeService corestore.KVStoreService, transferKeeper types.TransferKeeper, channelKeeper types.ChannelKeeper, bankKeeper types.BankKeeper, authority string,
) *Keeper {
	_ = "STUB: not implemented"
	return nil
}

// Defaults to using the channel keeper as the ICS4Wrapper
// This can be overridden later with WithICS4Wrapper (e.g. by the middleware stack wiring)

// WithICS4Wrapper sets the ICS4Wrapper for the keeper.
func (k *Keeper) WithICS4Wrapper(ics4Wrapper porttypes.ICS4Wrapper) {
	_ = "STUB: not implemented"
	return
}

// GetAuthority returns the module's authority.
func (k *Keeper) GetAuthority() string {
	_ = "STUB: not implemented"

	// GetAddressCodec returns the address codec used by the keeper.
	return ""
}

func (k *Keeper) GetAddressCodec() address.Codec {
	_ = "STUB: not implemented"
	return *

	// SetICS4Wrapper sets the ICS4 wrapper.
	new(address.Codec)
}

func (k *Keeper) SetICS4Wrapper(ics4Wrapper porttypes.ICS4Wrapper) {
	_ = "STUB: not implemented"
	return
}

// ICS4Wrapper gets the ICS4 Wrapper for PFM.
func (k *Keeper) ICS4Wrapper() porttypes.ICS4Wrapper {
	_ = "STUB: not implemented"
	return *

	// Logger returns a module-specific logger.
	new(porttypes.ICS4Wrapper)
}

func (*Keeper) Logger(ctx sdk.Context) log.Logger {
	_ = "STUB: not implemented"
	return *new(log.Logger)
}

func (k *Keeper) WriteAcknowledgementForForwardedPacket(ctx sdk.Context, packet channeltypes.Packet, transferDetail transfertypes.InternalTransferRepresentation, inFlightPacket *types.InFlightPacket, ack channeltypes.Acknowledgement) error {
	_ = "STUB: not implemented"
	// Lookup module by channel capability
	return nil
}

// For forwarded packets, the funds were moved into an escrow account if the denom originated on this chain.
// On an ack error or timeout on a forwarded packet, the funds in the escrow account
// should be moved to the other escrow account on the other side or burned.

// Sender chain is source

// funds were moved to escrow account for transfer, so they need to either:
// - move to the other escrow account, in the case of native denom
// - burn

// transfer funds from escrow account for forwarded packet to escrow account going back for refund.

// Transfer the coins from the escrow account to the module account and burn them.

// NOTE: should not happen as the module account was
// retrieved on the step above and it has enough balance
// to burn.

// Funds in the escrow account were burned,
// so on a timeout or acknowledgement error we need to mint the funds back to the escrow account.

// unescrowToken will update the total escrow by deducting the unescrowed token
// from the current total escrow.
func (k *Keeper) unescrowToken(ctx sdk.Context, token sdk.Coin) { _ = "STUB: not implemented"; return }

func (k *Keeper) ForwardTransferPacket(ctx sdk.Context, inFlightPacket *types.InFlightPacket, srcPacket channeltypes.Packet, srcPacketSender, receiver string, metadata types.ForwardMetadata, token sdk.Coin, maxRetries uint8, timeoutDelta time.Duration, labels []metrics.Label) error {
	_ = "STUB: not implemented"

	// set memo for next transfer with next from this transfer.
	return nil
}

// send tokens to destination

// Store the following information in keeper:
// key - information about forwarded packet: src_channel (parsedReceiver.Channel), src_port (parsedReceiver.Port), sequence
// value - information about original packet for refunding if necessary: retries, srcPacketSender, srcPacket.DestinationChannel, srcPacket.DestinationPort

// TimeoutShouldRetry returns inFlightPacket and no error if retry should be attempted. Error is returned if IBC refund should occur.
func (k *Keeper) TimeoutShouldRetry(ctx sdk.Context, packet channeltypes.Packet) (*types.InFlightPacket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Not a forwarded packet. Ignore.

// nolint:nilnil

func (k *Keeper) RetryTimeout(ctx sdk.Context, channel, port string, transferDetail transfertypes.InternalTransferRepresentation, inFlightPacket *types.InFlightPacket) error {
	_ = "STUB: not implemented"
	// send transfer again
	return nil
}

// srcPacket and srcPacketSender are empty because inFlightPacket is non-nil.

func (k *Keeper) SetInflightPacket(ctx sdk.Context, channel, port string, sequence uint64, packet *types.InFlightPacket) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Keeper) GetInflightPacket(ctx sdk.Context, packet channeltypes.Packet) (*types.InFlightPacket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// nolint:nilnil

func (k *Keeper) RemoveInFlightPacket(ctx sdk.Context, packet channeltypes.Packet) {
	_ = "STUB: not implemented"
	return
}

// not a forwarded packet, ignore.

// done with packet key now, delete.

// SendPacket wraps IBC ChannelKeeper's SendPacket function
func (k *Keeper) SendPacket(ctx sdk.Context, sourcePort, sourceChannel string, timeoutHeight clienttypes.Height, timeoutTimestamp uint64, data []byte) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// WriteAcknowledgement wraps IBC ICS4Wrapper WriteAcknowledgement function.
// ICS29 WriteAcknowledgement is used for asynchronous acknowledgements.
func (k *Keeper) WriteAcknowledgement(ctx sdk.Context, packet ibcexported.PacketI, acknowledgement ibcexported.Acknowledgement) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteAcknowledgement wraps IBC ICS4Wrapper GetAppVersion function.
func (k *Keeper) GetAppVersion(ctx sdk.Context, portID, channelID string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// LookupModuleByChannel wraps ChannelKeeper LookupModuleByChannel function.
func (k *Keeper) GetChannel(ctx sdk.Context, portID, channelID string) (channeltypes.Channel, bool) {
	_ = "STUB: not implemented"
	return *new(channeltypes.Channel), false
}
