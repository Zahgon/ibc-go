package keeper

import (
	"context"

	"github.com/cosmos/ibc-go/v11/modules/apps/27-interchain-accounts/controller/types"
)

var _ types.MsgServer = (*msgServer)(nil)

type msgServer struct {
	*Keeper
}

// NewMsgServerImpl returns an implementation of the ICS27 MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	_ = "STUB: not implemented"
	return *new(types.MsgServer)
}

// RegisterInterchainAccount defines a rpc handler for MsgRegisterInterchainAccount
func (s msgServer) RegisterInterchainAccount(goCtx context.Context, msg *types.MsgRegisterInterchainAccount) (*types.MsgRegisterInterchainAccountResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// use ORDER_UNORDERED as default in case msg's ordering is NONE

// SendTx defines a rpc handler for MsgSendTx
func (s msgServer) SendTx(goCtx context.Context, msg *types.MsgSendTx) (*types.MsgSendTxResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// the absolute timeout value is calculated using the controller chain block time + the relative timeout value
// this assumes time synchrony to a certain degree between the controller and counterparty host chain

// UpdateParams defines an rpc handler method for MsgUpdateParams. Updates the ica/controller submodule's parameters.
func (k *Keeper) UpdateParams(goCtx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
