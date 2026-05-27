package keeper

import (
	"context"

	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/core/store"
	"cosmossdk.io/log/v2"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/apps/27-gmp/types"
)

// Keeper defines the IBC fungible transfer keeper
type Keeper struct {
	cdc codec.Codec

	msgRouter types.MessageRouter

	accountKeeper types.AccountKeeper

	// the address capable of executing a MsgUpdateParams message. Typically, this
	// should be the x/gov module account.
	authority string

	// state management
	Schema collections.Schema
	// Accounts is a map of  (ClientID, Sender, Salt) to ICS27Account
	Accounts collections.Map[collections.Triple[string, string, []byte], types.ICS27Account]
	// AccountsByAddress is a map of sdk.AccAddress to ICS27Account, for reverse lookups
	AccountsByAddress collections.Map[sdk.AccAddress, types.ICS27Account]
}

// NewKeeper creates a new Keeper instance
func NewKeeper(
	cdc codec.Codec, storeService storetypes.KVStoreService,
	accountKeeper types.AccountKeeper, msgRouter types.MessageRouter,
	authority string,
) *Keeper {
	_ = "STUB: not implemented"
	return nil
}

// GetAuthority returns the module's authority.
func (k *Keeper) GetAuthority() string {
	_ = "STUB: not implemented"

	// Logger returns a module-specific logger.
	return ""
}

func (*Keeper) Logger(goCtx context.Context) log.Logger {
	_ = "STUB: not implemented"
	return *new(log.Logger)
}
