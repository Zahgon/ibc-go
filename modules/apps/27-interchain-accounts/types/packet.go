package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	ibcexported "github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var (
	_ ibcexported.PacketData         = (*InterchainAccountPacketData)(nil)
	_ ibcexported.PacketDataProvider = (*InterchainAccountPacketData)(nil)
)

// MaxMemoCharLength defines the maximum length for the InterchainAccountPacketData memo field
const MaxMemoCharLength = 32768

// ValidateBasic performs basic validation of the interchain account packet data.
// The memo may be empty.
func (iapd InterchainAccountPacketData) ValidateBasic() error {
	_ = "STUB: not implemented"
	return nil
}

// GetBytes returns the JSON marshalled interchain account packet data.
func (iapd InterchainAccountPacketData) GetBytes() []byte { _ = "STUB: not implemented"; return nil }

// UnmarshalJSON unmarshals raw JSON bytes into an InterchainAccountPacketData.
func (iapd *InterchainAccountPacketData) UnmarshalJSON(bz []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (ct CosmosTx) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// GetPacketSender returns the sender address of the interchain accounts packet data.
// It is obtained from the source port ID by cutting off the ControllerPortPrefix.
// If the source port ID does not have the ControllerPortPrefix, then an empty string is returned.
//
// NOTE:
//   - The sender address is set by the packet sender and may not have been validated a signature
//     check if the packet sender isn't the interchain accounts module.
//   - The sender address must only be used by modules on the sending chain.
func (InterchainAccountPacketData) GetPacketSender(sourcePortID string) string {
	_ = "STUB: not implemented"
	return ""
}

// GetCustomPacketData interprets the memo field of the packet data as a JSON object
// and returns the value associated with the given key.
// If the key is missing or the memo is not properly formatted, then nil is returned.
func (iapd InterchainAccountPacketData) GetCustomPacketData(key string) any {
	_ = "STUB: not implemented"
	return *new(any)
}
