package cli

import (
	"time"

	"github.com/spf13/cobra"

	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
)

const (
	// The controller chain channel version
	flagVersion = "version"
	// The channel ordering
	flagOrdering               = "ordering"
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
)

// defaultRelativePacketTimeoutTimestamp is the default packet timeout timestamp (in nanoseconds)
// relative to the current block timestamp of the counterparty chain provided by the client
// state. The timeout is disabled when set to 0. The default is currently set to a 10 minute
// timeout.
var defaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())

func newRegisterInterchainAccountCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func newSendTxCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// attempt to unmarshal ica msg data argument

// check for file path if JSON input is not provided

// parseOrdering gets the channel ordering from the flags.
func parseOrdering(cmd *cobra.Command) (channeltypes.Order, error) {
	_ = "STUB: not implemented"
	return *new(channeltypes.Order), nil
}
