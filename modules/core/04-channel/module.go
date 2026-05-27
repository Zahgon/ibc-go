package channel

import (
	"github.com/spf13/cobra"
)

// Name returns the IBC channel ICS name.
func Name() string { _ = "STUB: not implemented"; return "" }

// GetQueryCmd returns the root query command for IBC channels.
func GetQueryCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }
