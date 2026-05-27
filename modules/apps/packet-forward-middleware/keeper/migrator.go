package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Migrator is a struct for handling in-place state migrations.
type Migrator struct {
	keeper *Keeper
}

func NewMigrator(k *Keeper) Migrator { _ = "STUB: not implemented"; return *new(Migrator) }

// Migrate2to3 migrates the module state from the consensus version 2 to
// version 3
func (m Migrator) Migrate2to3(ctx sdk.Context) error { _ = "STUB: not implemented"; return nil }

// Migrate3to4 migrates the module state from the consensus version 3 to
// version 4.
func (m Migrator) Migrate3to4(ctx sdk.Context) error { _ = "STUB: not implemented"; return nil }
