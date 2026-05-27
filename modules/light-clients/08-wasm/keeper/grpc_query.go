package keeper

import (
	"context"

	"github.com/cosmos/ibc-go/modules/light-clients/08-wasm/v11/types"
)

var _ types.QueryServer = (*Keeper)(nil)

// Code implements the Query/Code gRPC method
func (k *Keeper) Code(goCtx context.Context, req *types.QueryCodeRequest) (*types.QueryCodeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Only return checksums we previously stored, not arbitrary checksums that might be stored via e.g Wasmd.

// Checksums implements the Query/Checksums gRPC method. It returns a list of hex encoded checksums stored.
func (k *Keeper) Checksums(goCtx context.Context, req *types.QueryChecksumsRequest) (*types.QueryChecksumsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
