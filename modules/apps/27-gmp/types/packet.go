package types

import (
	ibcexported "github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var (
	_ ibcexported.PacketData         = (*GMPPacketData)(nil)
	_ ibcexported.PacketDataProvider = (*GMPPacketData)(nil)
)

const (
	EncodingJSON     = "application/json"
	EncodingProtobuf = "application/x-protobuf"
	EncodingABI      = "application/x-solidity-abi"
)

// NewGMPPacketData creates a new GMPPacketData instance with the provided parameters.
func NewGMPPacketData(
	sender, receiver string, salt, payload []byte, memo string,
) GMPPacketData {
	_ = "STUB: not implemented"
	return *new(GMPPacketData)
}

func (p GMPPacketData) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// receiver is allowed to be empty

// MarshalPacketData attempts to marshal the provided GMPPacketData into bytes with the provided encoding.
func MarshalPacketData(data *GMPPacketData, ics27Version string, encoding string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalPacketData attempts to unmarshal the provided bytes into a GMPPacketData with the provided encoding.
func UnmarshalPacketData(bz []byte, ics27Version string, encoding string) (*GMPPacketData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPacketSender returns the sender address of the packet data.
// NOTE:
//   - The sender address is set by the module which requested the packet to be sent,
//     and this module may not have validated the sender address by a signature check.
//   - The sender address must only be used by modules on the sending chain.
//   - sourcePortID is not used in this implementation.
func (p GMPPacketData) GetPacketSender(sourcePortID string) string {
	_ = "STUB: not implemented"

	// GetCustomPacketData returns callback data for the callbacks middleware.
	// For "src_callback", returns sender as callback address (auto-registration).
	// For other keys, parses memo as JSON and returns the value for the given key.
	return ""
}

func (p GMPPacketData) GetCustomPacketData(key string) any {
	_ = "STUB: not implemented"
	return *new(any)
}
