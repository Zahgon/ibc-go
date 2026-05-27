package cli

import (
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the query commands for the IBC channel/v2.
func GetQueryCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// NewTxCmd returns the command to submit transactions defined for IBC channel/v2.
func NewTxCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// TODO: Add v2 packet commands: https://github.com/cosmos/ibc-go/issues/7853
