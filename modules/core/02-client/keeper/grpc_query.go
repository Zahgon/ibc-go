package keeper

import (
	"context"

	"github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
)

var _ types.QueryServer = (*queryServer)(nil)

// queryServer implements the 02-client types.QueryServer interface.
// It embeds the client keeper to leverage store access while limiting the api of the client keeper.
type queryServer struct {
	*Keeper
}

// NewQueryServer returns a new 02-client types.QueryServer implementation.
func NewQueryServer(k *Keeper) types.QueryServer {
	_ = "STUB: not implemented"
	return *new(types.QueryServer)
}

// ClientState implements the Query/ClientState gRPC method
func (q *queryServer) ClientState(goCtx context.Context, req *types.QueryClientStateRequest) (*types.QueryClientStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ClientStates implements the Query/ClientStates gRPC method
func (q *queryServer) ClientStates(goCtx context.Context, req *types.QueryClientStatesRequest) (*types.QueryClientStatesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// filter any metadata stored under client state key

// ConsensusState implements the Query/ConsensusState gRPC method
func (q *queryServer) ConsensusState(goCtx context.Context, req *types.QueryConsensusStateRequest) (*types.QueryConsensusStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConsensusStates implements the Query/ConsensusStates gRPC method
func (q *queryServer) ConsensusStates(goCtx context.Context, req *types.QueryConsensusStatesRequest) (*types.QueryConsensusStatesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// filter any metadata stored under consensus state key

// ConsensusStateHeights implements the Query/ConsensusStateHeights gRPC method
func (q *queryServer) ConsensusStateHeights(goCtx context.Context, req *types.QueryConsensusStateHeightsRequest) (*types.QueryConsensusStateHeightsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// filter any metadata stored under consensus state key

// ClientStatus implements the Query/ClientStatus gRPC method
func (q *queryServer) ClientStatus(goCtx context.Context, req *types.QueryClientStatusRequest) (*types.QueryClientStatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ClientCreator implements the Query/ClientCreator gRPC method
func (q *queryServer) ClientCreator(goCtx context.Context, req *types.QueryClientCreatorRequest) (*types.QueryClientCreatorResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ClientParams implements the Query/ClientParams gRPC method
func (q *queryServer) ClientParams(goCtx context.Context, _ *types.QueryClientParamsRequest) (*types.QueryClientParamsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpgradedClientState implements the Query/UpgradedClientState gRPC method
func (q *queryServer) UpgradedClientState(goCtx context.Context, req *types.QueryUpgradedClientStateRequest) (*types.QueryUpgradedClientStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpgradedConsensusState implements the Query/UpgradedConsensusState gRPC method
func (q *queryServer) UpgradedConsensusState(goCtx context.Context, req *types.QueryUpgradedConsensusStateRequest) (*types.QueryUpgradedConsensusStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// VerifyMembership implements the Query/VerifyMembership gRPC method
// NOTE: Any state changes made within this handler are discarded by leveraging a cached context. Gas is consumed for underlying state access.
// This gRPC method is intended to be used within the context of the state machine and delegates to light clients to verify proofs.
func (q *queryServer) VerifyMembership(goCtx context.Context, req *types.QueryVerifyMembershipRequest) (*types.QueryVerifyMembershipResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// cache the context to ensure clientState.VerifyMembership does not change state

// make sure we charge the higher level context even on panic

// consume flat gas fee for proof verification queries.
// NOTE: consuming gas prior to method invocation also provides protection against recursive calls reaching stack overflow
