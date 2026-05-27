package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Adds a denom to a blacklist to prevent all IBC transfers with that denom
func (k *Keeper) AddDenomToBlacklist(ctx sdk.Context, denom string) {
	_ = "STUB: not implemented"
	return
}

// Removes a denom from a blacklist to re-enable IBC transfers for that denom
func (k *Keeper) RemoveDenomFromBlacklist(ctx sdk.Context, denom string) {
	_ = "STUB: not implemented"
	return
}

// Check if a denom is currently blacklisted
func (k *Keeper) IsDenomBlacklisted(ctx sdk.Context, denom string) bool {
	_ = "STUB: not implemented"
	return false
}

// Get all the blacklisted denoms
func (k *Keeper) GetAllBlacklistedDenoms(ctx sdk.Context) []string {
	_ = "STUB: not implemented"
	return nil
}
