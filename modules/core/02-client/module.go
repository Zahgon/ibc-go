package client

import (
	"github.com/spf13/cobra"
)

// Name returns the IBC client name
func Name() string { _ = "STUB: not implemented"; return "" }

// GetQueryCmd returns no root query command for the IBC client
func GetQueryCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetTxCmd returns the root tx command for 02-client.
func GetTxCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }
