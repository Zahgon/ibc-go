package clientv2

import (
	"github.com/spf13/cobra"
)

// Name returns the IBC channel ICS name.
func Name() string { _ = "STUB: not implemented"; return "" }

// GetTxCmd returns the root tx command for IBC channels.
func GetTxCmd() *cobra.Command {
	_ = "STUB: not implemented"
	// TODO

	// GetQueryCmd returns the root query command for IBC channels.
	return nil
}

func GetQueryCmd() *cobra.Command {
	_ = "STUB: not implemented"
	// TODO
	return nil
}
