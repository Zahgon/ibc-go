package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

var (
	_ codectypes.UnpackInterfacesMessage = (*QueryClientStateResponse)(nil)
	_ codectypes.UnpackInterfacesMessage = (*QueryClientStatesResponse)(nil)
	_ codectypes.UnpackInterfacesMessage = (*QueryConsensusStateResponse)(nil)
	_ codectypes.UnpackInterfacesMessage = (*QueryConsensusStatesResponse)(nil)
)

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (qcsr QueryClientStatesResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryClientStateResponse creates a new QueryClientStateResponse instance.
func NewQueryClientStateResponse(
	clientStateAny *codectypes.Any, proof []byte, height Height,
) *QueryClientStateResponse {
	_ = "STUB: not implemented"
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (qcsr QueryClientStateResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (qcsr QueryConsensusStatesResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryConsensusStateResponse creates a new QueryConsensusStateResponse instance.
func NewQueryConsensusStateResponse(
	consensusStateAny *codectypes.Any, proof []byte, height Height,
) *QueryConsensusStateResponse {
	_ = "STUB: not implemented"
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (qcsr QueryConsensusStateResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}
