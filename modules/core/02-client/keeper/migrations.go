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

// MigrateToStatelessLocalhost deletes the localhost client state. The localhost
// implementation is now stateless.
func (m Migrator) MigrateToStatelessLocalhost(ctx sdk.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// delete the client state
