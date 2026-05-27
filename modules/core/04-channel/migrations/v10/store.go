package v10

import (
	corestore "cosmossdk.io/core/store"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ParamsKey defines the key to store the params in the keeper.
	ParamsKey               = "channelParams"
	KeyPruningSequenceStart = "pruningSequenceStart"

	KeyChannelUpgradePrefix = "channelUpgrades"
	KeyUpgradePrefix        = "upgrades"
	KeyUpgradeErrorPrefix   = "upgradeError"
	KeyCounterpartyUpgrade  = "counterpartyUpgrade"
)

// PruningSequenceStartKey returns the store key for the pruning sequence start of a particular channel
func PruningSequenceStartKey(portID, channelID string) []byte {
	_ = "STUB: not implemented"
	return nil
}

func ChannelUpgradeKey(portID, channelID string) []byte { _ = "STUB: not implemented"; return nil }

func ChannelUpgradeErrorKey(portID, channelID string) []byte { _ = "STUB: not implemented"; return nil }

func ChannelCounterpartyUpgradeKey(portID, channelID string) []byte {
	_ = "STUB: not implemented"
	return nil
}

// MigrateStore migrates the channel store to the ibc-go v10 store by:
// - Removing channel upgrade sequences
// - Removing any channel upgrade info (i.e. upgrades, counterparty upgrades, upgrade errors)
// - Removing channel params
// - Removing pruning sequences
// NOTE: This migration will fail if any channels are in the FLUSHING or FLUSHCOMPLETE state.
func MigrateStore(ctx sdk.Context, storeService corestore.KVStoreService, cdc codec.BinaryCodec, channelKeeper ChannelKeeper) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: See if there is more to migrate/delete from store

func handleChannelMigration(ctx sdk.Context, store corestore.KVStore, cdc codec.BinaryCodec, channelKeeper ChannelKeeper) error {
	_ = "STUB: not implemented"
	// Remove channel upgrade sequences and set in-upgrade channels back to open
	return nil
}

// Any pitfalls of doing this?

func deleteChannelUpgrades(store corestore.KVStore) error {
	_ = "STUB: not implemented"
	// Delete channel upgrades (i.e. upgrades, counterparty upgrades, upgrade errors, which are stored in the channelUpgrades prefix)
	return nil
}

func deleteParams(store corestore.KVStore) error {
	_ = "STUB: not implemented"
	// Delete channel params
	return nil
}

func deletePruneSequences(store corestore.KVStore) error {
	_ = "STUB: not implemented"
	// Delete all pruning sequences
	return nil
}
