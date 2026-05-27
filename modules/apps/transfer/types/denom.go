package types

import (
	"sort"

	cmtbytes "github.com/cometbft/cometbft/libs/bytes"
)

// NewDenom creates a new Denom instance given the base denomination and a variable number of hops.
func NewDenom(base string, trace ...Hop) Denom { _ = "STUB: not implemented"; return *new(Denom) }

// Validate performs a basic validation of the Denom fields.
func (d Denom) Validate() error {
	_ = "STUB: not implemented"
	// NOTE: base denom validation cannot be performed as each chain may define
	// its own base denom validation
	return nil
}

// Hash returns the hex bytes of the SHA256 hash of the Denom fields using the following formula:
//
// hash = sha256(trace + "/" + baseDenom)
func (d Denom) Hash() cmtbytes.HexBytes { _ = "STUB: not implemented"; return *new(cmtbytes.HexBytes) }

// IBCDenom a coin denomination for an ICS20 fungible token in the format
// 'ibc/{hash(trace + baseDenom)}'. If the trace is empty, it will return the base denomination.
func (d Denom) IBCDenom() string { _ = "STUB: not implemented"; return "" }

// Path returns the full denomination according to the ICS20 specification:
// trace + "/" + baseDenom
// If there exists no trace then the base denomination is returned.
func (d Denom) Path() string { _ = "STUB: not implemented"; return "" }

// nolint:revive // no error returned by WriteString
//nolint:revive // no error returned by WriteByte

//nolint:revive

// IsNative returns true if the denomination is native, thus containing no trace history.
func (d Denom) IsNative() bool { _ = "STUB: not implemented"; return false }

// HasPrefix returns true if the first element of the trace of the denom
// matches the provided portId and channelId.
func (d Denom) HasPrefix(portID, channelID string) bool {
	_ = "STUB: not implemented"
	// if the denom is native, then it is not prefixed by any port/channel pair
	return false
}

// Denoms defines a wrapper type for a slice of Denom.
type Denoms []Denom

// Validate performs a basic validation of each denomination trace info.
func (d Denoms) Validate() error { _ = "STUB: not implemented"; return nil }

var _ sort.Interface = (*Denoms)(nil)

// Len implements sort.Interface for Denoms
func (d Denoms) Len() int {
	_ = "STUB: not implemented"

	// Less implements sort.Interface for Denoms
	return 0
}

func (d Denoms) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Swap implements sort.Interface for Denoms
func (d Denoms) Swap(i, j int) { _ = "STUB: not implemented"; return }

// Sort is a helper function to sort the set of denomination in-place
func (d Denoms) Sort() Denoms { _ = "STUB: not implemented"; return *new(Denoms) }

// ExtractDenomFromPath returns the denom from the full path.
func ExtractDenomFromPath(fullPath string) Denom { _ = "STUB: not implemented"; return *new(Denom) }

// The IBC specification does not guarantee the expected format of the
// destination port or destination channel identifier. A short term solution
// to determine base denomination is to expect the channel identifier to be the
// one ibc-go specifies. A longer term solution is to separate the path and base
// denomination in the ICS20 packet. If an intermediate hop prefixes the full denom
// with a channel identifier format different from our own, the base denomination
// will be incorrectly parsed, but the token will continue to be treated correctly
// as an IBC denomination. The hash used to store the token internally on our chain
// will be the same value as the base denomination being correctly parsed.
// nolint:revive // Early return possible, but not needed here

// validateIBCDenom validates that the given denomination is either:
//
//   - A valid base denomination (eg: 'uatom' or 'gamm/pool/1' as in https://github.com/cosmos/ibc-go/issues/894)
//   - A valid fungible token representation (i.e 'ibc/{hash}') per ADR 001 https://github.com/cosmos/ibc-go/blob/main/docs/architecture/adr-001-coin-source-tracing.md
func validateIBCDenom(denom string) error { _ = "STUB: not implemented"; return nil }

// Valid base denomination or other valid format

// ParseHexHash parses a hex hash in string format to bytes and validates its correctness.
func ParseHexHash(hexHash string) (cmtbytes.HexBytes, error) {
	_ = "STUB: not implemented"
	return *new(cmtbytes.HexBytes), nil
}
