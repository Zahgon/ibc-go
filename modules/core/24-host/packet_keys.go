package host

const (
	KeySequencePrefix         = "sequences"
	KeyNextSeqSendPrefix      = "nextSequenceSend"
	KeyNextSeqRecvPrefix      = "nextSequenceRecv"
	KeyNextSeqAckPrefix       = "nextSequenceAck"
	KeyPacketCommitmentPrefix = "commitments"
	KeyPacketAckPrefix        = "acks"
	KeyPacketReceiptPrefix    = "receipts"
	KeyRecvStartSequence      = "recvStartSequence"
)

// ICS04
// The following paths are the keys to the store as defined in https://github.com/cosmos/ibc/tree/master/spec/core/ics-004-channel-and-packet-semantics#store-paths
// NOTE: NextSequenceSendKey has been removed and we only use the IBC v2 key in this repo.
// We can safely do this since the NextSequenceSendKey is not proven to counterparties, thus we can use any key format we want.
// so long as they do not collide with other keys in the store.

// NextSequenceRecvKey returns the store key for the receive sequence of a particular
// channel binded to a specific port
func NextSequenceRecvKey(portID, channelID string) []byte { _ = "STUB: not implemented"; return nil }

// NextSequenceAckKey returns the store key for the acknowledgement sequence of
// a particular channel binded to a specific port.
func NextSequenceAckKey(portID, channelID string) []byte { _ = "STUB: not implemented"; return nil }

// PacketCommitmentKey returns the store key of under which a packet commitment
// is stored
func PacketCommitmentKey(portID, channelID string, sequence uint64) []byte {
	_ = "STUB: not implemented"
	return nil
}

// PacketCommitmentPrefixKey defines the prefix for commitments to packet data fields store path.
func PacketCommitmentPrefixKey(portID, channelID string) []byte {
	_ = "STUB: not implemented"
	return nil
}

// PacketAcknowledgementKey returns the store key of under which a packet
// acknowledgement is stored
func PacketAcknowledgementKey(portID, channelID string, sequence uint64) []byte {
	_ = "STUB: not implemented"
	return nil
}

// PacketAcknowledgementPrefixKey defines the prefix for commitments to packet data fields store path.
func PacketAcknowledgementPrefixKey(portID, channelID string) []byte {
	_ = "STUB: not implemented"
	return nil
}

// PacketReceiptKey returns the store key of under which a packet
// receipt is stored
func PacketReceiptKey(portID, channelID string, sequence uint64) []byte {
	_ = "STUB: not implemented"
	return nil
}

// RecvStartSequenceKey returns the store key for the recv start sequence of a particular channel
func RecvStartSequenceKey(portID, channelID string) []byte { _ = "STUB: not implemented"; return nil }

func sequencePath(sequence uint64) string { _ = "STUB: not implemented"; return "" }
