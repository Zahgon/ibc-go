package types

// CommitPacket returns the V2 packet commitment bytes. The commitment consists of:
// ha256_hash(0x02 + sha256_hash(destinationClient) + sha256_hash(timeout) + sha256_hash(payload)) from a given packet.
// This results in a fixed length preimage of 32 bytes.
// NOTE: A fixed length preimage is ESSENTIAL to prevent relayers from being able
// to malleate the packet fields and create a commitment hash that matches the original packet.
func CommitPacket(packet Packet) []byte { _ = "STUB: not implemented"; return nil }

// hashPayload returns the hash of the payload.
func hashPayload(data Payload) []byte { _ = "STUB: not implemented"; return nil }

// CommitAcknowledgement returns the V2 acknowledgement commitment bytes. The commitment consists of:
// sha256_hash(0x02 + sha256_hash(ack1) + sha256_hash(ack2) + ...) from a given acknowledgement.
func CommitAcknowledgement(acknowledgement Acknowledgement) []byte {
	_ = "STUB: not implemented"
	return nil
}
