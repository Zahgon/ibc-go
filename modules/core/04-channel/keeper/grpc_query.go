package keeper

import (
	"context"

	"github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
)

var _ types.QueryServer = (*queryServer)(nil)

// queryServer implements the 04-channel types.QueryServer interface.
// It embeds the channel keeper to leverage store access while limiting the api of the channel keeper.
type queryServer struct {
	*Keeper
}

// NewQueryServer returns a new 04-channel types.QueryServer implementation.
func NewQueryServer(k *Keeper) types.QueryServer {
	_ = "STUB: not implemented"
	return *new(types.QueryServer)
}

// Channel implements the Query/Channel gRPC method
func (q *queryServer) Channel(goCtx context.Context, req *types.QueryChannelRequest) (*types.QueryChannelResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Channels implements the Query/Channels gRPC method
func (q *queryServer) Channels(goCtx context.Context, req *types.QueryChannelsRequest) (*types.QueryChannelsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConnectionChannels implements the Query/ConnectionChannels gRPC method
func (q *queryServer) ConnectionChannels(goCtx context.Context, req *types.QueryConnectionChannelsRequest) (*types.QueryConnectionChannelsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// filter any metadata stored under channel key

// ignore channel and continue to the next item if the connection is
// different than the requested one

// ChannelClientState implements the Query/ChannelClientState gRPC method
func (q *queryServer) ChannelClientState(goCtx context.Context, req *types.QueryChannelClientStateRequest) (*types.QueryChannelClientStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ChannelConsensusState implements the Query/ChannelConsensusState gRPC method
func (q *queryServer) ChannelConsensusState(goCtx context.Context, req *types.QueryChannelConsensusStateRequest) (*types.QueryChannelConsensusStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PacketCommitment implements the Query/PacketCommitment gRPC method
func (q *queryServer) PacketCommitment(goCtx context.Context, req *types.QueryPacketCommitmentRequest) (*types.QueryPacketCommitmentResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PacketCommitments implements the Query/PacketCommitments gRPC method
func (q *queryServer) PacketCommitments(goCtx context.Context, req *types.QueryPacketCommitmentsRequest) (*types.QueryPacketCommitmentsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PacketReceipt implements the Query/PacketReceipt gRPC method
func (q *queryServer) PacketReceipt(goCtx context.Context, req *types.QueryPacketReceiptRequest) (*types.QueryPacketReceiptResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PacketAcknowledgement implements the Query/PacketAcknowledgement gRPC method
func (q *queryServer) PacketAcknowledgement(goCtx context.Context, req *types.QueryPacketAcknowledgementRequest) (*types.QueryPacketAcknowledgementResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PacketAcknowledgements implements the Query/PacketAcknowledgements gRPC method
func (q *queryServer) PacketAcknowledgements(goCtx context.Context, req *types.QueryPacketAcknowledgementsRequest) (*types.QueryPacketAcknowledgementsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if a list of packet sequences is provided then query for each specific ack and return a list <= len(req.PacketCommitmentSequences)
// otherwise, maintain previous behaviour and perform paginated query

// UnreceivedPackets implements the Query/UnreceivedPackets gRPC method. Given
// a list of counterparty packet commitments, the querier checks if the packet
// has already been received by checking if a receipt exists on this
// chain for the packet sequence. All packets that haven't been received yet
// are returned in the response
// Usage: To use this method correctly, first query all packet commitments on
// the sending chain using the Query/PacketCommitments gRPC method.
// Then input the returned sequences into the QueryUnreceivedPacketsRequest
// and send the request to this Query/UnreceivedPackets on the **receiving**
// chain. This gRPC method will then return the list of packet sequences that
// are yet to be received on the receiving chain.
//
// NOTE: The querier makes the assumption that the provided list of packet
// commitments is correct and will not function properly if the list
// is not up to date. Ideally the query height should equal the latest height
// on the counterparty's client which represents this chain.
func (q *queryServer) UnreceivedPackets(goCtx context.Context, req *types.QueryUnreceivedPacketsRequest) (*types.QueryUnreceivedPacketsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// filter for invalid sequences to ensure they are not included in the response value.

// if the packet receipt does not exist, then it is unreceived

// filter for invalid sequences to ensure they are not included in the response value.

// Any sequence greater than or equal to the next sequence to be received is not received.

// UnreceivedAcks implements the Query/UnreceivedAcks gRPC method. Given
// a list of counterparty packet acknowledgements, the querier checks if the packet
// has already been received by checking if the packet commitment still exists on this
// chain (original sender) for the packet sequence.
// All acknowledgmeents that haven't been received yet are returned in the response.
// Usage: To use this method correctly, first query all packet acknowledgements on
// the original receiving chain (ie the chain that wrote the acks) using the Query/PacketAcknowledgements gRPC method.
// Then input the returned sequences into the QueryUnreceivedAcksRequest
// and send the request to this Query/UnreceivedAcks on the **original sending**
// chain. This gRPC method will then return the list of packet sequences whose
// acknowledgements are already written on the receiving chain but haven't yet
// been received back to the sending chain.
//
// NOTE: The querier makes the assumption that the provided list of packet
// acknowledgements is correct and will not function properly if the list
// is not up to date. Ideally the query height should equal the latest height
// on the counterparty's client which represents this chain.
func (q *queryServer) UnreceivedAcks(goCtx context.Context, req *types.QueryUnreceivedAcksRequest) (*types.QueryUnreceivedAcksResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if packet commitment still exists on the original sending chain, then packet ack has not been received
// since processing the ack will delete the packet commitment

// NextSequenceReceive implements the Query/NextSequenceReceive gRPC method
func (q *queryServer) NextSequenceReceive(goCtx context.Context, req *types.QueryNextSequenceReceiveRequest) (*types.QueryNextSequenceReceiveResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Return the next sequence received for ordered channels. Unordered channels
// do not make use of the next sequence receive.

// NextSequenceSend implements the Query/NextSequenceSend gRPC method
func (q *queryServer) NextSequenceSend(goCtx context.Context, req *types.QueryNextSequenceSendRequest) (*types.QueryNextSequenceSendResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
