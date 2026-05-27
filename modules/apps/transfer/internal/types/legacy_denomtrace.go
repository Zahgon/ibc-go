package types

import (
	cmtbytes "github.com/cometbft/cometbft/libs/bytes"
)

// Hash returns the hex bytes of the SHA256 hash of the DenomTrace fields using the following formula:
//
// hash = sha256(tracePath + "/" + baseDenom)
func (dt DenomTrace) Hash() cmtbytes.HexBytes {
	_ = "STUB: not implemented"
	return *new(cmtbytes.HexBytes)
}

// GetPrefix returns the receiving denomination prefix composed by the trace info and a separator.
func (dt DenomTrace) GetPrefix() string { _ = "STUB: not implemented"; return "" }

// IBCDenom a coin denomination for an ICS20 fungible token in the format
// 'ibc/{hash(tracePath + baseDenom)}'. If the trace is empty, it will return the base denomination.
func (dt DenomTrace) IBCDenom() string { _ = "STUB: not implemented"; return "" }

// GetFullDenomPath returns the full denomination according to the ICS20 specification:
// tracePath + "/" + baseDenom
// If there exists no trace then the base denomination is returned.
func (dt DenomTrace) GetFullDenomPath() string { _ = "STUB: not implemented"; return "" }
