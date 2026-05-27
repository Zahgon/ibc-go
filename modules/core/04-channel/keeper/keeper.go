package keeper

import (
	db "github.com/cosmos/cosmos-db"

	corestore "cosmossdk.io/core/store"
	"cosmossdk.io/log/v2"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	clientv2types "github.com/cosmos/ibc-go/v11/modules/core/02-client/v2/types"
	connectiontypes "github.com/cosmos/ibc-go/v11/modules/core/03-connection/types"
	"github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v11/modules/core/05-port/types"
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var _ porttypes.ICS4Wrapper = (*Keeper)(nil)

// Keeper defines the IBC channel keeper
type Keeper struct {
	// implements gRPC QueryServer interface
	types.QueryServer

	storeService     corestore.KVStoreService
	cdc              codec.BinaryCodec
	clientKeeper     types.ClientKeeper
	connectionKeeper types.ConnectionKeeper

	// V2 Keepers are only used for channel aliasing
	clientKeeperV2  types.ClientKeeperV2
	channelKeeperV2 types.ChannelKeeperV2
}

// NewKeeper creates a new IBC channel Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec,
	storeService corestore.KVStoreService,
	clientKeeper types.ClientKeeper,
	connectionKeeper types.ConnectionKeeper,
	clientKeeperV2 types.ClientKeeperV2,
	channelKeeperV2 types.ChannelKeeperV2,
) *Keeper {
	_ = "STUB: not implemented"
	return nil
}

// Logger returns a module-specific logger.
func (*Keeper) Logger(ctx sdk.Context) log.Logger {
	_ = "STUB: not implemented"
	return *new(log.Logger)
}

// GenerateChannelIdentifier returns the next channel identifier.
func (k *Keeper) GenerateChannelIdentifier(ctx sdk.Context) string {
	_ = "STUB: not implemented"
	return ""
}

// HasChannel true if the channel with the given identifiers exists in state.
func (k *Keeper) HasChannel(ctx sdk.Context, portID, channelID string) bool {
	_ = "STUB: not implemented"
	return false
}

// GetChannel returns a channel with a particular identifier binded to a specific port
func (k *Keeper) GetChannel(ctx sdk.Context, portID, channelID string) (types.Channel, bool) {
	_ = "STUB: not implemented"
	return *new(types.Channel), false
}

// SetChannel sets a channel to the store
func (k *Keeper) SetChannel(ctx sdk.Context, portID, channelID string, channel types.Channel) {
	_ = "STUB: not implemented"
	return
}

