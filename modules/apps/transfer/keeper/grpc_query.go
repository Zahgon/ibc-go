package keeper

import (
	"context"

	"github.com/cosmos/ibc-go/v11/modules/apps/transfer/types"
)

var _ types.QueryServer = (*Keeper)(nil)

// Denom implements the Query/Denom gRPC method
func (k *Keeper) Denom(goCtx context.Context, req *types.QueryDenomRequest) (*types.QueryDenomResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Denoms implements the Query/Denoms gRPC method
func (k *Keeper) Denoms(ctx context.Context, req *types.QueryDenomsRequest) (*types.QueryDenomsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Params implements the Query/Params gRPC method
func (k *Keeper) Params(goCtx context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DenomHash implements the Query/DenomHash gRPC method
func (k *Keeper) DenomHash(goCtx context.Context, req *types.QueryDenomHashRequest) (*types.QueryDenomHashResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Convert given request trace path to Denom struct to confirm the path in a valid denom trace format

// EscrowAddress implements the EscrowAddress gRPC method
func (k *Keeper) EscrowAddress(goCtx context.Context, req *types.QueryEscrowAddressRequest) (*types.QueryEscrowAddressResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TotalEscrowForDenom implements the TotalEscrowForDenom gRPC method.
func (k *Keeper) TotalEscrowForDenom(goCtx context.Context, req *types.QueryTotalEscrowForDenomRequest) (*types.QueryTotalEscrowForDenomResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
