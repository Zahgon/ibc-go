package v2

const (
	PacketCommitmentBasePrefix      = byte(1)
	PacketReceiptBasePrefix         = byte(2)
	PacketAcknowledgementBasePrefix = byte(3)
	KeyNextSeqSendPrefix            = "nextSequenceSend/"
)

// PacketCommitmentPrefixKey returns the store key prefix under which packet commitments for a particular channel are stored.
// channelID must be a generated identifier, not provided externally so key collisions are not possible.
func PacketCommitmentPrefixKey(channelID string) []byte { _ = "STUB: not implemented"; return nil }

// PacketCommitmentKey returns the store key of under which a packet commitment is stored.
// channelID must be a generated identifier, not provided externally so key collisions are not possible.
func PacketCommitmentKey(channelID string, sequence uint64) []byte {
	_ = "STUB: not implemented"
	return nil
}

// PacketReceiptPrefixKey returns the store key prefix under which packet receipts for a particular channel are stored.
// channelID must be a generated identifier, not provided externally so key collisions are not possible.
func PacketReceiptPrefixKey(channelID string) []byte { _ = "STUB: not implemented"; return nil }

// PacketReceiptKey returns the store key of under which a packet receipt is stored.
// channelID must be a generated identifier, not provided externally so key collisions are not possible.
func PacketReceiptKey(channelID string, sequence uint64) []byte {
	_ = "STUB: not implemented"
	return nil
}

// PacketAcknowledgementPrefixKey returns the store key prefix under which packet acknowledgements for a particular channel are stored.
// channelID must be a generated identifier, not provided externally so key collisions are not possible.
func PacketAcknowledgementPrefixKey(channelID string) []byte { _ = "STUB: not implemented"; return nil }

// PacketAcknowledgementKey returns the store key of under which a packet acknowledgement is stored.
// channelID must be a generated identifier, not provided externally so key collisions are not possible.
func PacketAcknowledgementKey(channelID string, sequence uint64) []byte {
	_ = "STUB: not implemented"
	return nil
}

// NextSequenceSendKey returns the store key for the next sequence send of a given channelID.
func NextSequenceSendKey(channelID string) []byte { _ = "STUB: not implemented"; return nil }
