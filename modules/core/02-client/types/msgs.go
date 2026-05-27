package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var (
	_ sdk.Msg = (*MsgCreateClient)(nil)
	_ sdk.Msg = (*MsgUpdateClient)(nil)
	_ sdk.Msg = (*MsgUpgradeClient)(nil)
	_ sdk.Msg = (*MsgUpdateParams)(nil)
	_ sdk.Msg = (*MsgIBCSoftwareUpgrade)(nil)
	_ sdk.Msg = (*MsgRecoverClient)(nil)
	_ sdk.Msg = (*MsgDeleteClientCreator)(nil)

	_ sdk.HasValidateBasic = (*MsgCreateClient)(nil)
	_ sdk.HasValidateBasic = (*MsgUpdateClient)(nil)
	_ sdk.HasValidateBasic = (*MsgUpgradeClient)(nil)
	_ sdk.HasValidateBasic = (*MsgUpdateParams)(nil)
	_ sdk.HasValidateBasic = (*MsgIBCSoftwareUpgrade)(nil)
	_ sdk.HasValidateBasic = (*MsgRecoverClient)(nil)
	_ sdk.HasValidateBasic = (*MsgDeleteClientCreator)(nil)

	_ codectypes.UnpackInterfacesMessage = (*MsgCreateClient)(nil)
	_ codectypes.UnpackInterfacesMessage = (*MsgUpdateClient)(nil)
	_ codectypes.UnpackInterfacesMessage = (*MsgUpgradeClient)(nil)
	_ codectypes.UnpackInterfacesMessage = (*MsgIBCSoftwareUpgrade)(nil)
)

const (
	// MaxClientStateSize is the maximum allowed size of the client state in bytes. (This is an arbitrarily chosen value)
	MaxClientStateSize = 32768
	// MaxConsensusStateSize is the maximum allowed size of the consensus state in bytes. (This is an arbitrarily chosen value)
	MaxConsensusStateSize = 32768
)

// NewMsgCreateClient creates a new MsgCreateClient instance
func NewMsgCreateClient(
	clientState exported.ClientState, consensusState exported.ConsensusState, signer string,
) (*MsgCreateClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgCreateClient) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// validate the total size of client state

// validate the total size of consensus state

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (msg MsgCreateClient) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// NewMsgUpdateClient creates a new MsgUpdateClient instance
func NewMsgUpdateClient(id string, clientMsg exported.ClientMessage, signer string) (*MsgUpdateClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgUpdateClient) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (msg MsgUpdateClient) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// NewMsgUpgradeClient creates a new MsgUpgradeClient instance
func NewMsgUpgradeClient(clientID string, clientState exported.ClientState, consState exported.ConsensusState,
	upgradeClientProof, upgradeConsensusStateProof []byte, signer string,
) (*MsgUpgradeClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgUpgradeClient) ValidateBasic() error {
	_ = "STUB: not implemented"
	// will not validate client state as committed client may not form a valid client state.
	// client implementations are responsible for ensuring final upgraded client is valid.
	return nil
}

// will not validate consensus state here since the trusted kernel may not form a valid consensus state.
// client implementations are responsible for ensuring client can submit new headers against this consensus state.

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (msg MsgUpgradeClient) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// NewMsgRecoverClient creates a new MsgRecoverClient instance
func NewMsgRecoverClient(signer, subjectClientID, substituteClientID string) *MsgRecoverClient {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic performs basic checks on a MsgRecoverClient.
func (msg *MsgRecoverClient) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewMsgIBCSoftwareUpgrade creates a new MsgIBCSoftwareUpgrade instance
func NewMsgIBCSoftwareUpgrade(signer string, plan upgradetypes.Plan, upgradedClientState exported.ClientState) (*MsgIBCSoftwareUpgrade, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ValidateBasic performs basic checks on a MsgIBCSoftwareUpgrade.
func (msg *MsgIBCSoftwareUpgrade) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// for the time being, we should implicitly be on tendermint when using ibc-go

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (msg *MsgIBCSoftwareUpgrade) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// NewMsgUpdateParams creates a new instance of MsgUpdateParams.
func NewMsgUpdateParams(signer string, params Params) *MsgUpdateParams {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic performs basic checks on a MsgUpdateParams.
func (msg *MsgUpdateParams) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewMsgDeleteClientCreator creates a new instance of MsgDeleteClientCreator.
func NewMsgDeleteClientCreator(clientID string, signer string) *MsgDeleteClientCreator {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic performs basic validation of the MsgDeleteClientCreator fields.
func (msg *MsgDeleteClientCreator) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
