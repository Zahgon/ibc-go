package cli

import (
	"time"

	"github.com/spf13/cobra"
)

const (
	flagPacketTimeoutHeight    = "packet-timeout-height"
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	flagAbsoluteTimeouts       = "absolute-timeouts"
	flagMemo                   = "memo"
)

// defaultRelativePacketTimeoutTimestamp is the default packet timeout timestamp (in nanoseconds)
// relative to the current block timestamp of the counterparty chain provided by the client
// state. For IBC v1 protocol, either timeout timestamp or timeout height must be set.
// If you are sending with IBC v2 protocol, timeout timestamp must be set.
// The default is currently set to a 10 minute timeout.
var defaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())

// NewTransferTxCmd returns the command to create a NewMsgTransfer transaction
func NewTransferTxCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// NOTE: relative timeouts using block height are not supported.
// if the timeouts are not absolute, CLI users rely solely on local clock time in order to calculate relative timestamps.

// use local clock time as reference time for calculating timeout timestamp.
