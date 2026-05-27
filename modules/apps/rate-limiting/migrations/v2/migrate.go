package v2

import (
	corestore "cosmossdk.io/core/store"

	"github.com/cosmos/cosmos-sdk/store/v2/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// oldPendingSendPacketChannelLength is hard-coded so the migration stays
	// correct even if types.PendingSendPacketChannelLength is bumped again
	// later.
	oldPendingSendPacketChannelLength = 16
	oldKeyLen                         = oldPendingSendPacketChannelLength + 8

	// newPendingSendPacketChannelLength is hard-coded so the migration stays
	// correct even if types.PendingSendPacketChannelLength is bumped again
	// later.
	newPendingSendPacketChannelLength = 64
	newKeyLen                         = newPendingSendPacketChannelLength + 8
)

// Migrate rewrites entries under types.PendingSendPacketPrefix from the old
// [16-byte channelID][8-byte sequence] layout to the new [64-byte channelID]
// [8-byte sequence] layout so IBC v2 channel IDs (up to 64 bytes) fit. Entries
// already in the new layout are skipped, making the migration idempotent.
func Migrate(ctx sdk.Context, storeService corestore.KVStoreService) error {
	_ = "STUB: not implemented"
	return nil
}

// get store entries that need to be migrated

// migrate store entries
// old key layout: [16-byte channelID][8-byte sequence]
// new key layout: [64-byte channelID][8-byte sequence]

// place 16 byte channel id from old key into first 64 bytes of new key

// put remaining 8 bytes sequence from old key into the final 8 bytes
// sequence of the new key

// remove old kv and set new kv

type entry struct {
	key   [oldKeyLen]byte
	value []byte
}

// collectLegacyEntries returns a list of entries in the prefix store that must
// be migrated from the oldKeyLen to the newKeyLen.
func collectLegacyEntries(store prefix.Store) ([]entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// validateMigratedKey ensures a legacy key transformation is valid, returns an
// error if not.
func validateMigratedKey(newKey []byte, oldKey [oldKeyLen]byte) error {
	_ = "STUB: not implemented"
	// channelID is right-padded with null bytes in the key (see
	// types.PendingSendPacketKey), so trim them before validating.
	return nil
}

// validate channelID in the newKey.

// we are using the client
// validator here since in v1 these will be channelID's, and in v2 they
// are clientID's, the client validator is slightly less strict and
// will accept both.

// ensure we have not modified the existing value

// ensure after the oldPendingSendPacketChannelLength, we have 48 bytes of 0's

// validate sequence number in the newKey.

// ensure we have not modified the existing value
