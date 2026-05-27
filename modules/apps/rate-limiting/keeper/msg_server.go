package keeper

import (
	"context"

	"github.com/cosmos/ibc-go/v11/modules/apps/rate-limiting/types"
)

type msgServer struct {
	*Keeper
}

// NewMsgServerImpl returns an implementation of the ratelimit MsgServer interface
func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	_ = "STUB: not implemented"
	return *new(types.MsgServer)
}

var _ types.MsgServer = msgServer{}

// Adds a new rate limit. Fails if the rate limit already exists or the channel value is 0
func (k msgServer) AddRateLimit(goCtx context.Context, msg *types.MsgAddRateLimit) (*types.MsgAddRateLimitResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates an existing rate limit. Fails if the rate limit doesn't exist
func (k msgServer) UpdateRateLimit(goCtx context.Context, msg *types.MsgUpdateRateLimit) (*types.MsgUpdateRateLimitResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Removes a rate limit. Fails if the rate limit doesn't exist
func (k msgServer) RemoveRateLimit(goCtx context.Context, msg *types.MsgRemoveRateLimit) (*types.MsgRemoveRateLimitResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Resets the flow on a rate limit. Fails if the rate limit doesn't exist
func (k msgServer) ResetRateLimit(goCtx context.Context, msg *types.MsgResetRateLimit) (*types.MsgResetRateLimitResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
