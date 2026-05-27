package v2

import (
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var _ exported.Path = (*MerklePath)(nil)

// NewMerklePath creates a new MerklePath instance
// The keys must be passed in from root-to-leaf order
func NewMerklePath(keyPath ...[]byte) MerklePath {
	_ = "STUB: not implemented"
	return *new(MerklePath)
}

// GetKey will return a byte representation of the key
func (mp MerklePath) GetKey(i uint64) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Empty returns true if the path is empty
func (mp MerklePath) Empty() bool { _ = "STUB: not implemented"; return false }

// ValidateAsPrefix validates the MerklePath to ensure it is a valid prefix
// Thus every element of the merkle path must be non-empty except for the last element
// which may be empty. In this case, the ICS24 path will be appended to the last element
// to form the full path.
// This is the MerklePath being stored in the CounterpartyInfo
// It is interpreted as a prefix to the full path in particular it is the keypath
// from the root to the provable ICS24 store.
// Since it is not the full path to a leaf, the last element may be empty.
// This can occur if the commitment structure is a nested merkle tree and the ICS24
// store is itself a merkle tree.
func (mp MerklePath) ValidateAsPrefix() error { _ = "STUB: not implemented"; return nil }

// ValidateAsPath validates the MerklePath as a fully constructed path.
// Here every element must be non-empty since the MerklePath is no longer
// acting as a prefix but is instead the full path intended for verification.
// This is the full path to a leaf in the commitment tree constructed by IBC handler
// and it will be passed to the client for verification. Thus, at this point
// every element must be non-empty.
func (mp MerklePath) ValidateAsPath() error { _ = "STUB: not implemented"; return nil }
