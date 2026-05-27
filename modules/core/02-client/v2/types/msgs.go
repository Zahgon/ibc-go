package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MaxCounterpartyMerklePrefixElements defines the maximum number of elements allowed in the counterparty merkle prefix. (This is an arbitrarily chosen value)
const MaxCounterpartyMerklePrefixElements = 32

var (
	_ sdk.Msg = (*MsgRegisterCounterparty)(nil)
	_ sdk.Msg = (*MsgUpdateClientConfig)(nil)

	_ sdk.HasValidateBasic = (*MsgRegisterCounterparty)(nil)
	_ sdk.HasValidateBasic = (*MsgUpdateClientConfig)(nil)
)

// NewMsgRegisterCounterparty creates a new instance of MsgRegisterCounterparty.
func NewMsgRegisterCounterparty(clientID string, merklePrefix [][]byte, counterpartyClientID string, signer string) *MsgRegisterCounterparty {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic performs basic checks on a MsgRegisterCounterparty.
func (msg *MsgRegisterCounterparty) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// This check must be done because the transfer v2 module assumes that the client IDs in the packet
// are in the format {clientID}-{sequence}

func NewMsgUpdateClientConfig(clientID string, signer string, config Config) *MsgUpdateClientConfig {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic performs basic validation of the MsgUpdateClientConfig fields.
func (msg *MsgUpdateClientConfig) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
