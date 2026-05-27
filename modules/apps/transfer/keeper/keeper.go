package keeper

import (
	"cosmossdk.io/core/address"
	corestore "cosmossdk.io/core/store"
	"cosmossdk.io/log/v2"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	cmtbytes "github.com/cometbft/cometbft/libs/bytes"

	"github.com/cosmos/ibc-go/v11/modules/apps/transfer/types"
	porttypes "github.com/cosmos/ibc-go/v11/modules/core/05-port/types"
)

// Keeper defines the IBC fungible transfer keeper
type Keeper struct {
	storeService corestore.KVStoreService
	cdc          codec.BinaryCodec
	addressCodec address.Codec

	ics4Wrapper   porttypes.ICS4Wrapper
	channelKeeper types.ChannelKeeper
	msgRouter     types.MessageRouter
	AuthKeeper    types.AccountKeeper
	BankKeeper    types.BankKeeper

	// the address capable of executing a MsgUpdateParams message. Typically, this
	// should be the x/gov module account.
	authority string
}

// NewKeeper creates a new IBC transfer Keeper instance
func NewKeeper(cdc codec.BinaryCodec, addressCodec address.Codec, storeService corestore.KVStoreService, channelKeeper types.ChannelKeeper, msgRouter types.MessageRouter, authKeeper types.AccountKeeper, bankKeeper types.BankKeeper, authority string) *Keeper {
	_ = "STUB: not implemented"
	// ensure ibc transfer module account is set
	return nil
}

// default ICS4Wrapper is the channel keeper

// WithICS4Wrapper sets the ICS4Wrapper. This function may be used after
// the keepers creation to set the middleware which is above this module
// in the IBC application stack.
func (k *Keeper) WithICS4Wrapper(wrapper porttypes.ICS4Wrapper) { _ = "STUB: not implemented"; return }

// GetICS4Wrapper returns the ICS4Wrapper.
func (k *Keeper) GetICS4Wrapper() porttypes.ICS4Wrapper {
	_ = "STUB: not implemented"
	return *

	// GetAuthority returns the transfer module's authority.
	new(porttypes.ICS4Wrapper)
}

func (k *Keeper) GetAuthority() string {
	_ = "STUB: not implemented"

	// GetAddressCodec returns the address codec used by the keeper.
	return ""
}

func (k *Keeper) GetAddressCodec() address.Codec {
	_ = "STUB: not implemented"
	return *

	// Logger returns a module-specific logger.
	new(address.Codec)
}

func (*Keeper) Logger(ctx sdk.Context) log.Logger {
	_ = "STUB: not implemented"
	return *new(log.Logger)
}

// GetPort returns the portID for the transfer module. Used in ExportGenesis
func (k *Keeper) GetPort(ctx sdk.Context) string { _ = "STUB: not implemented"; return "" }

// SetPort sets the portID for the transfer module. Used in InitGenesis
func (k *Keeper) SetPort(ctx sdk.Context, portID string) { _ = "STUB: not implemented"; return }

// GetParams returns the current transfer module parameters.
func (k *Keeper) GetParams(ctx sdk.Context) types.Params {
	_ = "STUB: not implemented"
	return *new(types.Params)
}

// only panic on unset params and not on empty params

// SetParams sets the transfer module parameters.
func (k *Keeper) SetParams(ctx sdk.Context, params types.Params) { _ = "STUB: not implemented"; return }

// GetDenom retrieves the denom from store given the hash of the denom.
func (k *Keeper) GetDenom(ctx sdk.Context, denomHash cmtbytes.HexBytes) (types.Denom, bool) {
	_ = "STUB: not implemented"
	return *new(types.Denom), false
}

// HasDenom checks if a the key with the given denomination hash exists on the store.
func (k *Keeper) HasDenom(ctx sdk.Context, denomHash cmtbytes.HexBytes) bool {
	_ = "STUB: not implemented"
	return false
}

// SetDenom sets a new {denom hash -> denom } pair to the store.
// This allows for reverse lookup of the denom given the hash.
func (k *Keeper) SetDenom(ctx sdk.Context, denom types.Denom) { _ = "STUB: not implemented"; return }

// GetAllDenoms returns all the denominations.
func (k *Keeper) GetAllDenoms(ctx sdk.Context) types.Denoms {
	_ = "STUB: not implemented"
	return *new(types.Denoms)
}

// IterateDenoms iterates over the denominations in the store and performs a callback function.
func (k *Keeper) IterateDenoms(ctx sdk.Context, cb func(denom types.Denom) bool) {
	_ = "STUB: not implemented"
	return
}

// SetDenomMetadata sets an IBC token's denomination metadata
func (k *Keeper) SetDenomMetadata(ctx sdk.Context, denom types.Denom) {
	_ = "STUB: not implemented"
	return
}

// Setting base as IBC hash denom since bank keepers's SetDenomMetadata uses
// Base as key path and the IBC hash is what gives this token uniqueness
// on the executing chain

// GetTotalEscrowForDenom gets the total amount of source chain tokens that
// are in escrow, keyed by the denomination.
//
// NOTE: if there is no value stored in state for the provided denom then a new Coin is returned for the denom with an initial value of zero.
// This accommodates callers to simply call `Add()` on the returned Coin as an empty Coin literal (e.g. sdk.Coin{}) will trigger a panic due to the absence of a denom.
func (k *Keeper) GetTotalEscrowForDenom(ctx sdk.Context, denom string) sdk.Coin {
	_ = "STUB: not implemented"
	return *new(sdk.Coin)
}

// SetTotalEscrowForDenom stores the total amount of source chain tokens that are in escrow.
// Amount is stored in state if and only if it is not equal to zero. The function will panic
// if the amount is negative.
func (k *Keeper) SetTotalEscrowForDenom(ctx sdk.Context, coin sdk.Coin) {
	_ = "STUB: not implemented"
	return
}

// delete the key since Cosmos SDK x/bank module will prune any non-zero balances

// GetAllTotalEscrowed returns the escrow information for all the denominations.
func (k *Keeper) GetAllTotalEscrowed(ctx sdk.Context) sdk.Coins {
	_ = "STUB: not implemented"
	return *new(sdk.Coins)
}

// IterateTokensInEscrow iterates over the denomination escrows in the store
// and performs a callback function. Denominations for which an invalid value
// (i.e. not integer) is stored, will be skipped.
func (k *Keeper) IterateTokensInEscrow(ctx sdk.Context, storeprefix []byte, cb func(denomEscrow sdk.Coin) bool) {
	_ = "STUB: not implemented"
	return
}

// denom is empty

// total escrow amount cannot be unmarshalled to integer

// IsBlockedAddr checks if the given address is allowed to send or receive tokens.
// The module account is always allowed to send and receive tokens.
func (k *Keeper) IsBlockedAddr(addr sdk.AccAddress) bool { _ = "STUB: not implemented"; return false }
