package keeper

import (
	corestore "cosmossdk.io/core/store"
	"cosmossdk.io/log/v2"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/core/03-connection/types"
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

// Keeper defines the IBC connection keeper
type Keeper struct {
	// implements gRPC QueryServer interface
	types.QueryServer

	storeService corestore.KVStoreService
	cdc          codec.BinaryCodec
	clientKeeper types.ClientKeeper
}

// NewKeeper creates a new IBC connection Keeper instance
func NewKeeper(cdc codec.BinaryCodec, storeService corestore.KVStoreService, ck types.ClientKeeper) *Keeper {
	_ = "STUB: not implemented"
	return nil
}

// Logger returns a module-specific logger.
func (*Keeper) Logger(ctx sdk.Context) log.Logger {
	_ = "STUB: not implemented"
	return *new(log.Logger)
}

// GetCommitmentPrefix returns the IBC connection store prefix as a commitment
// Prefix
func (*Keeper) GetCommitmentPrefix() exported.Prefix {
	_ = "STUB: not implemented"
	return *new(exported.Prefix)
}

// GenerateConnectionIdentifier returns the next connection identifier.
func (k *Keeper) GenerateConnectionIdentifier(ctx sdk.Context) string {
	_ = "STUB: not implemented"
	return ""
}

// GetConnection returns a connection with a particular identifier
func (k *Keeper) GetConnection(ctx sdk.Context, connectionID string) (types.ConnectionEnd, bool) {
	_ = "STUB: not implemented"
	return *new(types.ConnectionEnd), false
}

// HasConnection returns a true if the connection with the given identifier
// exists in the store.
func (k *Keeper) HasConnection(ctx sdk.Context, connectionID string) bool {
	_ = "STUB: not implemented"
	return false
}

// SetConnection sets a connection to the store
func (k *Keeper) SetConnection(ctx sdk.Context, connectionID string, connection types.ConnectionEnd) {
	_ = "STUB: not implemented"
	return
}

// GetClientConnectionPaths returns all the connection paths stored under a
// particular client
func (k *Keeper) GetClientConnectionPaths(ctx sdk.Context, clientID string) ([]string, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// SetClientConnectionPaths sets the connections paths for client
func (k *Keeper) SetClientConnectionPaths(ctx sdk.Context, clientID string, paths []string) {
	_ = "STUB: not implemented"
	return
}

// GetNextConnectionSequence gets the next connection sequence from the store.
func (k *Keeper) GetNextConnectionSequence(ctx sdk.Context) uint64 {
	_ = "STUB: not implemented"
	return 0
}

// SetNextConnectionSequence sets the next connection sequence to the store.
func (k *Keeper) SetNextConnectionSequence(ctx sdk.Context, sequence uint64) {
	_ = "STUB: not implemented"
	return
}

// GetAllClientConnectionPaths returns all stored clients connection id paths. It
// will ignore the clients that haven't initialized a connection handshake since
// no paths are stored.
func (k *Keeper) GetAllClientConnectionPaths(ctx sdk.Context) []types.ConnectionPaths {
	_ = "STUB: not implemented"
	return nil
}

// continue when connection handshake is not initialized

// IterateConnections provides an iterator over all ConnectionEnd objects.
// For each ConnectionEnd, cb will be called. If the cb returns true, the
// iterator will close and stop.
func (k *Keeper) IterateConnections(ctx sdk.Context, cb func(types.IdentifiedConnection) bool) {
	_ = "STUB: not implemented"
	return
}

// GetAllConnections returns all stored ConnectionEnd objects.
func (k *Keeper) GetAllConnections(ctx sdk.Context) []types.IdentifiedConnection {
	_ = "STUB: not implemented"
	return nil
}

// CreateSentinelLocalhostConnection creates and sets the sentinel localhost connection end in the IBC store.
func (k *Keeper) CreateSentinelLocalhostConnection(ctx sdk.Context) {
	_ = "STUB: not implemented"
	return
}

// addConnectionToClient is used to add a connection identifier to the set of
// connections associated with a client.
func (k *Keeper) addConnectionToClient(ctx sdk.Context, clientID, connectionID string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetParams returns the total set of ibc-connection parameters.
func (k *Keeper) GetParams(ctx sdk.Context) types.Params {
	_ = "STUB: not implemented"
	return *new(types.Params)
}

// only panic on unset params and not on empty params

// SetParams sets the total set of ibc-connection parameters.
func (k *Keeper) SetParams(ctx sdk.Context, params types.Params) { _ = "STUB: not implemented"; return }
