package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper *Keeper
}

// NewMigrator creates a new Migrator instance.
func NewMigrator(k *Keeper) Migrator { _ = "STUB: not implemented"; return *new(Migrator) }

// Migrate1to2 widens the PendingSendPacket key's channel-ID segment from
// 16 to 64 bytes so IBC v2 channel IDs fit.
func (m Migrator) Migrate1to2(ctx sdk.Context) error { _ = "STUB: not implemented"; return nil }
