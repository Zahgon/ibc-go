package simulation

import (
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/cosmos/ibc-go/v11/modules/core/keeper"
)

// NewDecodeStore returns a decoder function closure that unmarshals the KVPair's
// Value to the corresponding ibc type.
func NewDecodeStore(k keeper.Keeper) func(kvA, kvB kv.Pair) string {
	_ = "STUB: not implemented"
	return nil
}
