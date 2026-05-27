package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Before each hour epoch, check if any of the rate limits have expired,
// and reset them if they have
func (k *Keeper) BeginBlocker(ctx sdk.Context) { _ = "STUB: not implemented"; return }
