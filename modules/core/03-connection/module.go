package connection

import (
	"github.com/spf13/cobra"
)

// Name returns the IBC connection ICS name.
func Name() string { _ = "STUB: not implemented"; return "" }

// GetQueryCmd returns the root query command for the IBC connections.
func GetQueryCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }
