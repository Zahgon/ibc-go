package keeper

import (
	"context"

	"github.com/cosmos/ibc-go/modules/light-clients/08-wasm/v11/types"
)

var _ types.MsgServer = (*Keeper)(nil)

// StoreCode defines a rpc handler method for MsgStoreCode
func (k *Keeper) StoreCode(goCtx context.Context, msg *types.MsgStoreCode) (*types.MsgStoreCodeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RemoveChecksum defines a rpc handler method for MsgRemoveChecksum
func (k *Keeper) RemoveChecksum(goCtx context.Context, msg *types.MsgRemoveChecksum) (*types.MsgRemoveChecksumResponse,
	error,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

// unpin the code from the vm in-memory cache

// MigrateContract defines a rpc handler method for MsgMigrateContract
func (k *Keeper) MigrateContract(goCtx context.Context, msg *types.MsgMigrateContract) (*types.MsgMigrateContractResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// event emission is handled in migrateContractCode
