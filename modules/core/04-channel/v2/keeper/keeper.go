package keeper

import (
	corestore "cosmossdk.io/core/store"
	"cosmossdk.io/log/v2"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	clientv2keeper "github.com/cosmos/ibc-go/v11/modules/core/02-client/v2/keeper"
	connectionkeeper "github.com/cosmos/ibc-go/v11/modules/core/03-connection/keeper"
	"github.com/cosmos/ibc-go/v11/modules/core/04-channel/v2/types"
	"github.com/cosmos/ibc-go/v11/modules/core/api"
)

// Keeper defines the channel keeper v2.
type Keeper struct {
	storeService corestore.KVStoreService
	cdc          codec.BinaryCodec
	ClientKeeper types.ClientKeeper
	// clientV2Keeper is used for counterparty access.
	clientV2Keeper *clientv2keeper.Keeper

	connectionKeeper *connectionkeeper.Keeper

	// Router is used to route messages to the appropriate module callbacks
	// NOTE: it must be explicitly set before usage.
	Router *api.Router
}

// NewKeeper creates a new channel v2 keeper
func NewKeeper(
	cdc codec.BinaryCodec,
	storeService corestore.KVStoreService,
	clientKeeper types.ClientKeeper,
	clientV2Keeper *clientv2keeper.Keeper,
	connectionKeeper *connectionkeeper.Keeper,
) *Keeper {
	_ = "STUB: not implemented"
	return nil
}

// Logger returns a module-specific logger.
func (*Keeper) Logger(ctx sdk.Context) log.Logger {
	_ = "STUB: not implemented"
	return *new(log.Logger)
}

// GetPacketReceipt returns the packet receipt from the packet receipt path based on the clientID and sequence.
func (k *Keeper) GetPacketReceipt(ctx sdk.Context, clientID string, sequence uint64) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// HasPacketReceipt returns true if the packet receipt exists, otherwise false.
func (k *Keeper) HasPacketReceipt(ctx sdk.Context, clientID string, sequence uint64) bool {
	_ = "STUB: not implemented"
	return false
}

// SetPacketReceipt writes the packet receipt under the receipt path
// This is a public path that is standardized by the IBC V2 specification.
func (k *Keeper) SetPacketReceipt(ctx sdk.Context, clientID string, sequence uint64) {
	_ = "STUB: not implemented"
	return
}

// GetPacketAcknowledgement fetches the packet acknowledgement from the store.
func (k *Keeper) GetPacketAcknowledgement(ctx sdk.Context, clientID string, sequence uint64) []byte {
	_ = "STUB: not implemented"
	return nil
}

// SetPacketAcknowledgement writes the acknowledgement hash under the acknowledgement path
// This is a public path that is standardized by the IBC V2 specification.
func (k *Keeper) SetPacketAcknowledgement(ctx sdk.Context, clientID string, sequence uint64, ackHash []byte) {
	_ = "STUB: not implemented"
	return
}

// HasPacketAcknowledgement checks if the packet ack hash is already on the store.
func (k *Keeper) HasPacketAcknowledgement(ctx sdk.Context, clientID string, sequence uint64) bool {
	_ = "STUB: not implemented"
	return false
}

// GetPacketCommitment returns the packet commitment hash under the commitment path.
func (k *Keeper) GetPacketCommitment(ctx sdk.Context, clientID string, sequence uint64) []byte {
	_ = "STUB: not implemented"
	return nil
}

// SetPacketCommitment writes the commitment hash under the commitment path.
func (k *Keeper) SetPacketCommitment(ctx sdk.Context, clientID string, sequence uint64, commitment []byte) {
	_ = "STUB: not implemented"
	return
}

// DeletePacketCommitment deletes the packet commitment hash under the commitment path.
func (k *Keeper) DeletePacketCommitment(ctx sdk.Context, clientID string, sequence uint64) {
	_ = "STUB: not implemented"
	return
}

// GetNextSequenceSend returns the next send sequence from the sequence path
func (k *Keeper) GetNextSequenceSend(ctx sdk.Context, clientID string) (uint64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// SetNextSequenceSend writes the next send sequence under the sequence path
func (k *Keeper) SetNextSequenceSend(ctx sdk.Context, clientID string, sequence uint64) {
	_ = "STUB: not implemented"
	return
}

// SetAsyncPacket writes the packet under the async path
func (k *Keeper) SetAsyncPacket(ctx sdk.Context, clientID string, sequence uint64, packet types.Packet) {
	_ = "STUB: not implemented"
	return
}

// GetAsyncPacket fetches the packet from the async path
func (k *Keeper) GetAsyncPacket(ctx sdk.Context, clientID string, sequence uint64) (types.Packet, bool) {
	_ = "STUB: not implemented"
	return *new(types.Packet), false
}

// DeleteAsyncPacket deletes the packet from the async path
func (k *Keeper) DeleteAsyncPacket(ctx sdk.Context, clientID string, sequence uint64) {
	_ = "STUB: not implemented"
	return
}

// extractSequenceFromKey takes the full store key as well as a packet store prefix and extracts
// the encoded sequence number from the key.
//
// This function panics of the provided key once trimmed is larger than 8 bytes as the expected
// sequence byte length is always 8.
func extractSequenceFromKey(key, storePrefix []byte) uint64 { _ = "STUB: not implemented"; return 0 }

// GetAllPacketCommitmentsForClient returns all stored PacketCommitments objects for a specified
// client ID.
func (k *Keeper) GetAllPacketCommitmentsForClient(ctx sdk.Context, clientID string) []types.PacketState {
	_ = "STUB: not implemented"
	return nil
}

// GetAllPacketAcknowledgementsForClient returns all stored PacketAcknowledgements objects for a specified
// client ID.
func (k *Keeper) GetAllPacketAcknowledgementsForClient(ctx sdk.Context, clientID string) []types.PacketState {
	_ = "STUB: not implemented"
	return nil
}

// GetAllPacketReceiptsForClient returns all stored PacketReceipts objects for a specified
// client ID.
func (k *Keeper) GetAllPacketReceiptsForClient(ctx sdk.Context, clientID string) []types.PacketState {
	_ = "STUB: not implemented"
	return nil
}

// GetAllAsyncPacketsForClient returns all stored AsyncPackets objects for a specified
// client ID.
func (k *Keeper) GetAllAsyncPacketsForClient(ctx sdk.Context, clientID string) []types.PacketState {
	_ = "STUB: not implemented"
	return nil
}

// prefixKeyConstructor is a function that constructs a store key for a specific packet store using the provided
// clientID.
type prefixKeyConstructor func(clientID string) []byte

// getAllPacketStateForClient gets all PacketState objects for the specified clientID using a provided
// function for constructing the key prefix for the store.
//
// For example, to get all PacketReceipts for a clientID the hostv2.PacketReceiptPrefixKey function can be
// passed to get the PacketReceipt store key prefix.
func (k *Keeper) getAllPacketStateForClient(ctx sdk.Context, clientID string, prefixFn prefixKeyConstructor) []types.PacketState {
	_ = "STUB: not implemented"
	return nil
}

// SetClientForAlias sets the base client ID under the alias key for an aliased channelID.
func (k *Keeper) SetClientForAlias(ctx sdk.Context, alias string, baseClientID string) {
	_ = "STUB: not implemented"
	return
}

// GetClientForAlias get the base client ID under the alias key for an aliased channelID.
func (k *Keeper) GetClientForAlias(ctx sdk.Context, alias string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}
