package types

import (
	corestore "cosmossdk.io/core/store"

	storetypes "github.com/cosmos/cosmos-sdk/store/v2/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// StoreProvider encapsulates the IBC core store service and offers convenience methods for LightClientModules.
type StoreProvider struct {
	storeService corestore.KVStoreService
}

// NewStoreProvider creates and returns a new client StoreProvider.
func NewStoreProvider(storeService corestore.KVStoreService) StoreProvider {
	_ = "STUB: not implemented"
	return *new(StoreProvider)
}

// ClientStore returns isolated prefix store for each client so they can read/write in separate namespaces.
func (s StoreProvider) ClientStore(ctx sdk.Context, clientID string) storetypes.KVStore {
	_ = "STUB: not implemented"
	return *new(storetypes.KVStore)
}

// ClientModuleStore returns the module store for a provided client type.
func (s StoreProvider) ClientModuleStore(ctx sdk.Context, clientType string) storetypes.KVStore {
	_ = "STUB: not implemented"
	return *new(storetypes.KVStore)
}
