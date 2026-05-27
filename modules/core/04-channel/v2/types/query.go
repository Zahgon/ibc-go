package types

import (
	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
)

// NewQueryNextSequenceSendRequest creates a new next sequence send query.
func NewQueryNextSequenceSendRequest(clientID string) *QueryNextSequenceSendRequest {
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

// NewQueryPacketCommitmentRequest creates and returns a new packet commitment query request.
func NewQueryPacketCommitmentRequest(clientID string, sequence uint64) *QueryPacketCommitmentRequest {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryPacketCommitmentResponse creates and returns a new packet commitment query response.
func NewQueryPacketCommitmentResponse(commitmentHash []byte, proof []byte, proofHeight clienttypes.Height) *QueryPacketCommitmentResponse {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryPacketAcknowledgementRequest creates and returns a new packet acknowledgement query request.
func NewQueryPacketAcknowledgementRequest(clientID string, sequence uint64) *QueryPacketAcknowledgementRequest {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryPacketAcknowledgementResponse creates and returns a new packet acknowledgement query response.
func NewQueryPacketAcknowledgementResponse(acknowledgementHash []byte, proof []byte, proofHeight clienttypes.Height) *QueryPacketAcknowledgementResponse {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryPacketReceiptRequest creates and returns a new packet receipt query request.
func NewQueryPacketReceiptRequest(clientID string, sequence uint64) *QueryPacketReceiptRequest {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryPacketReceiptResponse creates and returns a new packet receipt query response.
func NewQueryPacketReceiptResponse(exists bool, proof []byte, height clienttypes.Height) *QueryPacketReceiptResponse {
	_ = "STUB: not implemented"
	return nil
}

// NewQueryPacketReceiptRequest creates and returns a new packet receipt query request.
func NewQueryUnreceivedPacketsRequest(clientID string, sequences []uint64) *QueryUnreceivedPacketsRequest {
	_ = "STUB: not implemented"
	return nil
}
