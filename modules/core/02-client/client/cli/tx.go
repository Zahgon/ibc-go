package cli

import (
	"github.com/spf13/cobra"
)

const FlagAuthority = "authority"

// newCreateClientCmd defines the command to create a new IBC light client.
func newCreateClientCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// attempt to unmarshal client state argument

// check for file path if JSON input is not provided

// attempt to unmarshal consensus state argument

// check for file path if JSON input is not provided

func newAddCounterpartyCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// newDeleteClientCreatorCmd defines the command to delete the client creator for a given client.
func newDeleteClientCreatorCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// newUpdateClientConfigCmd defines the command to update the client config (allowed relayers) for a given client.
func newUpdateClientConfigCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// NOTE: Make sure all fields are gathered from the user for the config objects, as it replaces the entire existing config.
// In other words, if we add a new field to the config object, we need to make sure it is gathered here

// newUpdateClientCmd defines the command to update an IBC client.
func newUpdateClientCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// check for file path if JSON input is not provided

// newUpgradeClientCmd defines the command to upgrade an IBC light client.
func newUpgradeClientCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// attempt to unmarshal client state argument

// check for file path if JSON input is not provided

// attempt to unmarshal consensus state argument

// check for file path if JSON input is not provided

// newSubmitRecoverClientProposalCmd defines the command to recover an IBC light client.
func newSubmitRecoverClientProposalCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// newScheduleIBCUpgradeProposalCmd defines the command for submitting an IBC software upgrade proposal.
func newScheduleIBCUpgradeProposalCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// attempt to unmarshal client state argument

// check for file path if JSON input is not provided
