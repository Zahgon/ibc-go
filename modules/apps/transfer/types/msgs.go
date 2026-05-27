package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
)

const (
	MaximumReceiverLength = 2048  // maximum length of the receiver address in bytes (value chosen arbitrarily)
	MaximumMemoLength     = 32768 // maximum length of the memo in bytes (value chosen arbitrarily)
)

var (
	_ sdk.Msg              = (*MsgUpdateParams)(nil)
	_ sdk.Msg              = (*MsgTransfer)(nil)
	_ sdk.HasValidateBasic = (*MsgUpdateParams)(nil)
	_ sdk.HasValidateBasic = (*MsgTransfer)(nil)
)

// NewMsgUpdateParams creates a new MsgUpdateParams instance
func NewMsgUpdateParams(signer string, params Params) *MsgUpdateParams {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgUpdateParams) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewMsgTransfer creates a new MsgTransfer instance
func NewMsgTransfer(
	sourcePort, sourceChannel string,
	token sdk.Coin, sender, receiver string,
	timeoutHeight clienttypes.Height, timeoutTimestamp uint64,
	memo string,
) *MsgTransfer {
	_ = "STUB: not implemented"
	return nil
}

// NewMsgTransferAliased creates a new MsgTransfer instance
// with isV2 set to true, indicating that it is using the V2 protocol
// with v1 channel identifiers.
func NewMsgTransferAliased(
	sourcePort, sourceChannel string,
	token sdk.Coin, sender, receiver string,
	timeoutHeight clienttypes.Height, timeoutTimestamp uint64,
	memo string,
) *MsgTransfer {
	_ = "STUB: not implemented"
	return nil
}

// This indicates that the message is using the V2 protocol with aliased channel identifiers

// NewMsgTransferWithEncoding creates a new MsgTransfer instance
// with the provided encoding
func NewMsgTransferWithEncoding(
	sourcePort, sourceChannel string,
	token sdk.Coin, sender, receiver string,
	timeoutHeight clienttypes.Height, timeoutTimestamp uint64,
	memo string, encoding string, useAliasing bool,
) *MsgTransfer {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic performs a basic check of the MsgTransfer fields.
// NOTE: If you are sending with V1 protocol, timeoutHeight or timeoutTimestamp must be non-zero,
// if you are sending with V2 protocol, timeoutTimestamp must be non-zero and timeoutHeight must be zero
// NOTE: The recipient addresses format is not validated as the format defined by
// the chain is not known to IBC.
func (msg MsgTransfer) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// validateIdentifiers checks if the source port and channel identifiers are valid
func (msg MsgTransfer) validateIdentifiers() error { _ = "STUB: not implemented"; return nil }

// if we are using aliasing, then the source channel must be in the channel id format
// expected by ibc-go
// otherwise, it may be either a client id using v2 directly or a channel id using ibc v1
// thus, we perform a less strict check

// isValidIBCCoin returns true if the token provided is valid,
// and should be used to transfer tokens.
func isValidIBCCoin(coin sdk.Coin) bool { _ = "STUB: not implemented"; return false }

// validateIBCCoin returns true if the token provided is valid,
// and should be used to transfer tokens. The token must
// have a positive amount.
func validateIBCCoin(coin sdk.Coin) error { _ = "STUB: not implemented"; return nil }
