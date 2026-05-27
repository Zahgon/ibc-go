package cli

import (
	"github.com/spf13/cobra"
)

const (
	flagSequences = "sequences"
)

// getCmdQueryNextSequenceSend defines the command to query a next send sequence for a given client
func getCmdQueryNextSequenceSend() *cobra.Command { _ = "STUB: not implemented"; return nil }

func getCmdQueryPacketCommitment() *cobra.Command { _ = "STUB: not implemented"; return nil }

func getCmdQueryPacketCommitments() *cobra.Command { _ = "STUB: not implemented"; return nil }

func getCmdQueryPacketAcknowledgement() *cobra.Command { _ = "STUB: not implemented"; return nil }

func getCmdQueryPacketReceipt() *cobra.Command { _ = "STUB: not implemented"; return nil }

// getCmdQueryUnreceivedPackets defines the command to query all the unreceived
// packets on the receiving chain
func getCmdQueryUnreceivedPackets() *cobra.Command { _ = "STUB: not implemented"; return nil }

// getCmdQueryUnreceivedAcks defines the command to query all the unreceived acks on the original sending chain
func getCmdQueryUnreceivedAcks() *cobra.Command { _ = "STUB: not implemented"; return nil }
