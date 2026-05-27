package types

import (
	commitmenttypesv2 "github.com/cosmos/ibc-go/v11/modules/core/23-commitment/types/v2"
)

// BuildMerklePath takes the merkle path prefix and an ICS24 path
// and builds a new path by appending the ICS24 path to the last element of the merkle path prefix.
func BuildMerklePath(prefix [][]byte, path []byte) commitmenttypesv2.MerklePath {
	_ = "STUB: not implemented"
	return *new(commitmenttypesv2.MerklePath)
}

// copy prefix to avoid modifying the original slice

// append path to last element
