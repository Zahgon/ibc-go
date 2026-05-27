package keeper

import (
	"context"

	"github.com/cosmos/ibc-go/v11/modules/apps/27-interchain-accounts/host/types"
)

var _ types.QueryServer = (*Keeper)(nil)

// Params implements the Query/Params gRPC method
func (k *Keeper) Params(goCtx context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
