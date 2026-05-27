package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var (
	_ codectypes.UnpackInterfacesMessage = (*QueryChannelClientStateResponse)(nil)
	_ codectypes.UnpackInterfacesMessage = (*QueryChannelConsensusStateResponse)(nil)
)

// NewQueryChannelResponse creates a new QueryChannelResponse instance
func NewQueryChannelResponse(channel Channel, proof []byte, height clienttypes.Height) *QueryChannelResponse {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryChannelClientStateResponse creates a newQueryChannelClientStateResponse instance
func NewQueryChannelClientStateResponse(identifiedClientState clienttypes.IdentifiedClientState, proof []byte, height clienttypes.Height) *QueryChannelClientStateResponse {
	_ = "STUB: not implemented"
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (resp QueryChannelClientStateResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryChannelConsensusStateResponse creates a newQueryChannelConsensusStateResponse instance
func NewQueryChannelConsensusStateResponse(clientID string, anyConsensusState *codectypes.Any, consensusStateHeight exported.Height, proof []byte, height clienttypes.Height) *QueryChannelConsensusStateResponse {
	_ = "STUB: not implemented"
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (resp QueryChannelConsensusStateResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryPacketCommitmentResponse creates a new QueryPacketCommitmentResponse instance
func NewQueryPacketCommitmentResponse(
	commitment []byte, proof []byte, height clienttypes.Height,
) *QueryPacketCommitmentResponse {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryPacketReceiptResponse creates a new QueryPacketReceiptResponse instance
func NewQueryPacketReceiptResponse(
	recvd bool, proof []byte, height clienttypes.Height,
) *QueryPacketReceiptResponse {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryPacketAcknowledgementResponse creates a new QueryPacketAcknowledgementResponse instance
func NewQueryPacketAcknowledgementResponse(
	acknowledgement []byte, proof []byte, height clienttypes.Height,
) *QueryPacketAcknowledgementResponse {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryNextSequenceReceiveResponse creates a new QueryNextSequenceReceiveResponse instance
func NewQueryNextSequenceReceiveResponse(
	sequence uint64, proof []byte, height clienttypes.Height,
) *QueryNextSequenceReceiveResponse {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryNextSequenceSendResponse creates a new QueryNextSequenceSendResponse instance
func NewQueryNextSequenceSendResponse(
	sequence uint64, proof []byte, height clienttypes.Height,
) *QueryNextSequenceSendResponse {
	_ = "STUB: not implemented"
	return nil
}
