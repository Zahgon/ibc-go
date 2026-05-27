package types

import (
	sdkmath "cosmossdk.io/math"
)

// CheckExceedsQuota checks if new in/out flow is going to reach the max in/out or not
func (q *Quota) CheckExceedsQuota(direction PacketDirection, amount sdkmath.Int, totalValue sdkmath.Int) bool {
	_ = "STUB: not implemented"
	// If there's no channel value (this should be almost impossible), it means there is no
	// supply of the asset, so we shouldn't prevent inflows/outflows
	return false
}

// Revert to GT check as in the original reference module
