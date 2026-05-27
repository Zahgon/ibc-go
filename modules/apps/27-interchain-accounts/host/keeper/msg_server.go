package keeper

import (
	"context"

	"github.com/cosmos/ibc-go/v11/modules/apps/27-interchain-accounts/host/types"
)

var _ types.MsgServer = (*msgServer)(nil)

type msgServer struct {
	*Keeper
}

// NewMsgServerImpl returns an implementation of the ICS27 host MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	_ = "STUB: not implemented"
	return *new(types.MsgServer)
}

// ModuleQuerySafe routes the queries to the keeper's query router if they are module_query_safe.
// This handler doesn't use the signer.
func (m msgServer) ModuleQuerySafe(goCtx context.Context, msg *types.MsgModuleQuerySafe) (*types.MsgModuleQuerySafeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateParams updates the host submodule's params.
func (m msgServer) UpdateParams(goCtx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
