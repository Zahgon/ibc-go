package types

import (
	ibcexported "github.com/cosmos/ibc-go/v11/modules/core/exported"
)

// InternalTransferRepresentation defines a struct used internally by the transfer application to represent a fungible token transfer
type InternalTransferRepresentation struct {
	// the tokens to be transferred
	Token Token
	// the sender address
	Sender string
	// the recipient address on the destination chain
	Receiver string
	// optional memo
	Memo string
}

var (
	_ ibcexported.PacketData         = (*FungibleTokenPacketData)(nil)
	_ ibcexported.PacketDataProvider = (*FungibleTokenPacketData)(nil)
)

const (
	EncodingJSON     = "application/json"
	EncodingProtobuf = "application/x-protobuf"
	EncodingABI      = "application/x-solidity-abi"
)

// NewFungibleTokenPacketData constructs a new FungibleTokenPacketData instance
func NewFungibleTokenPacketData(
	denom string, amount string,
	sender, receiver string,
	memo string,
) FungibleTokenPacketData {
	_ = "STUB: not implemented"
	return *new(FungibleTokenPacketData)
}

// ValidateBasic is used for validating the token transfer.
// NOTE: The addresses formats are not validated as the sender and recipient can have different
// formats defined by their corresponding chains that are not known to IBC.
func (ftpd FungibleTokenPacketData) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// GetBytes is a helper for serialising the packet to bytes.
// The memo field of FungibleTokenPacketData is marked with the JSON omitempty tag
// ensuring that the memo field is not included in the marshalled bytes if one is not specified.
func (ftpd FungibleTokenPacketData) GetBytes() []byte { _ = "STUB: not implemented"; return nil }

// GetPacketSender returns the sender address embedded in the packet data.
//
// NOTE:
//   - The sender address is set by the module which requested the packet to be sent,
//     and this module may not have validated the sender address by a signature check.
//   - The sender address must only be used by modules on the sending chain.
//   - sourcePortID is not used in this implementation.
func (ftpd FungibleTokenPacketData) GetPacketSender(sourcePortID string) string {
	_ = "STUB: not implemented"
	return ""

	// GetCustomPacketData interprets the memo field of the packet data as a JSON object
	// and returns the value associated with the given key.
	// If the key is missing or the memo is not properly formatted, then nil is returned.
}

func (ftpd FungibleTokenPacketData) GetCustomPacketData(key string) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// NewInternalTransferRepresentation constructs a new InternalTransferRepresentation instance
func NewInternalTransferRepresentation(
	token Token,
	sender, receiver string,
	memo string,
) InternalTransferRepresentation {
	_ = "STUB: not implemented"
	return *new(InternalTransferRepresentation)
}

// ValidateBasic is used for validating the token transfer.
// NOTE: The addresses formats are not validated as the sender and recipient can have different
// formats defined by their corresponding chains that are not known to IBC.
func (ftpd InternalTransferRepresentation) ValidateBasic() error {
	_ = "STUB: not implemented"
	return nil
}

// GetCustomPacketData interprets the memo field of the packet data as a JSON object
// and returns the value associated with the given key.
// If the key is missing or the memo is not properly formatted, then nil is returned.
func (ftpd InternalTransferRepresentation) GetCustomPacketData(key string) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// GetPacketSender returns the sender address embedded in the packet data.
//
// NOTE:
//   - The sender address is set by the module which requested the packet to be sent,
//     and this module may not have validated the sender address by a signature check.
//   - The sender address must only be used by modules on the sending chain.
//   - sourcePortID is not used in this implementation.
func (ftpd InternalTransferRepresentation) GetPacketSender(sourcePortID string) string {
	_ = "STUB: not implemented"
	return ""

	// MarshalPacketData attempts to marshal the provided FungibleTokenPacketData into bytes with the provided encoding.
}

func MarshalPacketData(data FungibleTokenPacketData, ics20Version string, encoding string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalPacketData attempts to unmarshal the provided packet data bytes into a InternalTransferRepresentation.
func UnmarshalPacketData(bz []byte, ics20Version string, encoding string) (InternalTransferRepresentation, error) {
	_ = "STUB: not implemented"
	return *new(InternalTransferRepresentation), nil
}

// Depending on the ics20 version, we use a different default encoding (json for V1, proto for V2)
// and we have a different type to unmarshal the data into.

// Here we perform the unmarshaling based on the specified encoding.
// The functions act on the generic "data" variable which is of type proto.Message (an interface).

// When the unmarshaling is done, we want to retrieve the underlying data type based on the value of ics20Version
// Since it has to be v1, we convert the data to FungibleTokenPacketData and then call the conversion function to construct
// the v2 type.

// We should never get here, as we manually constructed the type at the beginning of the file

// The call to ValidateBasic for V1 is done inside PacketDataV1toV2.

// PacketDataV1ToV2 converts a v1 packet data to a v2 packet data. The packet data is validated
// before conversion.
func PacketDataV1ToV2(packetData FungibleTokenPacketData) (InternalTransferRepresentation, error) {
	_ = "STUB: not implemented"
	return *new(InternalTransferRepresentation), nil
}
