package solomachine

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

// Interface implementation checks.
var _, _, _, _ codectypes.UnpackInterfacesMessage = (*ClientState)(nil), (*ConsensusState)(nil), (*Header)(nil), (*HeaderData)(nil)

// Data is an interface used for all the signature data bytes proto definitions.
type Data any

// UnpackInterfaces implements the UnpackInterfaceMessages.UnpackInterfaces method
func (cs ClientState) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// UnpackInterfaces implements the UnpackInterfaceMessages.UnpackInterfaces method
func (cs ConsensusState) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// UnpackInterfaces implements the UnpackInterfaceMessages.UnpackInterfaces method
func (h Header) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// UnpackInterfaces implements the UnpackInterfaceMessages.UnpackInterfaces method
func (hd HeaderData) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}
