package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	internaltypes "github.com/cosmos/ibc-go/v11/modules/apps/transfer/internal/types"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper) Migrator { _ = "STUB: not implemented"; return *new(Migrator) }

// MigrateDenomMetadata sets token metadata for all the IBC denom traces
func (m Migrator) MigrateDenomMetadata(ctx sdk.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// check if the metadata for the given denom trace does not already exist

// MigrateTotalEscrowForDenom migrates the total amount of source chain tokens in escrow.
func (m Migrator) MigrateTotalEscrowForDenom(ctx sdk.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// MigrateDenomTraceToDenom migrates storage from using DenomTrace to Denom.
func (m Migrator) MigrateDenomTraceToDenom(ctx sdk.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// convert denomTrace to denom

// defense in depth

// This migration must not change the SDK coin denom.
// A panic should occur to prevent the chain from using corrupted state.

// setDenomTrace sets a new {trace hash -> denom trace} pair to the store.
func (k *Keeper) setDenomTrace(ctx sdk.Context, denomTrace internaltypes.DenomTrace) {
	_ = "STUB: not implemented"
	return
}

// deleteDenomTrace deletes the denom trace
func (k *Keeper) deleteDenomTrace(ctx sdk.Context, denomTrace internaltypes.DenomTrace) {
	_ = "STUB: not implemented"
	return
}

// iterateDenomTraces iterates over the denomination traces in the store
// and performs a callback function.
func (k *Keeper) iterateDenomTraces(ctx sdk.Context, cb func(denomTrace internaltypes.DenomTrace) bool) {
	_ = "STUB: not implemented"
	return
}

// setDenomMetadataWithDenomTrace sets an IBC token's denomination metadata
func (k *Keeper) setDenomMetadataWithDenomTrace(ctx sdk.Context, denomTrace internaltypes.DenomTrace) {
	_ = "STUB: not implemented"
	return
}

// Setting base as IBC hash denom since bank keepers's SetDenomMetadata uses
// Base as key path and the IBC hash is what gives this token uniqueness
// on the executing chain
