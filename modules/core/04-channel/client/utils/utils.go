package utils

import (
	"github.com/cosmos/cosmos-sdk/client"

	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
)

// QueryChannel returns a channel end.
// If prove is true, it performs an ABCI store query in order to retrieve the merkle proof. Otherwise,
// it uses the gRPC query client.
func QueryChannel(
	clientCtx client.Context, portID, channelID string, prove bool,
) (*types.QueryChannelResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func queryChannelABCI(clientCtx client.Context, portID, channelID string) (*types.QueryChannelResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check if channel exists

// QueryChannelClientState returns the ClientState of a channel end. If
// prove is true, it performs an ABCI store query in order to retrieve the
// merkle proof. Otherwise, it uses the gRPC query client.
func QueryChannelClientState(
	clientCtx client.Context, portID, channelID string, prove bool,
) (*types.QueryChannelClientStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// use client state returned from ABCI query in case query height differs

// QueryChannelConsensusState returns the ConsensusState of a channel end. If
// prove is true, it performs an ABCI store query in order to retrieve the
// merkle proof. Otherwise, it uses the gRPC query client.
func QueryChannelConsensusState(
	clientCtx client.Context, portID, channelID string, height clienttypes.Height, prove bool,
) (*types.QueryChannelConsensusStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// QueryNextSequenceReceive returns the next sequence receive.
// If prove is true, it performs an ABCI store query in order to retrieve the merkle proof. Otherwise,
// it uses the gRPC query client.
func QueryNextSequenceReceive(
	clientCtx client.Context, portID, channelID string, prove bool,
) (*types.QueryNextSequenceReceiveResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func queryNextSequenceRecvABCI(clientCtx client.Context, portID, channelID string) (*types.QueryNextSequenceReceiveResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check if next sequence receive exists

// QueryNextSequenceSend returns the next sequence send.
// If prove is true, it performs an ABCI store query in order to retrieve the merkle proof. Otherwise,
// it uses the gRPC query client.
func QueryNextSequenceSend(
	clientCtx client.Context, portID, channelID string, prove bool,
) (*types.QueryNextSequenceSendResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func queryNextSequenceSendABCI(clientCtx client.Context, portID, channelID string) (*types.QueryNextSequenceSendResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check if next sequence send exists

// QueryPacketCommitment returns a packet commitment.
// If prove is true, it performs an ABCI store query in order to retrieve the merkle proof. Otherwise,
// it uses the gRPC query client.
func QueryPacketCommitment(
	clientCtx client.Context, portID, channelID string,
	sequence uint64, prove bool,
) (*types.QueryPacketCommitmentResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func queryPacketCommitmentABCI(
	clientCtx client.Context, portID, channelID string, sequence uint64,
) (*types.QueryPacketCommitmentResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check if packet commitment exists

// QueryPacketReceipt returns data about a packet receipt.
// If prove is true, it performs an ABCI store query in order to retrieve the merkle proof. Otherwise,
// it uses the gRPC query client.
func QueryPacketReceipt(
	clientCtx client.Context, portID, channelID string,
	sequence uint64, prove bool,
) (*types.QueryPacketReceiptResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func queryPacketReceiptABCI(
	clientCtx client.Context, portID, channelID string, sequence uint64,
) (*types.QueryPacketReceiptResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// QueryPacketAcknowledgement returns the data about a packet acknowledgement.
// If prove is true, it performs an ABCI store query in order to retrieve the merkle proof. Otherwise,
// it uses the gRPC query client
func QueryPacketAcknowledgement(clientCtx client.Context, portID, channelID string, sequence uint64, prove bool) (*types.QueryPacketAcknowledgementResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func queryPacketAcknowledgementABCI(clientCtx client.Context, portID, channelID string, sequence uint64) (*types.QueryPacketAcknowledgementResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
