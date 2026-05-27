package types

import (
	sdkmath "cosmossdk.io/math"
)

func (r *RateLimit) UpdateFlow(direction PacketDirection, amount sdkmath.Int) error {
	_ = "STUB: not implemented"
	return nil
}
