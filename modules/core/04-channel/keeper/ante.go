package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
)

// RecvPacketReCheckTx applies replay protection ensuring that when relay messages are
// re-executed in ReCheckTx, we can appropriately filter out redundant relay transactions.
func (k *Keeper) RecvPacketReCheckTx(ctx sdk.Context, packet types.Packet) error {
	_ = "STUB: not implemented"
	return nil
}