// GetAppVersion gets the version for the specified channel.
func (k *Keeper) GetAppVersion(ctx sdk.Context, portID, channelID string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// GetNextChannelSequence gets the next channel sequence from the store.
func (k *Keeper) GetNextChannelSequence(ctx sdk.Context) uint64 {
	_ = "STUB: not implemented"
	return 0
}

// SetNextChannelSequence sets the next channel sequence to the store.
func (k *Keeper) SetNextChannelSequence(ctx sdk.Context, sequence uint64) {
	_ = "STUB: not implemented"
	return
}

// GetNextSequenceSend gets a channel's next send sequence from the store
// NOTE: Even though we are using IBCv1 protocol, we are using the v2 NextSequenceSendKey
// this allows us to use the same identifiers for both v1 and v2 packets without having the sequences
// collide.
// The v2 NextSequenceSendKey does not include the port ID, and only uses the client ID.
// It is safe for us to only use the channel ID as the client ID since channel ids are unique chain identifiers
// in ibc-go.
func (k *Keeper) GetNextSequenceSend(ctx sdk.Context, portID, channelID string) (uint64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// SetNextSequenceSend sets a channel's next send sequence to the store
// NOTE: Even though we are using IBCv1 protocol, we are using the v2 NextSequenceSendKey
// this allows us to use the same identifiers for both v1 and v2 packets without having the sequences
// collide.
// The v2 NextSequenceSendKey does not include the port ID, and only uses the client ID.
// It is safe for us to only use the channel ID as the client ID since channel ids are unique chain identifiers
// in ibc-go.
func (k *Keeper) SetNextSequenceSend(ctx sdk.Context, portID, channelID string, sequence uint64) {
	_ = "STUB: not implemented"
	return
}

// GetNextSequenceRecv gets a channel's next receive sequence from the store
func (k *Keeper) GetNextSequenceRecv(ctx sdk.Context, portID, channelID string) (uint64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// SetNextSequenceRecv sets a channel's next receive sequence to the store
func (k *Keeper) SetNextSequenceRecv(ctx sdk.Context, portID, channelID string, sequence uint64) {
	_ = "STUB: not implemented"
	return
}

// GetNextSequenceAck gets a channel's next ack sequence from the store
func (k *Keeper) GetNextSequenceAck(ctx sdk.Context, portID, channelID string) (uint64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// SetNextSequenceAck sets a channel's next ack sequence to the store
func (k *Keeper) SetNextSequenceAck(ctx sdk.Context, portID, channelID string, sequence uint64) {
	_ = "STUB: not implemented"
	return
}

// GetPacketReceipt gets a packet receipt from the store
func (k *Keeper) GetPacketReceipt(ctx sdk.Context, portID, channelID string, sequence uint64) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// SetPacketReceipt sets an empty packet receipt to the store
func (k *Keeper) SetPacketReceipt(ctx sdk.Context, portID, channelID string, sequence uint64) {
	_ = "STUB: not implemented"
	return
}

// GetPacketCommitment gets the packet commitment hash from the store
func (k *Keeper) GetPacketCommitment(ctx sdk.Context, portID, channelID string, sequence uint64) []byte {
	_ = "STUB: not implemented"
	return nil
}

// HasPacketCommitment returns true if the packet commitment exists
func (k *Keeper) HasPacketCommitment(ctx sdk.Context, portID, channelID string, sequence uint64) bool {
	_ = "STUB: not implemented"
	return false
}

// SetPacketCommitment sets the packet commitment hash to the store
func (k *Keeper) SetPacketCommitment(ctx sdk.Context, portID, channelID string, sequence uint64, commitmentHash []byte) {
	_ = "STUB: not implemented"
	return
}

func (k *Keeper) deletePacketCommitment(ctx sdk.Context, portID, channelID string, sequence uint64) {
	_ = "STUB: not implemented"
	return
}

// SetPacketAcknowledgement sets the packet ack hash to the store
func (k *Keeper) SetPacketAcknowledgement(ctx sdk.Context, portID, channelID string, sequence uint64, ackHash []byte) {
	_ = "STUB: not implemented"
	return
}

// GetPacketAcknowledgement gets the packet ack hash from the store
func (k *Keeper) GetPacketAcknowledgement(ctx sdk.Context, portID, channelID string, sequence uint64) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// HasPacketAcknowledgement check if the packet ack hash is already on the store
func (k *Keeper) HasPacketAcknowledgement(ctx sdk.Context, portID, channelID string, sequence uint64) bool {
	_ = "STUB: not implemented"
	return false
}

// IteratePacketSequence provides an iterator over all send, receive or ack sequences.
// For each sequence, cb will be called. If the cb returns true, the iterator
// will close and stop.
// NOTE: This function will no longer work for NextSequenceSend
func (k *Keeper) IteratePacketSequence(ctx sdk.Context, iterator db.Iterator, cb func(portID, channelID string, sequence uint64) bool) {
	_ = "STUB: not implemented"
	return
}

// return if the key is not a channel key

// GetAllPacketSendSeqs returns all stored next send sequences.
// NOTE: Implemented differently from NextSequenceRecv/Ack since the key format is different
func (k *Keeper) GetAllPacketSendSeqs(ctx sdk.Context) []types.PacketSequence {
	_ = "STUB: not implemented"
	return nil
}

// GetAllPacketRecvSeqs returns all stored next recv sequences.
func (k *Keeper) GetAllPacketRecvSeqs(ctx sdk.Context) []types.PacketSequence {
	_ = "STUB: not implemented"
	return nil
}

// GetAllPacketAckSeqs returns all stored next acknowledgements sequences.
func (k *Keeper) GetAllPacketAckSeqs(ctx sdk.Context) []types.PacketSequence {
	_ = "STUB: not implemented"
	return nil
}

// IteratePacketCommitment provides an iterator over all PacketCommitment objects. For each
// packet commitment, cb will be called. If the cb returns true, the iterator will close
// and stop.
func (k *Keeper) IteratePacketCommitment(ctx sdk.Context, cb func(portID, channelID string, sequence uint64, hash []byte) bool) {
	_ = "STUB: not implemented"
	return
}

// GetAllPacketCommitments returns all stored PacketCommitments objects.
func (k *Keeper) GetAllPacketCommitments(ctx sdk.Context) []types.PacketState {
	_ = "STUB: not implemented"
	return nil
}

// IteratePacketCommitmentAtChannel provides an iterator over all PacketCommitment objects
// at a specified channel. For each packet commitment, cb will be called. If the cb returns
// true, the iterator will close and stop.
func (k *Keeper) IteratePacketCommitmentAtChannel(ctx sdk.Context, portID, channelID string, cb func(_, _ string, sequence uint64, hash []byte) bool) {
	_ = "STUB: not implemented"
	return
}

// GetAllPacketCommitmentsAtChannel returns all stored PacketCommitments objects for a specified
// port ID and channel ID.
func (k *Keeper) GetAllPacketCommitmentsAtChannel(ctx sdk.Context, portID, channelID string) []types.PacketState {
	_ = "STUB: not implemented"
	return nil
}

// IteratePacketReceipt provides an iterator over all PacketReceipt objects. For each
// receipt, cb will be called. If the cb returns true, the iterator will close
// and stop.
func (k *Keeper) IteratePacketReceipt(ctx sdk.Context, cb func(portID, channelID string, sequence uint64, receipt []byte) bool) {
	_ = "STUB: not implemented"
	return
}

// GetAllPacketReceipts returns all stored PacketReceipt objects.
func (k *Keeper) GetAllPacketReceipts(ctx sdk.Context) []types.PacketState {
	_ = "STUB: not implemented"
	return nil
}

// IteratePacketAcknowledgement provides an iterator over all PacketAcknowledgement objects. For each
// acknowledgement, cb will be called. If the cb returns true, the iterator will close
// and stop.
func (k *Keeper) IteratePacketAcknowledgement(ctx sdk.Context, cb func(portID, channelID string, sequence uint64, hash []byte) bool) {
	_ = "STUB: not implemented"
	return
}

// GetAllPacketAcks returns all stored PacketAcknowledgements objects.
func (k *Keeper) GetAllPacketAcks(ctx sdk.Context) []types.PacketState {
	_ = "STUB: not implemented"
	return nil
}

// IterateChannels provides an iterator over all Channel objects. For each
// Channel, cb will be called. If the cb returns true, the iterator will close
// and stop.
func (k *Keeper) IterateChannels(ctx sdk.Context, cb func(types.IdentifiedChannel) bool) {
	_ = "STUB: not implemented"
	return
}

// GetAllChannelsWithPortPrefix returns all channels with the specified port prefix. If an empty prefix is provided
// all channels will be returned.
func (k *Keeper) GetAllChannelsWithPortPrefix(ctx sdk.Context, portPrefix string) []types.IdentifiedChannel {
	_ = "STUB: not implemented"
	return nil
}

// GetAllChannels returns all stored Channel objects.
func (k *Keeper) GetAllChannels(ctx sdk.Context) []types.IdentifiedChannel {
	_ = "STUB: not implemented"
	return nil
}

// GetChannelClientState returns the associated client state with its ID, from a port and channel identifier.
func (k *Keeper) GetChannelClientState(ctx sdk.Context, portID, channelID string) (string, exported.ClientState, error) {
	_ = "STUB: not implemented"
	return "", *new(exported.ClientState), nil
}

// GetConnection wraps the connection keeper's GetConnection function.
func (k *Keeper) GetConnection(ctx sdk.Context, connectionID string) (connectiontypes.ConnectionEnd, error) {
	_ = "STUB: not implemented"
	return *new(connectiontypes.ConnectionEnd), nil
}

// GetChannelConnection returns the connection ID and state associated with the given port and channel identifier.
func (k *Keeper) GetChannelConnection(ctx sdk.Context, portID, channelID string) (string, connectiontypes.ConnectionEnd, error) {
	_ = "STUB: not implemented"
	return "", *new(connectiontypes.ConnectionEnd), nil
}

// common functionality for IteratePacketCommitment and IteratePacketAcknowledgement
func (k *Keeper) iterateHashes(ctx sdk.Context, iterator db.Iterator, cb func(portID, channelID string, sequence uint64, hash []byte) bool) {
	_ = "STUB: not implemented"
	return
}

// HasInflightPackets returns true if there are packet commitments stored at the specified
// port and channel, and false otherwise.
func (k *Keeper) HasInflightPackets(ctx sdk.Context, portID, channelID string) bool {
	_ = "STUB: not implemented"
	return false
}

// setRecvStartSequence sets the channel's recv start sequence to the store.
func (k *Keeper) setRecvStartSequence(ctx sdk.Context, portID, channelID string, sequence uint64) {
	_ = "STUB: not implemented"
	return
}

// GetRecvStartSequence gets a channel's recv start sequence from the store.
// The recv start sequence will be set to the counterparty's next sequence send
// upon a successful channel upgrade. It will be used for replay protection of
// historical packets and as the upper bound for pruning stale packet receives.
func (k *Keeper) GetRecvStartSequence(ctx sdk.Context, portID, channelID string) (uint64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (k *Keeper) GetV2Counterparty(ctx sdk.Context, portID string, channelID string) (clientv2types.CounterpartyInfo, bool) {
	_ = "STUB: not implemented"
	return *new(clientv2types.CounterpartyInfo), false
}

// Do not allow channel to be converted into a version 2 counterparty
// if the channel is not OPEN or if it is ORDERED
