package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper *Keeper
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper *Keeper) Migrator { _ = "STUB: not implemented"; return *new(Migrator) }

// Migrate7To8 migrates the channel store from module version 7 to 8 by:
// - Removing channel upgrade sequences
// - Removing any channel upgrade info (i.e. upgrades, counterparty upgrades, upgrade errors)
// - Removing channel params
// - Removing pruning sequences
// NOTE: This migration will fail if any channels are in the FLUSHING or FLUSHCOMPLETE state.
func (m *Migrator) Migrate7To8(ctx sdk.Context) error { _ = "STUB: not implemented"; return nil }
