package types

import (
	sdkmath "cosmossdk.io/math"
)

// Initializes a new flow from the channel value
func NewFlow(channelValue sdkmath.Int) Flow { _ = "STUB: not implemented"; return *new(Flow) }

// Adds an amount to the rate limit's flow after an incoming packet was received
// Returns an error if the new inflow will cause the rate limit to exceed its quota
func (f *Flow) AddInflow(amount sdkmath.Int, quota Quota) error {
	_ = "STUB: not implemented"
	return nil
}

// Adds an amount to the rate limit's flow after a packet was sent
// Returns an error if the new outflow will cause the rate limit to exceed its quota
func (f *Flow) AddOutflow(amount sdkmath.Int, quota Quota) error {
	_ = "STUB: not implemented"
	return nil
}
