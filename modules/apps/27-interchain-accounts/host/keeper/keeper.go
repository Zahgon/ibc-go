package keeper

import (
	corestore "cosmossdk.io/core/store"
	"cosmossdk.io/log/v2"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	genesistypes "github.com/cosmos/ibc-go/v11/modules/apps/27-interchain-accounts/genesis/types"
	"github.com/cosmos/ibc-go/v11/modules/apps/27-interchain-accounts/host/types"
	icatypes "github.com/cosmos/ibc-go/v11/modules/apps/27-interchain-accounts/types"
	porttypes "github.com/cosmos/ibc-go/v11/modules/core/05-port/types"
)

// Keeper defines the IBC interchain accounts host keeper
type Keeper struct {
	storeService corestore.KVStoreService
	cdc          codec.Codec

	ics4Wrapper   porttypes.ICS4Wrapper
	channelKeeper icatypes.ChannelKeeper
	accountKeeper icatypes.AccountKeeper

	msgRouter   icatypes.MessageRouter
	queryRouter icatypes.QueryRouter

	// mqsAllowList is a list of all module safe query paths
	mqsAllowList []string

	// the address capable of executing a MsgUpdateParams message. Typically, this
	// should be the x/gov module account.
	authority string
}

// NewKeeper creates a new interchain accounts host Keeper instance
func NewKeeper(
	cdc codec.Codec, storeService corestore.KVStoreService,
	channelKeeper icatypes.ChannelKeeper,
	accountKeeper icatypes.AccountKeeper, msgRouter icatypes.MessageRouter, queryRouter icatypes.QueryRouter, authority string,
) *Keeper {
	_ = "STUB: not implemented"
	// ensure ibc interchain accounts module account is set
	return nil
}

// Defaults to using the channel keeper as the ICS4Wrapper
// This can be overridden later with WithICS4Wrapper (e.g. by the middleware stack wiring)

// WithICS4Wrapper sets the ICS4Wrapper. This function may be used after
// the keepers creation to set the middleware which is above this module
// in the IBC application stack.
func (k *Keeper) WithICS4Wrapper(wrapper porttypes.ICS4Wrapper) { _ = "STUB: not implemented"; return }

// GetICS4Wrapper returns the ICS4Wrapper.
func (k *Keeper) GetICS4Wrapper() porttypes.ICS4Wrapper {
	_ = "STUB: not implemented"
	return *

	// Logger returns the application logger, scoped to the associated module
	new(porttypes.ICS4Wrapper)
}

func (*Keeper) Logger(ctx sdk.Context) log.Logger {
	_ = "STUB: not implemented"
	return *new(log.Logger)
}

// setPort sets the provided portID in state.
func (k *Keeper) setPort(ctx sdk.Context, portID string) { _ = "STUB: not implemented"; return }

// GetAppVersion calls the ICS4Wrapper GetAppVersion function.
func (k *Keeper) GetAppVersion(ctx sdk.Context, portID, channelID string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// getAppMetadata retrieves the interchain accounts channel metadata from the store associated with the provided portID and channelID
func (k *Keeper) getAppMetadata(ctx sdk.Context, portID, channelID string) (icatypes.Metadata, error) {
	_ = "STUB: not implemented"
	return *new(icatypes.Metadata), nil
}

// GetActiveChannelID retrieves the active channelID from the store keyed by the provided connectionID and portID
func (k *Keeper) GetActiveChannelID(ctx sdk.Context, connectionID, portID string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// GetOpenActiveChannel retrieves the active channelID from the store, keyed by the provided connectionID and portID & checks if the channel in question is in state OPEN
func (k *Keeper) GetOpenActiveChannel(ctx sdk.Context, connectionID, portID string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// GetAllActiveChannels returns a list of all active interchain accounts host channels and their associated connection and port identifiers
func (k *Keeper) GetAllActiveChannels(ctx sdk.Context) []genesistypes.ActiveChannel {
	_ = "STUB: not implemented"
	return nil
}

// SetActiveChannelID stores the active channelID, keyed by the provided connectionID and portID
func (k *Keeper) SetActiveChannelID(ctx sdk.Context, connectionID, portID, channelID string) {
	_ = "STUB: not implemented"
	return
}

// IsActiveChannel returns true if there exists an active channel for the provided connectionID and portID, otherwise false
func (k *Keeper) IsActiveChannel(ctx sdk.Context, connectionID, portID string) bool {
	_ = "STUB: not implemented"
	return false
}

// GetInterchainAccountAddress retrieves the InterchainAccount address from the store associated with the provided connectionID and portID
func (k *Keeper) GetInterchainAccountAddress(ctx sdk.Context, connectionID, portID string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// GetAllInterchainAccounts returns a list of all registered interchain account addresses and their associated connection and controller port identifiers
func (k *Keeper) GetAllInterchainAccounts(ctx sdk.Context) []genesistypes.RegisteredInterchainAccount {
	_ = "STUB: not implemented"
	return nil
}

// SetInterchainAccountAddress stores the InterchainAccount address, keyed by the associated connectionID and portID
func (k *Keeper) SetInterchainAccountAddress(ctx sdk.Context, connectionID, portID, address string) {
	_ = "STUB: not implemented"
	return
}

// GetAuthority returns the 27-interchain-accounts host submodule's authority.
func (k *Keeper) GetAuthority() string {
	_ = "STUB: not implemented"

	// GetParams returns the total set of the host submodule parameters.
	return ""
}

func (k *Keeper) GetParams(ctx sdk.Context) types.Params {
	_ = "STUB: not implemented"
	return *new(types.Params)
}

// only panic on unset params and not on empty params

// SetParams sets the total set of the host submodule parameters.
func (k *Keeper) SetParams(ctx sdk.Context, params types.Params) { _ = "STUB: not implemented"; return }

// newModuleQuerySafeAllowList returns a list of all query paths labeled with module_query_safe in the proto files.
func newModuleQuerySafeAllowList() []string { _ = "STUB: not implemented"; return nil }

// Get the service descriptor

// Skip services that are annotated with the "cosmos.msg.v1.service" option.

// Get the method descriptor

// Skip methods that are not annotated with the "cosmos.query.v1.module_query_safe" option.

// Add the method to the whitelist
