package keeper

import (
	"context"

	"github.com/cosmos/ibc-go/v11/modules/core/02-client/v2/types"
)

var _ types.QueryServer = (*queryServer)(nil)

// queryServer implements the 02-client/v2 types.QueryServer interface.
// It embeds the client keeper to leverage store access while limiting the api of the client keeper.
type queryServer struct {
	*Keeper
}

// NewQueryServer returns a new 02-client/v2 types.QueryServer implementation.
func NewQueryServer(k *Keeper) types.QueryServer {
	_ = "STUB: not implemented"
	return *new(types.QueryServer)
}

// CounterpartyInfo gets the CounterpartyInfo from the store corresponding to the request client ID.
func (q queryServer) CounterpartyInfo(goCtx context.Context, req *types.QueryCounterpartyInfoRequest) (*types.QueryCounterpartyInfoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Config queries the configuration of the ibc client v2 module.
func (q queryServer) Config(goCtx context.Context, req *types.QueryConfigRequest) (*types.QueryConfigResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
