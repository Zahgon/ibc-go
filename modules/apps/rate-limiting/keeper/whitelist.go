package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/apps/rate-limiting/types"
)

// Adds an pair of sender and receiver addresses to the whitelist to allow all
// IBC transfers between those addresses to skip all flow calculations
func (k *Keeper) SetWhitelistedAddressPair(ctx sdk.Context, whitelist types.WhitelistedAddressPair) {
	_ = "STUB: not implemented"
	return
}

// Removes a whitelisted address pair so that it's transfers are counted in the quota
func (k *Keeper) RemoveWhitelistedAddressPair(ctx sdk.Context, sender, receiver string) {
	_ = "STUB: not implemented"
	return
}

// Check if a sender/receiver address pair is currently whitelisted
func (k *Keeper) IsAddressPairWhitelisted(ctx sdk.Context, sender, receiver string) bool {
	_ = "STUB: not implemented"
	return false
}

// Get all the whitelisted addresses
func (k *Keeper) GetAllWhitelistedAddressPairs(ctx sdk.Context) []types.WhitelistedAddressPair {
	_ = "STUB: not implemented"
	return nil
}
