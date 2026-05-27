package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg              = (*MsgStoreCode)(nil)
	_ sdk.Msg              = (*MsgMigrateContract)(nil)
	_ sdk.Msg              = (*MsgRemoveChecksum)(nil)
	_ sdk.HasValidateBasic = (*MsgStoreCode)(nil)
	_ sdk.HasValidateBasic = (*MsgMigrateContract)(nil)
	_ sdk.HasValidateBasic = (*MsgRemoveChecksum)(nil)
)

// NewMsgStoreCode creates a new MsgStoreCode instance
func NewMsgStoreCode(signer string, code []byte) *MsgStoreCode {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.HasValidateBasic
func (m MsgStoreCode) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewMsgRemoveChecksum creates a new MsgRemoveChecksum instance
func NewMsgRemoveChecksum(signer string, checksum []byte) *MsgRemoveChecksum {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.HasValidateBasic
func (m MsgRemoveChecksum) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// MsgMigrateContract creates a new MsgMigrateContract instance
func NewMsgMigrateContract(signer, clientID string, checksum, migrateMsg []byte) *MsgMigrateContract {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.HasValidateBasic
func (m MsgMigrateContract) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
