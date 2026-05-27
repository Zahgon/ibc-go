package keeper

import (
	"cosmossdk.io/core/address"
	corestore "cosmossdk.io/core/store"
	"cosmossdk.io/log/v2"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/apps/rate-limiting/types"
	porttypes "github.com/cosmos/ibc-go/v11/modules/core/05-port/types"
)

// Keeper maintains the link to storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeService corestore.KVStoreService
	cdc          codec.BinaryCodec
	addressCodec address.Codec

	ics4Wrapper   porttypes.ICS4Wrapper
	channelKeeper types.ChannelKeeper
	clientKeeper  types.ClientKeeper

	bankKeeper types.BankKeeper
	authority  string
}

// NewKeeper creates a new rate-limiting Keeper instance
func NewKeeper(cdc codec.BinaryCodec, addressCodec address.Codec, storeService corestore.KVStoreService, channelKeeper types.ChannelKeeper, clientKeeper types.ClientKeeper, bankKeeper types.BankKeeper, authority string) *Keeper {
	_ = "STUB: not implemented"
	return nil
}

// Defaults to using the channel keeper as the ICS4Wrapper
// This can be overridden later with WithICS4Wrapper (e.g. by the middleware stack wiring)

// SetICS4Wrapper sets the ICS4Wrapper.
// It is used after the middleware is created since the keeper needs the underlying module's SendPacket capability,
// creating a dependency cycle.
func (k *Keeper) SetICS4Wrapper(ics4Wrapper porttypes.ICS4Wrapper) {
	_ = "STUB: not implemented"
	return
}

// ICS4Wrapper returns the ICS4Wrapper to send packets downstream.
func (k *Keeper) ICS4Wrapper() porttypes.ICS4Wrapper {
	_ = "STUB: not implemented"
	return *

	// GetAuthority returns the module's authority.
	new(porttypes.ICS4Wrapper)
}

func (k *Keeper) GetAuthority() string {
	_ = "STUB: not implemented"

	// Logger returns a module-specific logger.
	return ""
}

func (*Keeper) Logger(ctx sdk.Context) log.Logger {
	_ = "STUB: not implemented"
	return *new(log.Logger)
}
