package types

import (
	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

const MaximumPayloadsSize = 262144 // 256 KiB. This is the maximum size of all payloads combined

// CommitPacket returns the packet commitment bytes. The commitment consists of:
// sha256_hash(timeout_timestamp + timeout_height.RevisionNumber + timeout_height.RevisionHeight + sha256_hash(data))
// from a given packet. This results in a fixed length preimage.
// NOTE: A fixed length preimage is ESSENTIAL to prevent relayers from being able
// to malleate the packet fields and create a commitment hash that matches the original packet.
// NOTE: sdk.Uint64ToBigEndian sets the uint64 to a slice of length 8.
func CommitPacket(packet Packet) []byte { _ = "STUB: not implemented"; return nil }

// CommitAcknowledgement returns the hash of commitment bytes
func CommitAcknowledgement(data []byte) []byte { _ = "STUB: not implemented"; return nil }

// NewPacket creates a new Packet instance. It panics if the provided
// packet data interface is not registered.
func NewPacket(
	data []byte,
	sequence uint64, sourcePort, sourceChannel,
	destinationPort, destinationChannel string,
	timeoutHeight clienttypes.Height, timeoutTimestamp uint64,
) Packet {
	_ = "STUB: not implemented"
	return *new(Packet)
}

// GetSequence implements PacketI interface
func (p Packet) GetSequence() uint64 {
	_ = "STUB: not implemented"

	// GetSourcePort implements PacketI interface
	return 0
}

func (p Packet) GetSourcePort() string {
	_ = "STUB: not implemented"

	// GetSourceChannel implements PacketI interface
	return ""
}

func (p Packet) GetSourceChannel() string { _ = "STUB: not implemented"; return "" }

// GetDestPort implements PacketI interface
func (p Packet) GetDestPort() string { _ = "STUB: not implemented"; return "" }

// GetDestChannel implements PacketI interface
func (p Packet) GetDestChannel() string { _ = "STUB: not implemented"; return "" }

// GetData implements PacketI interface
func (p Packet) GetData() []byte {
	_ = "STUB: not implemented"

	// GetTimeoutHeight implements PacketI interface
	return nil
}

func (p Packet) GetTimeoutHeight() exported.Height {
	_ = "STUB: not implemented"
	return *

	// GetTimeoutTimestamp implements PacketI interface
	new(exported.Height)
}

func (p Packet) GetTimeoutTimestamp() uint64 { _ = "STUB: not implemented"; return 0 }

// ValidateBasic implements PacketI interface
func (p Packet) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// Validates a PacketId
func (p PacketId) Validate() error { _ = "STUB: not implemented"; return nil }

// NewPacketID returns a new instance of PacketId
func NewPacketID(portID, channelID string, seq uint64) PacketId {
	_ = "STUB: not implemented"
	return *new(PacketId)
}
