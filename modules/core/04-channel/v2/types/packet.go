package types

// NewPacket constructs a new packet.
func NewPacket(sequence uint64, sourceClient, destinationClient string, timeoutTimestamp uint64, payloads ...Payload) Packet {
	_ = "STUB: not implemented"
	return *new(Packet)
}

// NewPayload constructs a new Payload
func NewPayload(sourcePort, destPort, version, encoding string, value []byte) Payload {
	_ = "STUB: not implemented"
	return *new(Payload)
}

// ValidateBasic validates that a Packet satisfies the basic requirements.
func (p Packet) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// ValidateBasic validates a Payload.
func (p Payload) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
