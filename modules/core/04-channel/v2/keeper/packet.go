package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/core/04-channel/v2/types"
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

// sendPacket constructs a packet from the input arguments, writes a packet commitment to state
// in order for the packet to be sent to the counterparty.
func (k *Keeper) sendPacket(
	ctx sdk.Context,
	sourceClient string,
	timeoutTimestamp uint64,
	payloads []types.Payload,
) (uint64, string, error) {
	_ = "STUB: not implemented"
	// lookup counterparty from packet identifiers
	// note this will be either the client identifier for IBC V2 paths
	// or an aliased channel identifier for IBC V1 paths
	return 0, "", nil
}

// Note, the validate basic function in sendPacket does the timeoutTimestamp != 0 check and other stateless checks on the packet.
// timeoutTimestamp must be greater than current block time

// timeoutTimestamp must be less than current block time + MaxTimeoutDelta

// construct packet from given fields and channel state

// Before we do client keeper level checks, we first get underlying base clientID

// check that the client of counterparty chain is still active

// retrieve latest height and timestamp of the client of counterparty chain

// client timestamps are in nanoseconds while packet timeouts are in seconds
// thus to compare them, we convert the client timestamp to seconds in uint64
// to be consistent with IBC V2 specified timeout behaviour

// bump the sequence and set the packet commitment, so it is provable by the counterparty

// recvPacket implements the packet receiving logic required by a packet handler.￼
// The packet is checked for correctness including asserting that the packet was
// sent and received on clients which are counterparties for one another.
// If the packet has already been received a no-op error is returned.
// The packet handler will verify that the packet has not timed out and that the
// counterparty stored a packet commitment. If successful, a packet receipt is stored
// to indicate to the counterparty successful delivery.
func (k *Keeper) recvPacket(
	ctx sdk.Context,
	packet types.Packet,
	proof []byte,
	proofHeight exported.Height,
) error {
	_ = "STUB: not implemented"
	// lookup counterparty from packet identifiers
	// note this will be either the client identifier for IBC V2 paths
	// or an aliased channel identifier for IBC V1 paths
	return nil
}

// REPLAY PROTECTION: Packet receipts will indicate that a packet has already been received
// Packet receipts must not be pruned, unless it has been marked stale
// by the increase of the recvStartSequence.

// This error indicates that the packet has already been relayed. Core IBC will
// treat this error as a no-op in order to prevent an entire relay transaction
// from failing and consuming unnecessary fees.

// Before we do client keeper level checks, we first get underlying base clientID

// Set Packet Receipt to prevent timeout from occurring on counterparty

// writeAcknowledgement writes the acknowledgement to the store and emits the packet and acknowledgement
// for relayers to relay the acknowledgement to the counterparty chain.
func (k *Keeper) writeAcknowledgement(
	ctx sdk.Context,
	packet types.Packet,
	ack types.Acknowledgement,
) error {
	_ = "STUB: not implemented"
	// Validate the acknowledgement
	return nil
}

// Validate the acknowledgement against the payload length

// lookup counterparty from packet identifiers
// note this will be either the client identifier for IBC V2 paths
// or an aliased channel identifier for IBC V1 paths

// NOTE: IBC app modules might have written the acknowledgement synchronously on
// the OnRecvPacket callback so we need to check if the acknowledgement is already
// set on the store and return an error if so.

// set the acknowledgement so that it can be verified on the other side

// WriteAcknowledgement writes the acknowledgement and emits events for asynchronous acknowledgements
// this is the method to be called by external apps when they want to write an acknowledgement asyncrhonously
func (k *Keeper) WriteAcknowledgement(ctx sdk.Context, clientID string, sequence uint64, ack types.Acknowledgement) error {
	_ = "STUB: not implemented"
	// get saved async packet from store
	return nil
}

// Write the acknowledgement to the store

// Delete the packet from the async store

func (k *Keeper) acknowledgePacket(ctx sdk.Context, packet types.Packet, acknowledgement types.Acknowledgement, proof []byte, proofHeight exported.Height) error {
	_ = "STUB: not implemented"
	// lookup counterparty from packet identifiers
	// note this will be either the client identifier for IBC V2 paths
	// or an aliased channel identifier for IBC V1 paths
	return nil
}

// This error indicates that the acknowledgement has already been relayed
// or there is a misconfigured relayer attempting to prove an acknowledgement
// for a packet never sent. Core IBC will treat this error as a no-op in order to
// prevent an entire relay transaction from failing and consuming unnecessary fees.

// verify we sent the packet and haven't cleared it out yet

// Before we do client keeper level checks, we first get underlying base clientID

// timeoutPacket implements the timeout logic required by a packet handler.
// The packet is checked for correctness including asserting that the packet was
// sent and received on clients which are counterparties for one another.
// If no packet commitment exists, a no-op error is returned, otherwise
// an absence proof of the packet receipt is performed to ensure that the packet
// was never delivered to the counterparty. If successful, the packet commitment
// is deleted and the packet has completed its lifecycle.
func (k *Keeper) timeoutPacket(
	ctx sdk.Context,
	packet types.Packet,
	proof []byte,
	proofHeight exported.Height,
) error {
	_ = "STUB: not implemented"
	// lookup counterparty from packet identifiers
	// note this will be either the client identifier for IBC V2 paths
	// or an aliased channel identifier for IBC V1 paths
	return nil
}

// Before we do client keeper level checks, we first get underlying base clientID

// check that timeout timestamp has passed on the other end
// client timestamps are in nanoseconds while packet timeouts are in seconds
// so we convert client timestamp to seconds in uint64 to be consistent
// with IBC V2 timeout behaviour

// check that the commitment has not been cleared and that it matches the packet sent by relayer

// This error indicates that the timeout has already been relayed
// or there is a misconfigured relayer attempting to prove a timeout
// for a packet never sent. Core IBC will treat this error as a no-op in order to
// prevent an entire relay transaction from failing and consuming unnecessary fees.

// verify we sent the packet and haven't cleared it out yet

// verify packet receipt absence

// delete packet commitment to prevent replay
