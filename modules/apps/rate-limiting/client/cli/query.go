package cli

import (
	"github.com/spf13/cobra"
)

const (
	FlagDenom = "denom"
)

// GetCmdQueryRateLimit implements a command to query rate limits by channel-id or client-id and denom
func GetCmdQueryRateLimit() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Query all rate limits for the channel/client ID if denom is not specified.

// Query specific rate limit if denom is provided

// GetCmdQueryAllRateLimits return all available rate limits.
func GetCmdQueryAllRateLimits() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryRateLimitsByChainID return all rate limits that exist between this chain
// and the specified ChainId
func GetCmdQueryRateLimitsByChainID() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryAllBlacklistedDenoms returns the command to query all blacklisted denoms
func GetCmdQueryAllBlacklistedDenoms() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryAllWhitelistedAddresses returns the command to query all whitelisted address pairs
func GetCmdQueryAllWhitelistedAddresses() *cobra.Command { _ = "STUB: not implemented"; return nil }
