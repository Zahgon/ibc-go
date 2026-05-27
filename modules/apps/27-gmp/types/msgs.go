package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	MaximumSenderLength   = 2048  // maximum length of the receiver address in bytes (value chosen arbitrarily)
	MaximumReceiverLength = 2048  // maximum length of the receiver address in bytes (value chosen arbitrarily)
	MaximumMemoLength     = 32768 // maximum length of the memo in bytes (value chosen arbitrarily)
	MaximumSaltLength     = 32    // maximum length of the salt in bytes (value chosen arbitrarily)
	MaximumPayloadLength  = 32768 // maximum length of the payload in bytes (value chosen arbitrarily)
)

var (
	_ sdk.Msg              = (*MsgSendCall)(nil)
	_ sdk.HasValidateBasic = (*MsgSendCall)(nil)
)

// NewMsgSendCall creates a new MsgSendCall instance
func NewMsgSendCall(sourceClient, sender, receiver string, payload, salt []byte, timeoutTimestamp uint64, encoding, memo string) *MsgSendCall {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic performs a basic check of the MsgSendCall fields.
// NOTE: The recipient addresses format is not validated as the format defined by
// the chain is not known to IBC.
func (msg MsgSendCall) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// receiver is allowed to be empty

// validateIdentifiers checks if the IBC identifiers are valid
func (msg MsgSendCall) validateIdentifiers() error { _ = "STUB: not implemented"; return nil }

// validateEncoding checks if the encoding is valid
func validateEncoding(encoding string) error { _ = "STUB: not implemented"; return nil }
