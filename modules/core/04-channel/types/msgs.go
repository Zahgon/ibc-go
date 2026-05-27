package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
)

var (
	_ sdk.Msg = (*MsgChannelOpenInit)(nil)
	_ sdk.Msg = (*MsgChannelOpenTry)(nil)
	_ sdk.Msg = (*MsgChannelOpenAck)(nil)
	_ sdk.Msg = (*MsgChannelOpenConfirm)(nil)
	_ sdk.Msg = (*MsgChannelCloseInit)(nil)
	_ sdk.Msg = (*MsgChannelCloseConfirm)(nil)
	_ sdk.Msg = (*MsgRecvPacket)(nil)
	_ sdk.Msg = (*MsgAcknowledgement)(nil)
	_ sdk.Msg = (*MsgTimeout)(nil)
	_ sdk.Msg = (*MsgTimeoutOnClose)(nil)

	_ sdk.HasValidateBasic = (*MsgChannelOpenInit)(nil)
	_ sdk.HasValidateBasic = (*MsgChannelOpenTry)(nil)
	_ sdk.HasValidateBasic = (*MsgChannelOpenAck)(nil)
	_ sdk.HasValidateBasic = (*MsgChannelOpenConfirm)(nil)
	_ sdk.HasValidateBasic = (*MsgChannelCloseInit)(nil)
	_ sdk.HasValidateBasic = (*MsgChannelCloseConfirm)(nil)
	_ sdk.HasValidateBasic = (*MsgRecvPacket)(nil)
	_ sdk.HasValidateBasic = (*MsgAcknowledgement)(nil)
	_ sdk.HasValidateBasic = (*MsgTimeout)(nil)
	_ sdk.HasValidateBasic = (*MsgTimeoutOnClose)(nil)
)

// NewMsgChannelOpenInit creates a new MsgChannelOpenInit. It sets the counterparty channel
// identifier to be empty.
func NewMsgChannelOpenInit(
	portID, version string, channelOrder Order, connectionHops []string,
	counterpartyPortID string, signer string,
) *MsgChannelOpenInit {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgChannelOpenInit) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewMsgChannelOpenTry creates a new MsgChannelOpenTry instance
// The version string is deprecated and will be ignored by core IBC.
// It is left as an argument for go API backwards compatibility.
func NewMsgChannelOpenTry(
	portID, version string, channelOrder Order, connectionHops []string,
	counterpartyPortID, counterpartyChannelID, counterpartyVersion string,
	initProof []byte, proofHeight clienttypes.Height, signer string,
) *MsgChannelOpenTry {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgChannelOpenTry) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// counterparty validate basic allows empty counterparty channel identifiers

// NewMsgChannelOpenAck creates a new MsgChannelOpenAck instance
func NewMsgChannelOpenAck(
	portID, channelID, counterpartyChannelID string, cpv string, tryProof []byte, proofHeight clienttypes.Height,
	signer string,
) *MsgChannelOpenAck {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgChannelOpenAck) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewMsgChannelOpenConfirm creates a new MsgChannelOpenConfirm instance
func NewMsgChannelOpenConfirm(
	portID, channelID string, ackProof []byte, proofHeight clienttypes.Height,
	signer string,
) *MsgChannelOpenConfirm {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgChannelOpenConfirm) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewMsgChannelCloseInit creates a new MsgChannelCloseInit instance
func NewMsgChannelCloseInit(
	portID string, channelID string, signer string,
) *MsgChannelCloseInit {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgChannelCloseInit) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewMsgChannelCloseConfirm creates a new MsgChannelCloseConfirm instance
func NewMsgChannelCloseConfirm(
	portID, channelID string, initProof []byte, proofHeight clienttypes.Height,
	signer string,
) *MsgChannelCloseConfirm {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgChannelCloseConfirm) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewMsgRecvPacket constructs new MsgRecvPacket
func NewMsgRecvPacket(
	packet Packet, commitmentProof []byte, proofHeight clienttypes.Height,
	signer string,
) *MsgRecvPacket {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgRecvPacket) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// GetDataSignBytes returns the base64-encoded bytes used for the
// data field when signing the packet.
func (msg MsgRecvPacket) GetDataSignBytes() []byte { _ = "STUB: not implemented"; return nil }

// NewMsgTimeout constructs new MsgTimeout
func NewMsgTimeout(
	packet Packet, nextSequenceRecv uint64, unreceivedProof []byte,
	proofHeight clienttypes.Height, signer string,
) *MsgTimeout {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgTimeout) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewMsgTimeoutOnClose constructs new MsgTimeoutOnClose
func NewMsgTimeoutOnClose(
	packet Packet, nextSequenceRecv uint64,
	unreceivedProof, closeProof []byte,
	proofHeight clienttypes.Height, signer string,
) *MsgTimeoutOnClose {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgTimeoutOnClose) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewMsgAcknowledgement constructs a new MsgAcknowledgement
func NewMsgAcknowledgement(
	packet Packet,
	ack, ackedProof []byte,
	proofHeight clienttypes.Height,
	signer string,
) *MsgAcknowledgement {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic implements sdk.Msg
func (msg MsgAcknowledgement) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
