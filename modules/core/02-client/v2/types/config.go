package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Maximum length of the allowed relayers list
const MaxAllowedRelayersLength = 20

// NewConfig instantiates a new allowed relayer list for a client with provided addresses
func NewConfig(allowedRelayers ...string) Config { _ = "STUB: not implemented"; return *new(Config) }

// DefaultConfig is empty and therefore permissionless
func DefaultConfig() Config {
	_ = "STUB: not implemented"

	// Validate ensures all provided addresses are valid sdk Addresses
	return *new(Config)
}

func (c Config) Validate() error { _ = "STUB: not implemented"; return nil }

// IsAllowedRelayer checks if the given address is registered on the allowlist.
func (c Config) IsAllowedRelayer(relayer sdk.AccAddress) bool {
	_ = "STUB: not implemented"
	return false
}

func validateRelayers(allowedRelayers []string) error { _ = "STUB: not implemented"; return nil }
