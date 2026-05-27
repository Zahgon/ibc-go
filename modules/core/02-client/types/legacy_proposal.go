package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

const (
	// ProposalTypeClientUpdate defines the type for a ClientUpdateProposal
	ProposalTypeClientUpdate = "ClientUpdate"
	// ProposalTypeUpgrade defines the type for an UpgradeProposal
	ProposalTypeUpgrade = "IBCUpgrade"
)

var (
	_ govtypes.Content                   = &ClientUpdateProposal{}
	_ govtypes.Content                   = &UpgradeProposal{}
	_ codectypes.UnpackInterfacesMessage = &UpgradeProposal{}
)

// func init() {
// 	govtypes.RegisterProposalType(ProposalTypeClientUpdate)
// 	govtypes.RegisterProposalType(ProposalTypeUpgrade)
// }

// NewClientUpdateProposal creates a new client update proposal.
//
// Deprecated: The legacy v1beta1 gov ClientUpdateProposal is deprecated
// and will be removed in a future release. Please use MsgRecoverClient instead.
func NewClientUpdateProposal(title, description, subjectClientID, substituteClientID string) govtypes.Content {
	_ = "STUB: not implemented"
	return *new(govtypes.Content)
}

// GetTitle returns the title of a client update proposal.
func (cup *ClientUpdateProposal) GetTitle() string {
	_ = "STUB: not implemented"

	// GetDescription returns the description of a client update proposal.
	return ""
}

func (cup *ClientUpdateProposal) GetDescription() string { _ = "STUB: not implemented"; return "" }

// ProposalRoute returns the routing key of a client update proposal.
func (*ClientUpdateProposal) ProposalRoute() string {
	_ = "STUB: not implemented"

	// ProposalType returns the type of a client update proposal.
	return ""
}

func (*ClientUpdateProposal) ProposalType() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic runs basic stateless validity checks
func (cup *ClientUpdateProposal) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewUpgradeProposal creates a new IBC breaking upgrade proposal.
//
// Deprecated: The legacy v1beta1 gov UpgradeProposal is deprecated
// and will be removed in a future release. Please use MsgIBCSoftwareUpgrade instead.
func NewUpgradeProposal(title, description string, plan upgradetypes.Plan, upgradedClientState exported.ClientState) (govtypes.Content, error) {
	_ = "STUB: not implemented"
	return *new(govtypes.Content), nil
}

// GetTitle returns the title of a upgrade proposal.
func (up *UpgradeProposal) GetTitle() string {
	_ = "STUB: not implemented"

	// GetDescription returns the description of a upgrade proposal.
	return ""
}

func (up *UpgradeProposal) GetDescription() string { _ = "STUB: not implemented"; return "" }

// ProposalRoute returns the routing key of a upgrade proposal.
func (*UpgradeProposal) ProposalRoute() string {
	_ = "STUB: not implemented"

	// ProposalType returns the upgrade proposal type.
	return ""
}

func (*UpgradeProposal) ProposalType() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic runs basic stateless validity checks
func (up *UpgradeProposal) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// String returns the string representation of the UpgradeProposal.
func (up UpgradeProposal) String() string { _ = "STUB: not implemented"; return "" }

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (up UpgradeProposal) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}
