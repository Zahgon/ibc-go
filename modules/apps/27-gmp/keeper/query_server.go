package keeper

import (
	"context"

	"github.com/cosmos/ibc-go/v11/modules/apps/27-gmp/types"
)

var _ types.QueryServer = (*Keeper)(nil)

// AccountAddress defines the handler for the Query/AccountAddress RPC method.
func (k *Keeper) AccountAddress(ctx context.Context, req *types.QueryAccountAddressRequest) (*types.QueryAccountAddressResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AccountIdentifier defines the handler for the Query/AccountIdentifier RPC method.
func (k *Keeper) AccountIdentifier(ctx context.Context, req *types.QueryAccountIdentifierRequest) (*types.QueryAccountIdentifierResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
