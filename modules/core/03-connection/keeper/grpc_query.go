package keeper

import (
	"context"

	"github.com/cosmos/ibc-go/v11/modules/core/03-connection/types"
)

var _ types.QueryServer = (*queryServer)(nil)

// queryServer implements the 03-connection types.QueryServer interface.
// It embeds the connection keeper to leverage store access while limiting the api of the connection keeper.
type queryServer struct {
	*Keeper
}

// NewQueryServer returns a new 03-connection types.QueryServer implementation.
func NewQueryServer(k *Keeper) types.QueryServer {
	_ = "STUB: not implemented"
	return *new(types.QueryServer)
}

// Connection implements the Query/Connection gRPC method
func (q *queryServer) Connection(goCtx context.Context, req *types.QueryConnectionRequest) (*types.QueryConnectionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Connections implements the Query/Connections gRPC method
func (q *queryServer) Connections(goCtx context.Context, req *types.QueryConnectionsRequest) (*types.QueryConnectionsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ClientConnections implements the Query/ClientConnections gRPC method
func (q *queryServer) ClientConnections(goCtx context.Context, req *types.QueryClientConnectionsRequest) (*types.QueryClientConnectionsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConnectionClientState implements the Query/ConnectionClientState gRPC method
func (q *queryServer) ConnectionClientState(goCtx context.Context, req *types.QueryConnectionClientStateRequest) (*types.QueryConnectionClientStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConnectionConsensusState implements the Query/ConnectionConsensusState gRPC method
func (q *queryServer) ConnectionConsensusState(goCtx context.Context, req *types.QueryConnectionConsensusStateRequest) (*types.QueryConnectionConsensusStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConnectionParams implements the Query/ConnectionParams gRPC method.
func (q *queryServer) ConnectionParams(goCtx context.Context, req *types.QueryConnectionParamsRequest) (*types.QueryConnectionParamsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
