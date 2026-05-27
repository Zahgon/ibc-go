package keeper

import (
	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/apps/rate-limiting/types"
)

// If the rate limit is exceeded or the denom is blacklisted, we emit an event
func EmitTransferDeniedEvent(ctx sdk.Context, reason, denom, channelOrClientID string, direction types.PacketDirection, amount sdkmath.Int, err error) {
	_ = "STUB: not implemented"
	return
}

// packet_send or packet_recv
