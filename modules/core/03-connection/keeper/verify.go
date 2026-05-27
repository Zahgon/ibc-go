package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

// VerifyConnectionState verifies a proof of the connection state of the
// specified connection end stored on the target machine.
func (k *Keeper) VerifyConnectionState(
	ctx sdk.Context,
	connection types.ConnectionEnd,
	height exported.Height,
	proof []byte,
	connectionID string,
	counterpartyConnection types.ConnectionEnd, // opposite connection
) error {
	_ = "STUB: not implemented"
	return nil
}

// skip delay period checks for non-packet processing verification

// VerifyChannelState verifies a proof of the channel state of the specified
// channel end, under the specified port, stored on the target machine.
func (k *Keeper) VerifyChannelState(
	ctx sdk.Context,
	connection types.ConnectionEnd,
	height exported.Height,
	proof []byte,
	portID,
	channelID string,
	channel channeltypes.Channel,
) error {
	_ = "STUB: not implemented"
	return nil
}

// skip delay period checks for non-packet processing verification

// VerifyPacketCommitment verifies a proof of an outgoing packet commitment at
// the specified port, specified channel, and specified sequence.
func (k *Keeper) VerifyPacketCommitment(
	ctx sdk.Context,
	connection types.ConnectionEnd,
	height exported.Height,
	proof []byte,
	portID,
	channelID string,
	sequence uint64,
	commitmentBytes []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

// get time and block delays

// VerifyPacketAcknowledgement verifies a proof of an incoming packet
// acknowledgement at the specified port, specified channel, and specified sequence.
func (k *Keeper) VerifyPacketAcknowledgement(
	ctx sdk.Context,
	connection types.ConnectionEnd,
	height exported.Height,
	proof []byte,
	portID,
	channelID string,
	sequence uint64,
	acknowledgement []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

// get time and block delays

// VerifyPacketReceiptAbsence verifies a proof of the absence of an
// incoming packet receipt at the specified port, specified channel, and
// specified sequence.
func (k *Keeper) VerifyPacketReceiptAbsence(
	ctx sdk.Context,
	connection types.ConnectionEnd,
	height exported.Height,
	proof []byte,
	portID,
	channelID string,
	sequence uint64,
) error {
	_ = "STUB: not implemented"
	return nil
}

// get time and block delays

// VerifyNextSequenceRecv verifies a proof of the next sequence number to be
// received of the specified channel at the specified port.
func (k *Keeper) VerifyNextSequenceRecv(
	ctx sdk.Context,
	connection types.ConnectionEnd,
	height exported.Height,
	proof []byte,
	portID,
	channelID string,
	nextSequenceRecv uint64,
) error {
	_ = "STUB: not implemented"
	return nil
}

// get time and block delays

// getBlockDelay calculates the block delay period from the time delay of the connection
// and the maximum expected time per block.
func (k *Keeper) getBlockDelay(ctx sdk.Context, connection types.ConnectionEnd) uint64 {
	_ = "STUB: not implemented"
	// expectedTimePerBlock should never be zero, however if it is then return a 0 block delay for safety
	// as the expectedTimePerBlock parameter was not set.
	return 0
}

// calculate minimum block delay by dividing time delay period
// by the expected time per block. Round up the block delay.
