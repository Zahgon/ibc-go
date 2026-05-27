package keeper

import (
	corestore "cosmossdk.io/core/store"
	"cosmossdk.io/log/v2"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/v2/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

// Keeper represents a type that grants read and write permissions to any client
// state information
type Keeper struct {
	storeService  corestore.KVStoreService
	cdc           codec.BinaryCodec
	router        *types.Router
	upgradeKeeper types.UpgradeKeeper
}

// NewKeeper creates a new NewKeeper instance
func NewKeeper(cdc codec.BinaryCodec, storeService corestore.KVStoreService, uk types.UpgradeKeeper) *Keeper {
	_ = "STUB: not implemented"
	return nil
}

// Codec returns the IBC Client module codec.
func (k *Keeper) Codec() codec.BinaryCodec {
	_ = "STUB: not implemented"

	// Logger returns a module-specific logger.
	return *new(codec.BinaryCodec)
}

func (*Keeper) Logger(ctx sdk.Context) log.Logger {
	_ = "STUB: not implemented"
	return *new(log.Logger)
}

// AddRoute adds a new route to the underlying router.
func (k *Keeper) AddRoute(clientType string, module exported.LightClientModule) {
	_ = "STUB: not implemented"
	return
}

// GetStoreProvider returns the light client store provider.
func (k *Keeper) GetStoreProvider() types.StoreProvider {
	_ = "STUB: not implemented"
	return *new(types.StoreProvider)
}

// Route returns the light client module for the given client identifier.
func (k *Keeper) Route(ctx sdk.Context, clientID string) (exported.LightClientModule, error) {
	_ = "STUB: not implemented"
	return *new(exported.LightClientModule), nil
}

// GenerateClientIdentifier returns the next client identifier.
func (k *Keeper) GenerateClientIdentifier(ctx sdk.Context, clientType string) string {
	_ = "STUB: not implemented"
	return ""
}

// GetClientState gets a particular client from the store
func (k *Keeper) GetClientState(ctx sdk.Context, clientID string) (exported.ClientState, bool) {
	_ = "STUB: not implemented"
	return *new(exported.ClientState), false
}

// SetClientState sets a particular Client to the store
func (k *Keeper) SetClientState(ctx sdk.Context, clientID string, clientState exported.ClientState) {
	_ = "STUB: not implemented"
	return
}

// GetClientCreator returns the creator of a client
func (k *Keeper) GetClientCreator(ctx sdk.Context, clientID string) sdk.AccAddress {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress)
}

// SetClientCreator sets the creator of a client
func (k *Keeper) SetClientCreator(ctx sdk.Context, clientID string, creator sdk.AccAddress) {
	_ = "STUB: not implemented"
	return
}

// DeleteClientCreator deletes the creator of a client
func (k *Keeper) DeleteClientCreator(ctx sdk.Context, clientID string) {
	_ = "STUB: not implemented"
	return
}

// GetClientConsensusState gets the stored consensus state from a client at a given height.
func (k *Keeper) GetClientConsensusState(ctx sdk.Context, clientID string, height exported.Height) (exported.ConsensusState, bool) {
	_ = "STUB: not implemented"
	return *new(exported.ConsensusState), false
}

// SetClientConsensusState sets a ConsensusState to a particular client at the given
// height
func (k *Keeper) SetClientConsensusState(ctx sdk.Context, clientID string, height exported.Height, consensusState exported.ConsensusState) {
	_ = "STUB: not implemented"
	return
}

// GetNextClientSequence gets the next client sequence from the store.
func (k *Keeper) GetNextClientSequence(ctx sdk.Context) uint64 { _ = "STUB: not implemented"; return 0 }

// SetNextClientSequence sets the next client sequence to the store.
func (k *Keeper) SetNextClientSequence(ctx sdk.Context, sequence uint64) {
	_ = "STUB: not implemented"
	return
}

// IterateConsensusStates provides an iterator over all stored consensus states.
// objects. For each State object, cb will be called. If the cb returns true,
// the iterator will close and stop.
func (k *Keeper) IterateConsensusStates(ctx sdk.Context, cb func(clientID string, cs types.ConsensusStateWithHeight) bool) {
	_ = "STUB: not implemented"
	return
}

// consensus key is in the format "clients/<clientID>/consensusStates/<height>"

// iterateMetadata provides an iterator over all stored metadata keys in the client store.
// For each metadata object, it will perform a callback.
func (k *Keeper) iterateMetadata(ctx sdk.Context, cb func(clientID string, key, value []byte) bool) {
	_ = "STUB: not implemented"
	return
}

// skip client state keys

// skip consensus state keys

// GetAllGenesisClients returns all the clients in state with their client ids returned as IdentifiedClientState
func (k *Keeper) GetAllGenesisClients(ctx sdk.Context) types.IdentifiedClientStates {
	_ = "STUB: not implemented"
	return *new(types.IdentifiedClientStates)
}

