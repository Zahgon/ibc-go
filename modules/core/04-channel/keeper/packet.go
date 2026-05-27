package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

// SendPacket is called by a module in order to send an IBC packet on a channel.
// The packet sequence generated for the packet to be sent is returned. An error
// is returned if one occurs.
func (k *Keeper) SendPacket(
	ctx sdk.Context,
	sourcePort string,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
	data []byte,
) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// construct packet from given fields and channel state

// prevent accidental sends with clients that cannot be updated

// check if packet is timed out on the receiving chain

// RecvPacket is called by a module in order to receive & process an IBC packet
// sent on the corresponding channel end on the counterparty chain.
func (k *Keeper) RecvPacket(
	ctx sdk.Context,
	packet types.Packet,
	proof []byte,
	proofHeight exported.Height,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// packet must come from the channel's counterparty

// Connection must be OPEN to receive a packet. It is possible for connection to not yet be open if packet was
// sent optimistically before connection and channel handshake completed. However, to receive a packet,
// connection and channel must both be open

// check if packet timed out by comparing it with the latest height of the chain

// verify that the counterparty did commit to sending this packet

// log that a packet has been received & executed

// emit an event that the relayer can query for

// applyReplayProtection ensures a packet has not already been received
// and performs the necessary state changes to ensure it cannot be received again.
func (k *Keeper) applyReplayProtection(ctx sdk.Context, packet types.Packet, channel types.Channel) error {
	_ = "STUB: not implemented"
	// REPLAY PROTECTION: The recvStartSequence will prevent historical proofs from allowing replay
	// attacks on packets processed in previous lifecycles of a channel. After a successful channel
	// upgrade all packets under the recvStartSequence will have been processed and thus should be
	// rejected.
	return nil
}

// REPLAY PROTECTION: Packet receipts will indicate that a packet has already been received
// on unordered channels. Packet receipts must not be pruned, unless it has been marked stale
// by the increase of the recvStartSequence.

// This error indicates that the packet has already been relayed. Core IBC will
// treat this error as a no-op in order to prevent an entire relay transaction
// from failing and consuming unnecessary fees.

// All verification complete, update state
// For unordered channels we must set the receipt so it can be verified on the other side.
// This receipt does not contain any data, since the packet has not yet been processed,
// it's just a single store key set to a single byte to indicate that the packet has been received

// check if the packet is being received in order

// This error indicates that the packet has already been relayed. Core IBC will
// treat this error as a no-op in order to prevent an entire relay transaction
// from failing and consuming unnecessary fees.

// REPLAY PROTECTION: Ordered channels require packets to be received in a strict order.
// Any out of order or previously received packets are rejected.

// All verification complete, update state
// In ordered case, we must increment nextSequenceRecv

// incrementing nextSequenceRecv and storing under this chain's channelEnd identifiers
// Since this is the receiving chain, our channelEnd is packet's destination port and channel

// WriteAcknowledgement writes the packet execution acknowledgement to the state,
// which will be verified by the counterparty chain using AcknowledgePacket.
//
// CONTRACT:
//
// 1) For synchronous execution, this function is be called in the IBC handler .
// For async handling, it needs to be called directly by the module which originally
// processed the packet.
//
// 2) Assumes that packet receipt has been written (unordered), or nextSeqRecv was incremented (ordered)
// previously by RecvPacket.
func (k *Keeper) WriteAcknowledgement(
	ctx sdk.Context,
	packet exported.PacketI,
	acknowledgement exported.Acknowledgement,
) error {
	_ = "STUB: not implemented"
	return nil
}

// REPLAY PROTECTION: The recvStartSequence will prevent historical proofs from allowing replay
// attacks on packets processed in previous lifecycles of a channel. After a successful channel
// upgrade all packets under the recvStartSequence will have been processed and thus should be
// rejected. Any asynchronous acknowledgement writes from packets processed in a previous lifecycle of a channel
// will also be rejected.

// NOTE: IBC app modules might have written the acknowledgement synchronously on
// the OnRecvPacket callback so we need to check if the acknowledgement is already
// set on the store and return an error if so.

// set the acknowledgement so that it can be verified on the other side

// log that a packet acknowledgement has been written

// AcknowledgePacket is called by a module to process the acknowledgement of a
// packet previously sent by the calling module on a channel to a counterparty
// module on the counterparty chain. Its intended usage is within the ante
// handler. AcknowledgePacket will clean up the packet commitment,
// which is no longer necessary since the packet has been received and acted upon.
// It will also increment NextSequenceAck in case of ORDERED channels.
func (k *Keeper) AcknowledgePacket(
	ctx sdk.Context,
	packet types.Packet,
	acknowledgement []byte,
	proof []byte,
	proofHeight exported.Height,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// packet must have been sent to the channel's counterparty

// This error indicates that the acknowledgement has already been relayed
// or there is a misconfigured relayer attempting to prove an acknowledgement
// for a packet never sent. Core IBC will treat this error as a no-op in order to
// prevent an entire relay transaction from failing and consuming unnecessary fees.

// verify we sent the packet and haven't cleared it out yet

// assert packets acknowledged in order

// All verification complete, in the case of ORDERED channels we must increment nextSequenceAck

// incrementing NextSequenceAck and storing under this chain's channelEnd identifiers
// Since this is the original sending chain, our channelEnd is packet's source port and channel

// Delete packet commitment, since the packet has been acknowledged, the commitment is no longer necessary

// log that a packet has been acknowledged

// emit an event marking that we have processed the acknowledgement
