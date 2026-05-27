package keeper

import (
	corestore "cosmossdk.io/core/store"
	"cosmossdk.io/log/v2"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/apps/27-interchain-accounts/controller/types"
	genesistypes "github.com/cosmos/ibc-go/v11/modules/apps/27-interchain-accounts/genesis/types"
	icatypes "github.com/cosmos/ibc-go/v11/modules/apps/27-interchain-accounts/types"
	porttypes "github.com/cosmos/ibc-go/v11/modules/core/05-port/types"
)

// Keeper defines the IBC interchain accounts controller keeper
type Keeper struct {
	storeService  corestore.KVStoreService
	cdc           codec.Codec
	ics4Wrapper   porttypes.ICS4Wrapper
	channelKeeper icatypes.ChannelKeeper

	msgRouter icatypes.MessageRouter

	// the address capable of executing a MsgUpdateParams message. Typically, this
	// should be the x/gov module account.
	authority string
}

// NewKeeper creates a new interchain accounts controller Keeper instance
func NewKeeper(
	cdc codec.Codec, storeService corestore.KVStoreService,
	channelKeeper icatypes.ChannelKeeper,
	msgRouter icatypes.MessageRouter, authority string,
) *Keeper {
	_ = "STUB: not implemented"
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

// GetConnectionID returns the connection id for the given port and channelIDs.
func (k *Keeper) GetConnectionID(ctx sdk.Context, portID, channelID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetAllPorts returns all ports to which the interchain accounts controller module is bound. Used in ExportGenesis
func (k *Keeper) GetAllPorts(ctx sdk.Context) []string { _ = "STUB: not implemented"; return nil }

// setPort sets the provided portID in state
func (k *Keeper) setPort(ctx sdk.Context, portID string) { _ = "STUB: not implemented"; return }

// GetAppVersion calls the ICS4Wrapper GetAppVersion function.
func (k *Keeper) GetAppVersion(ctx sdk.Context, portID, channelID string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// GetActiveChannelID retrieves the active channelID from the store, keyed by the provided connectionID and portID
func (k *Keeper) GetActiveChannelID(ctx sdk.Context, connectionID, portID string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// GetOpenActiveChannel retrieves the active channelID from the store, keyed by the provided connectionID and portID & checks if the channel in question is in state OPEN
func (k *Keeper) GetOpenActiveChannel(ctx sdk.Context, connectionID, portID string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// IsActiveChannelClosed retrieves the active channel from the store and returns true if the channel state is CLOSED, otherwise false
func (k *Keeper) IsActiveChannelClosed(ctx sdk.Context, connectionID, portID string) bool {
	_ = "STUB: not implemented"
	return false
}

// GetAllActiveChannels returns a list of all active interchain accounts controller channels and their associated connection and port identifiers
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

// IsMiddlewareEnabled returns true if the underlying application callbacks are enabled for given port and connection identifier pair, otherwise false
func (k *Keeper) IsMiddlewareEnabled(ctx sdk.Context, portID, connectionID string) bool {
	_ = "STUB: not implemented"
	return false
}

// IsMiddlewareDisabled returns true if the underlying application callbacks are disabled for the given port and connection identifier pair, otherwise false
func (k *Keeper) IsMiddlewareDisabled(ctx sdk.Context, portID, connectionID string) bool {
	_ = "STUB: not implemented"
	return false
}

// SetMiddlewareEnabled stores a flag to indicate that the underlying application callbacks should be enabled for the given port and connection identifier pair
func (k *Keeper) SetMiddlewareEnabled(ctx sdk.Context, portID, connectionID string) {
	_ = "STUB: not implemented"
	return
}

// SetMiddlewareDisabled stores a flag to indicate that the underlying application callbacks should be disabled for the given port and connection identifier pair
func (k *Keeper) SetMiddlewareDisabled(ctx sdk.Context, portID, connectionID string) {
	_ = "STUB: not implemented"
	return
}

// DeleteMiddlewareEnabled deletes the middleware enabled flag stored in state
func (k *Keeper) DeleteMiddlewareEnabled(ctx sdk.Context, portID, connectionID string) {
	_ = "STUB: not implemented"
	return
}

// GetAuthority returns the ica/controller submodule's authority.
func (k *Keeper) GetAuthority() string {
	_ = "STUB: not implemented"

	// getAppMetadata retrieves the interchain accounts channel metadata from the store associated with the provided portID and channelID
	return ""
}

func (k *Keeper) getAppMetadata(ctx sdk.Context, portID, channelID string) (icatypes.Metadata, error) {
	_ = "STUB: not implemented"
	return *new(icatypes.Metadata), nil
}

// GetParams returns the current ica/controller submodule parameters.
func (k *Keeper) GetParams(ctx sdk.Context) types.Params {
	_ = "STUB: not implemented"
	return *new(types.Params)
}

// only panic on unset params and not on empty params

// SetParams sets the ica/controller submodule parameters.
func (k *Keeper) SetParams(ctx sdk.Context, params types.Params) { _ = "STUB: not implemented"; return }
