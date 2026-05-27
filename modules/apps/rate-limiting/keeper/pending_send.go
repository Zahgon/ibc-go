package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Sets the sequence number of a packet that was just sent
func (k *Keeper) SetPendingSendPacket(ctx sdk.Context, channelID string, sequence uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// Remove a pending packet sequence number from the store
// Used after the ack or timeout for a packet has been received
func (k *Keeper) RemovePendingSendPacket(ctx sdk.Context, channelID string, sequence uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// Checks whether the packet sequence number is in the store - indicating that it was
// sent during the current quota
func (k *Keeper) CheckPacketSentDuringCurrentQuota(ctx sdk.Context, channelID string, sequence uint64) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Get all pending packet sequence numbers
func (k *Keeper) GetAllPendingSendPackets(ctx sdk.Context) (pendingPackets []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// removes null bytes from suffix

// Remove all pending sequence numbers from the store
// This is executed when the quota resets
func (k *Keeper) RemoveAllChannelPendingSendPackets(ctx sdk.Context, channelID string) (err error) {
	_ = "STUB: not implemented"
	return nil
}
