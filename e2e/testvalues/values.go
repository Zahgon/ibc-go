package testvalues

import (
	"time"

	"github.com/cosmos/interchaintest/v11/ibc"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/e2e/semverutil"
)

const (
	StartingTokenAmount             int64  = 500_000_000_000
	IBCTransferAmount               int64  = 10_000
	InvalidAddress                  string = "<invalid-address>"
	DefaultGovV1ProposalTokenAmount        = 500_000_000
)

// VotingPeriod may differ per test.
var VotingPeriod = time.Second * 30

// ImmediatelyTimeout returns an ibc.IBCTimeout which will cause an IBC transfer to timeout immediately.
func ImmediatelyTimeout() *ibc.IBCTimeout { _ = "STUB: not implemented"; return nil }

func DefaultTransferAmount(denom string) sdk.Coin { _ = "STUB: not implemented"; return *new(sdk.Coin) }

func TransferAmount(amount int64, denom string) sdk.Coin {
	_ = "STUB: not implemented"
	return *new(sdk.Coin)
}

func TendermintClientID(id int) string { _ = "STUB: not implemented"; return "" }

func SolomachineClientID(id int) string { _ = "STUB: not implemented"; return "" }

var ReflectionServiceFeatureReleases = semverutil.FeatureReleases{
	MajorVersion: "v7",
}

// TokenMetadataFeatureReleases represents the releases the token metadata was released in.
var TokenMetadataFeatureReleases = semverutil.FeatureReleases{
	MajorVersion: "v8",
}

// GovGenesisFeatureReleases represents the releases the governance module genesis
// was upgraded from v1beta1 to v1.
var GovGenesisFeatureReleases = semverutil.FeatureReleases{
	MajorVersion: "v7",
}

// SelfParamsFeatureReleases represents the releases the transfer module started managing its own params.
var SelfParamsFeatureReleases = semverutil.FeatureReleases{
	MajorVersion: "v8",
}

// TotalEscrowFeatureReleases represents the releases the total escrow state entry was released in.
var TotalEscrowFeatureReleases = semverutil.FeatureReleases{
	MajorVersion: "v8",
	MinorVersions: []string{
		"v7.1",
	},
}

// IbcErrorsFeatureReleases represents the releases the IBC module level errors was released in.
var IbcErrorsFeatureReleases = semverutil.FeatureReleases{
	MajorVersion: "v8",
}

// LocalhostClientFeatureReleases represents the releases the localhost client was released in.
var LocalhostClientFeatureReleases = semverutil.FeatureReleases{
	MajorVersion: "v8",
	MinorVersions: []string{
		"v7.1",
	},
}

// AllowAllClientsWildcardFeatureReleases represents the releases the allow all clients wildcard was released in.
var AllowAllClientsWildcardFeatureReleases = semverutil.FeatureReleases{
	MajorVersion: "v9",
	MinorVersions: []string{
		"v8.1",
	},
}

// ChannelParamsFeatureReleases represents the releases the params for 04-channel was released in.
var ChannelParamsFeatureReleases = semverutil.FeatureReleases{
	MajorVersion: "v9",
	MinorVersions: []string{
		"v8.1",
	},
}

// GovV1MessagesFeatureReleases represents the releases the support for x/gov v1 messages was released in.
var GovV1MessagesFeatureReleases = semverutil.FeatureReleases{
	MajorVersion: "v8",
}

// TransactionEventQueryFeatureReleases represents the releases the support for --query flag
// in "query txs" for searching transactions that match exact events (since Cosmos SDK v0.50) was released in.
var TransactionEventQueryFeatureReleases = semverutil.FeatureReleases{
	MajorVersion: "v8",
}

var ChannelsV2FeatureReleases = semverutil.FeatureReleases{
	MajorVersion: "v10",
}

var ClientV2FeatureReleases = semverutil.FeatureReleases{
	MajorVersion: "v10",
}

var LocalhostWithDashFeatureReleases = semverutil.FeatureReleases{
	MajorVersion: "v10",
}
