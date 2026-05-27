package types

import (
	"time"

	ibcexported "github.com/cosmos/ibc-go/v11/modules/core/exported"
)

// PacketMetadata represents the metadata for a packet in the packet-forward middleware.
// Do not use this type directly with json encoding/decoding, as it is not json serializable.
// Instead use the provided helper methods to convert it, or use the Metadata keys defined in this package.
type PacketMetadata struct {
	Forward ForwardMetadata
}

// ForwardMetadata represents the metadata for forwarding a packet.
// Do not use this type directly with json encoding/decoding, as it is not json serializable.
// Instead use the provided helper methods to convert it, or use the Metadata keys defined in this package.
type ForwardMetadata struct {
	Receiver string
	Port     string
	Channel  string
	Timeout  time.Duration
	Retries  *uint8

	Next *PacketMetadata // Next is a pointer to allow nil values
}

func (m ForwardMetadata) Validate() error { _ = "STUB: not implemented"; return nil }

func (m ForwardMetadata) ToMap() map[string]any { _ = "STUB: not implemented"; return nil }

func (m PacketMetadata) toMap() map[string]any { _ = "STUB: not implemented"; return nil }

func (m PacketMetadata) ToMemo() (string, error) { _ = "STUB: not implemented"; return "", nil }

func GetPacketMetadataFromPacketdata(transferDetail ibcexported.PacketDataProvider) (PacketMetadata, bool, error) {
	_ = "STUB: not implemented"
	return *new(PacketMetadata), false, nil
}

func getForwardMetadata(forwardData map[string]any) (ForwardMetadata, error) {
	_ = "STUB: not implemented"
	return *new(ForwardMetadata), nil
}

func getForwardMetadataFromNext(nextData any) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseDuration(duration any) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}
