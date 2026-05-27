package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	clientv1keeper "github.com/cosmos/ibc-go/v11/modules/core/02-client/keeper"
	"github.com/cosmos/ibc-go/v11/modules/core/02-client/v2/types"
)

type Keeper struct {
	cdc            codec.BinaryCodec
	ClientV1Keeper *clientv1keeper.Keeper
}

// NewKeeper creates a new client v2 keeper
func NewKeeper(
	cdc codec.BinaryCodec,
	clientV1Keeper *clientv1keeper.Keeper,
) *Keeper {
	_ = "STUB: not implemented"
	return nil
}

// SetClientCounterparty sets counterpartyInfo for a given clientID
func (k *Keeper) SetClientCounterparty(ctx sdk.Context, clientID string, counterparty types.CounterpartyInfo) {
	_ = "STUB: not implemented"
	return
}

// GetClientCounterparty gets counterpartyInfo for a given clientID
func (k *Keeper) GetClientCounterparty(ctx sdk.Context, clientID string) (types.CounterpartyInfo, bool) {
	_ = "STUB: not implemented"
	return *new(types.CounterpartyInfo), false
}

// GetConfig returns the ibc-client v2 configuration for the given clientID.
func (k *Keeper) GetConfig(ctx sdk.Context, clientID string) types.Config {
	_ = "STUB: not implemented"
	return *new(types.Config)
}

// SetConfig sets ibc-client v2 configuration for the given clientID.
func (k *Keeper) SetConfig(ctx sdk.Context, clientID string, config types.Config) {
	_ = "STUB: not implemented"
	return
}
