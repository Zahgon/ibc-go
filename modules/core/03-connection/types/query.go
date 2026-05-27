package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var (
	_ codectypes.UnpackInterfacesMessage = (*QueryConnectionClientStateResponse)(nil)
	_ codectypes.UnpackInterfacesMessage = (*QueryConnectionConsensusStateResponse)(nil)
)

// NewQueryConnectionResponse creates a new QueryConnectionResponse instance
func NewQueryConnectionResponse(
	connection ConnectionEnd, proof []byte, height clienttypes.Height,
) *QueryConnectionResponse {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryClientConnectionsResponse creates a new ConnectionPaths instance
func NewQueryClientConnectionsResponse(
	connectionPaths []string, proof []byte, height clienttypes.Height,
) *QueryClientConnectionsResponse {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryClientConnectionsRequest creates a new QueryClientConnectionsRequest instance
func NewQueryClientConnectionsRequest(clientID string) *QueryClientConnectionsRequest {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryConnectionClientStateResponse creates a newQueryConnectionClientStateResponse instance
func NewQueryConnectionClientStateResponse(identifiedClientState clienttypes.IdentifiedClientState, proof []byte, height clienttypes.Height) *QueryConnectionClientStateResponse {
	_ = "STUB: not implemented"
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (resp QueryConnectionClientStateResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryConnectionConsensusStateResponse creates a newQueryConnectionConsensusStateResponse instance
func NewQueryConnectionConsensusStateResponse(clientID string, anyConsensusState *codectypes.Any, consensusStateHeight exported.Height, proof []byte, height clienttypes.Height) *QueryConnectionConsensusStateResponse {
	_ = "STUB: not implemented"
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (resp QueryConnectionConsensusStateResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}
