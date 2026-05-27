package cli

import (
	"github.com/spf13/cobra"
)

// GetCmdQueryDenom defines the command to query a denomination from a given hash or ibc denom.
func GetCmdQueryDenom() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryDenoms defines the command to query all the denominations that this chain maintains.
func GetCmdQueryDenoms() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdParams returns the command handler for ibc-transfer parameter querying.
func GetCmdParams() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryEscrowAddress returns the command handler for ibc-transfer parameter querying.
func GetCmdQueryEscrowAddress() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryDenomHash defines the command to query a denomination hash from a given trace.
func GetCmdQueryDenomHash() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetCmdQueryTotalEscrowForDenom defines the command to query the total amount of tokens in escrow for a denom
func GetCmdQueryTotalEscrowForDenom() *cobra.Command { _ = "STUB: not implemented"; return nil }
