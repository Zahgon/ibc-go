package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	icatypes "github.com/cosmos/ibc-go/v11/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
)

const MaximumOwnerLength = 2048 // maximum length of the owner in bytes (value chosen arbitrarily)

var (
	_ sdk.Msg = (*MsgRegisterInterchainAccount)(nil)
	_ sdk.Msg = (*MsgSendTx)(nil)
	_ sdk.Msg = (*MsgUpdateParams)(nil)

	_ sdk.HasValidateBasic = (*MsgRegisterInterchainAccount)(nil)
	_ sdk.HasValidateBasic = (*MsgSendTx)(nil)
	_ sdk.HasValidateBasic = (*MsgUpdateParams)(nil)
)

// NewMsgRegisterInterchainAccount creates a new instance of MsgRegisterInterchainAccount
func NewMsgRegisterInterchainAccount(connectionID, owner, version string, ordering channeltypes.Order) *MsgRegisterInterchainAccount {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgRegisterInterchainAccount) ValidateBasic() error {
	_ = "STUB: not implemented"
	return nil
}

// NewMsgSendTx creates a new instance of MsgSendTx
func NewMsgSendTx(owner, connectionID string, relativeTimeoutTimestamp uint64, packetData icatypes.InterchainAccountPacketData) *MsgSendTx {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgSendTx) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewMsgUpdateParams creates a new MsgUpdateParams instance
func NewMsgUpdateParams(signer string, params Params) *MsgUpdateParams {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgUpdateParams) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
