package sanitize

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/e2e/semverutil"
)

var (
	// govv1ProposalTitleAndSummary represents the releases that support the new title and summary fields.
	govv1ProposalTitleAndSummary = semverutil.FeatureReleases{
		MajorVersion: "v7",
	}
	// icaUnorderedChannelFeatureReleases represents the releasees that support the new ordering field.
	icaUnorderedChannelFeatureReleases = semverutil.FeatureReleases{
		MajorVersion: "v9",
		MinorVersions: []string{
			"v7.5",
			"v8.1",
		},
	}
)

// Messages removes any fields that are not supported by the chain version.
// For example, any fields that have been added in later sdk releases.
func Messages(tag string, msgs ...sdk.Msg) []sdk.Msg { _ = "STUB: not implemented"; return nil }

// removeUnknownFields removes any fields that are not supported by the chain version.
// The input message is returned if no changes are made.
func removeUnknownFields(tag string, msg sdk.Msg) sdk.Msg {
	_ = "STUB: not implemented"
	return *new(sdk.Msg)
}

// sanitize messages contained in the x/gov proposal
