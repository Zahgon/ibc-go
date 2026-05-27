package keeper

import (
	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/apps/rate-limiting/types"
)

// The total value on a given path (aka, the denominator in the percentage calculation)
// is the total supply of the given denom
func (k *Keeper) GetChannelValue(ctx sdk.Context, denom string) sdkmath.Int {
	_ = "STUB: not implemented"
	return *new(sdkmath.Int)
}

// CheckRateLimitAndUpdateFlow checks whether the given packet will exceed the rate limit.
// Called by OnRecvPacket and OnSendPacket
func (k *Keeper) CheckRateLimitAndUpdateFlow(ctx sdk.Context, direction types.PacketDirection, packetInfo RateLimitedPacketInfo) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// First check if the denom is blacklisted

// If there's no rate limit yet for this denom, no action is necessary

// Check if the sender/receiver pair is whitelisted
// If so, return a success without modifying the quota

// Update the flow object with the change in amount

// If the rate limit was exceeded, emit an event

// If there's no quota error, update the rate limit object in the store with the new flow

// If a SendPacket fails or times out, undo the outflow increment that happened during the send
func (k *Keeper) UndoSendPacket(ctx sdk.Context, channelOrClientID string, sequence uint64, denom string, amount sdkmath.Int) error {
	_ = "STUB: not implemented"
	return nil
}

// If the packet was sent during this quota, decrement the outflow
// Otherwise, it can be ignored
