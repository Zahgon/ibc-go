package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg              = (*MsgUpdateParams)(nil)
	_ sdk.HasValidateBasic = (*MsgUpdateParams)(nil)

	_ sdk.Msg              = (*MsgModuleQuerySafe)(nil)
	_ sdk.HasValidateBasic = (*MsgModuleQuerySafe)(nil)
)

// NewMsgUpdateParams creates a new MsgUpdateParams instance
func NewMsgUpdateParams(signer string, params Params) *MsgUpdateParams {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.HasValidateBasic
func (msg MsgUpdateParams) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewMsgModuleQuerySafe creates a new MsgModuleQuerySafe instance
func NewMsgModuleQuerySafe(signer string, requests []QueryRequest) *MsgModuleQuerySafe {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.HasValidateBasic
func (msg MsgModuleQuerySafe) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
