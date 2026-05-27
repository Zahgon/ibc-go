package simulation

import (
	"github.com/cosmos/cosmos-sdk/types/kv"
)

// NewDecodeStore returns a decoder function closure that unmarshals the KVPair's
// Value to the corresponding DenomTrace type.
func NewDecodeStore() func(kvA, kvB kv.Pair) string { _ = "STUB: not implemented"; return nil }
