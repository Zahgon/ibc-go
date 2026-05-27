package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/apps/rate-limiting/types"
)

// Stores/Updates a rate limit object in the store
func (k *Keeper) SetRateLimit(ctx sdk.Context, rateLimit types.RateLimit) {
	_ = "STUB: not implemented"
	return
}

// Removes a rate limit object from the store using denom and channel-id
func (k *Keeper) RemoveRateLimit(ctx sdk.Context, denom string, channelID string) {
	_ = "STUB: not implemented"
	return
}

// Grabs and returns a rate limit object from the store using denom and channel-id
func (k *Keeper) GetRateLimit(ctx sdk.Context, denom string, channelID string) (types.RateLimit, bool) {
	_ = "STUB: not implemented"
	return *new(types.RateLimit), false
}

// Returns all rate limits stored
func (k *Keeper) GetAllRateLimits(ctx sdk.Context) []types.RateLimit {
	_ = "STUB: not implemented"
	return nil
}

// Log the error and skip this entry if unmarshalling fails

// Adds a new rate limit. Fails if the rate limit already exists or the channel value is 0
func (k *Keeper) AddRateLimit(ctx sdk.Context, msg *types.MsgAddRateLimit) error {
	_ = "STUB: not implemented"
	return nil
}

// Confirm the channel or client exists

// Check if the channelId is actually a clientId

// If the status is Unauthorized or Unknown, it means the client doesn't exist or is invalid

// Return specific error indicating neither channel nor client was found

// If status is Active, Expired, or Frozen, the client exists, proceed.

// Create and store the rate limit object

// Updates an existing rate limit. Fails if the rate limit doesn't exist
func (k *Keeper) UpdateRateLimit(ctx sdk.Context, msg *types.MsgUpdateRateLimit) error {
	_ = "STUB: not implemented"
	return nil
}

// Update the rate limit object with the new quota information
// The flow should also get reset to 0

// Reset the rate limit after expiration
// The inflow and outflow should get reset to 0, the channelValue should be updated,
// and all pending send packet sequence numbers should be removed
func (k *Keeper) ResetRateLimit(ctx sdk.Context, denom string, channelID string) error {
	_ = "STUB: not implemented"
	return nil
}
