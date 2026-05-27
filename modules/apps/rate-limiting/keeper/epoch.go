package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/apps/rate-limiting/types"
)

// Stores the hour epoch
func (k *Keeper) SetHourEpoch(ctx sdk.Context, epoch types.HourEpoch) error {
	_ = "STUB: not implemented"
	return nil
}

// Reads the hour epoch from the store
// Returns a zero-value epoch and logs an error if the epoch is not found or fails to unmarshal.
func (k *Keeper) GetHourEpoch(ctx sdk.Context) (types.HourEpoch, error) {
	_ = "STUB: not implemented"
	return *new(types.HourEpoch), nil
}

// Checks if it's time to start the new hour epoch.
// This function returns epochStarting, epochNumber and a possible error.
func (k *Keeper) CheckHourEpochStarting(ctx sdk.Context) (bool, uint64, error) {
	_ = "STUB: not implemented"
	return false, 0, nil
}

// If GetHourEpoch returned a zero-value epoch (due to error or missing key),
// we cannot proceed with the check.

// If the block time is later than the current epoch start time + epoch duration,
// move onto the next epoch by incrementing the epoch number, height, and start time

// Otherwise, indicate that a new epoch is not starting
