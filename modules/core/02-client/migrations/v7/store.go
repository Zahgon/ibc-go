package v7

import (
	corestore "cosmossdk.io/core/store"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/v2/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	solomachine "github.com/cosmos/ibc-go/v11/modules/light-clients/06-solomachine"
)

// Localhost is the client type for a localhost client. It is also used as the clientID
// for the localhost client.
const Localhost string = "09-localhost"

// MigrateStore performs in-place store migrations from ibc-go v6 to ibc-go v7.
// The migration includes:
//
// - Migrating solo machine client states from v2 to v3 protobuf definition
// - Pruning all solo machine consensus states
// - Removing the localhost client
// - Asserting existing tendermint clients are properly registered on the chain codec
func MigrateStore(ctx sdk.Context, storeService corestore.KVStoreService, cdc codec.BinaryCodec, clientKeeper ClientKeeper) error {
	_ = "STUB: not implemented"
	return nil
}

// handleSolomachineMigration iterates over the solo machine clients and migrates client state from
// protobuf definition v2 to v3. All consensus states stored outside of the client state are pruned.
func handleSolomachineMigration(ctx sdk.Context, store storetypes.KVStore, cdc codec.BinaryCodec, clientKeeper ClientKeeper) error {
	_ = "STUB: not implemented"
	return nil
}

// update solomachine in store

// handleTendermintMigration asserts that the tendermint client in state can be decoded properly.
// This ensures the upgrading chain properly registered the tendermint client types on the chain codec.
func handleTendermintMigration(ctx sdk.Context, store storetypes.KVStore, clientKeeper ClientKeeper) error {
	_ = "STUB: not implemented"
	return nil
}

// no-op if no tm clients exist

// unregistered tendermint client types will panic when unmarshaling the client state
// in GetClientState

// handleLocalhostMigration removes all client and consensus states associated with the localhost client type.
func handleLocalhostMigration(ctx sdk.Context, store storetypes.KVStore, clientKeeper ClientKeeper) error {
	_ = "STUB: not implemented"
	return nil
}

// delete the client state

// collectClients will iterate over the provided client type prefix in the client store
// and return a list of clientIDs associated with the client type. This is necessary to
// avoid state corruption as modifying state during iteration is unsafe. A special case
// for tendermint clients is included as only one tendermint clientID is required for
// v7 migrations.
func collectClients(ctx sdk.Context, store storetypes.KVStore, clientType string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// skip non client state keys

// optimization: exit after a single tendermint client iteration

// removeAllClientConsensusStates removes all client consensus states from the associated
// client store.
func removeAllClientConsensusStates(clientStore storetypes.KVStore) {
	_ = "STUB: not implemented"
	return
}

// key is in the format "consensusStates/<height>"

// collect consensus states to be pruned

// delete all consensus states

// migrateSolomachine migrates the solomachine from v2 to v3 solo machine protobuf definition.
// Notably it drops the AllowUpdateAfterProposal field.
func migrateSolomachine(clientState ClientState) solomachine.ClientState {
	_ = "STUB: not implemented"
	return *new(solomachine.ClientState)
}
