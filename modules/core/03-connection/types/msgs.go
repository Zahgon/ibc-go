package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
	commitmenttypes "github.com/cosmos/ibc-go/v11/modules/core/23-commitment/types"
)

var (
	_ sdk.Msg = (*MsgConnectionOpenInit)(nil)
	_ sdk.Msg = (*MsgConnectionOpenConfirm)(nil)
	_ sdk.Msg = (*MsgConnectionOpenAck)(nil)
	_ sdk.Msg = (*MsgConnectionOpenTry)(nil)
	_ sdk.Msg = (*MsgUpdateParams)(nil)

	_ sdk.HasValidateBasic = (*MsgConnectionOpenInit)(nil)
	_ sdk.HasValidateBasic = (*MsgConnectionOpenConfirm)(nil)
	_ sdk.HasValidateBasic = (*MsgConnectionOpenAck)(nil)
	_ sdk.HasValidateBasic = (*MsgConnectionOpenTry)(nil)
	_ sdk.HasValidateBasic = (*MsgUpdateParams)(nil)
)

// NewMsgConnectionOpenInit creates a new MsgConnectionOpenInit instance. It sets the
// counterparty connection identifier to be empty.
func NewMsgConnectionOpenInit(
	clientID, counterpartyClientID string,
	counterpartyPrefix commitmenttypes.MerklePrefix,
	version *Version, delayPeriod uint64, signer string,
) *MsgConnectionOpenInit {
	_ = "STUB: not implemented"
	// counterparty must have the same delay period
	return nil
}

// ValidateBasic implements sdk.Msg.
func (msg MsgConnectionOpenInit) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NOTE: Version can be nil on MsgConnectionOpenInit

// NewMsgConnectionOpenTry creates a new MsgConnectionOpenTry instance
func NewMsgConnectionOpenTry(
	clientID, counterpartyConnectionID, counterpartyClientID string,
	counterpartyPrefix commitmenttypes.MerklePrefix,
	counterpartyVersions []*Version, delayPeriod uint64,
	initProof []byte, proofHeight clienttypes.Height, signer string,
) *MsgConnectionOpenTry {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgConnectionOpenTry) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// counterparty validate basic allows empty counterparty connection identifiers

// NewMsgConnectionOpenAck creates a new MsgConnectionOpenAck instance
func NewMsgConnectionOpenAck(
	connectionID, counterpartyConnectionID string, tryProof []byte,
	proofHeight clienttypes.Height, version *Version, signer string,
) *MsgConnectionOpenAck {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgConnectionOpenAck) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewMsgConnectionOpenConfirm creates a new MsgConnectionOpenConfirm instance
func NewMsgConnectionOpenConfirm(
	connectionID string, ackProof []byte, proofHeight clienttypes.Height,
	signer string,
) *MsgConnectionOpenConfirm {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgConnectionOpenConfirm) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewMsgUpdateParams creates a new MsgUpdateParams instance
func NewMsgUpdateParams(signer string, params Params) *MsgUpdateParams {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic performs basic checks on a MsgUpdateParams.
func (msg *MsgUpdateParams) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