// GetAllClientMetadata will take a list of IdentifiedClientState and return a list
// of IdentifiedGenesisMetadata necessary for exporting and importing client metadata
// into the client store.
func (k *Keeper) GetAllClientMetadata(ctx sdk.Context, genClients []types.IdentifiedClientState) ([]types.IdentifiedGenesisMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetAllClientMetadata takes a list of IdentifiedGenesisMetadata and stores all of the metadata in the client store at the appropriate paths.
func (k *Keeper) SetAllClientMetadata(ctx sdk.Context, genMetadata []types.IdentifiedGenesisMetadata) {
	_ = "STUB: not implemented"
	return
}

// create client store

// set all metadata kv pairs in client store

// GetAllConsensusStates returns all stored client consensus states.
func (k *Keeper) GetAllConsensusStates(ctx sdk.Context) types.ClientsConsensusStates {
	_ = "STUB: not implemented"
	return *new(types.ClientsConsensusStates)
}

// HasClientConsensusState returns if keeper has a ConsensusState for a particular
// client at the given height
func (k *Keeper) HasClientConsensusState(ctx sdk.Context, clientID string, height exported.Height) bool {
	_ = "STUB: not implemented"
	return false
}

// GetLatestClientConsensusState gets the latest ConsensusState stored for a given client
func (k *Keeper) GetLatestClientConsensusState(ctx sdk.Context, clientID string) (exported.ConsensusState, bool) {
	_ = "STUB: not implemented"
	return *new(exported.ConsensusState), false
}

// VerifyMembership retrieves the light client module for the clientID and verifies the proof of the existence of a key-value pair at a specified height.
func (k *Keeper) VerifyMembership(ctx sdk.Context, clientID string, height exported.Height, delayTimePeriod uint64, delayBlockPeriod uint64, proof []byte, path exported.Path, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyNonMembership retrieves the light client module for the clientID and verifies the absence of a given key at a specified height.
func (k *Keeper) VerifyNonMembership(ctx sdk.Context, clientID string, height exported.Height, delayTimePeriod uint64, delayBlockPeriod uint64, proof []byte, path exported.Path) error {
	_ = "STUB: not implemented"
	return nil
}

// GetUpgradePlan executes the upgrade keeper GetUpgradePlan function.
func (k *Keeper) GetUpgradePlan(ctx sdk.Context) (upgradetypes.Plan, error) {
	_ = "STUB: not implemented"
	return *new(upgradetypes.Plan), nil
}

// GetUpgradedClient executes the upgrade keeper GetUpgradeClient function.
func (k *Keeper) GetUpgradedClient(ctx sdk.Context, planHeight int64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetUpgradedConsensusState returns the upgraded consensus state
func (k *Keeper) GetUpgradedConsensusState(ctx sdk.Context, planHeight int64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetUpgradedConsensusState executes the upgrade keeper SetUpgradedConsensusState function.
func (k *Keeper) SetUpgradedConsensusState(ctx sdk.Context, planHeight int64, bz []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// IterateClientStates provides an iterator over all stored ibc ClientState
// objects using the provided store prefix. For each ClientState object, cb will be called. If the cb returns true,
// the iterator will close and stop.
func (k *Keeper) IterateClientStates(ctx sdk.Context, storePrefix []byte, cb func(clientID string, cs exported.ClientState) bool) {
	_ = "STUB: not implemented"
	return
}

// skip non client state keys

// GetAllClients returns all stored light client State objects.
func (k *Keeper) GetAllClients(ctx sdk.Context) []exported.ClientState {
	_ = "STUB: not implemented"
	return nil
}

// ClientStore returns isolated prefix store for each client so they can read/write in separate
// namespace without being able to read/write other client's data
func (k *Keeper) ClientStore(ctx sdk.Context, clientID string) storetypes.KVStore {
	_ = "STUB: not implemented"
	return *new(storetypes.KVStore)
}

// GetClientStatus returns the status for a client state  given a client identifier. If the client type is not in the allowed
// clients param field, Unauthorized is returned, otherwise the client state status is returned.
func (k *Keeper) GetClientStatus(ctx sdk.Context, clientID string) exported.Status {
	_ = "STUB: not implemented"
	return *new(exported.Status)
}

// GetClientLatestHeight returns the latest height of a client state for a given client identifier. If the client type is not in the allowed
// clients param field, a zero value height is returned, otherwise the client state latest height is returned.
func (k *Keeper) GetClientLatestHeight(ctx sdk.Context, clientID string) types.Height {
	_ = "STUB: not implemented"
	return *new(types.Height)
}

// GetClientTimestampAtHeight returns the timestamp in nanoseconds of the consensus state at the given height.
func (k *Keeper) GetClientTimestampAtHeight(ctx sdk.Context, clientID string, height exported.Height) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GetParams returns the total set of ibc-client parameters.
func (k *Keeper) GetParams(ctx sdk.Context) types.Params {
	_ = "STUB: not implemented"
	return *new(types.Params)
}

// only panic on unset params and not on empty params

// SetParams sets the total set of ibc-client parameters.
func (k *Keeper) SetParams(ctx sdk.Context, params types.Params) { _ = "STUB: not implemented"; return }

// ScheduleIBCSoftwareUpgrade schedules an upgrade for the IBC client.
func (k *Keeper) ScheduleIBCSoftwareUpgrade(ctx sdk.Context, plan upgradetypes.Plan, upgradedClientState exported.ClientState) error {
	_ = "STUB: not implemented"
	// zero out any custom fields before setting
	return nil
}

// sets the new upgraded client last height committed on this chain at plan.Height,
// since the chain will panic at plan.Height and new chain will resume at plan.Height

// emitting an event for scheduling an upgrade plan
