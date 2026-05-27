package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
)

const MaxTimeoutDelta time.Duration = 24 * time.Hour

var (
	_ sdk.Msg              = (*MsgSendPacket)(nil)
	_ sdk.HasValidateBasic = (*MsgSendPacket)(nil)

	_ sdk.Msg              = (*MsgRecvPacket)(nil)
	_ sdk.HasValidateBasic = (*MsgRecvPacket)(nil)

	_ sdk.Msg              = (*MsgTimeout)(nil)
	_ sdk.HasValidateBasic = (*MsgTimeout)(nil)

	_ sdk.Msg              = (*MsgAcknowledgement)(nil)
	_ sdk.HasValidateBasic = (*MsgAcknowledgement)(nil)
)

// NewMsgSendPacket creates a new MsgSendPacket instance.
func NewMsgSendPacket(sourceClient string, timeoutTimestamp uint64, signer string, payloads ...Payload) *MsgSendPacket {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic performs basic checks on a MsgSendPacket.
func (msg *MsgSendPacket) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewMsgRecvPacket creates a new MsgRecvPacket instance.
func NewMsgRecvPacket(packet Packet, proofCommitment []byte, proofHeight clienttypes.Height, signer string) *MsgRecvPacket {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic performs basic checks on a MsgRecvPacket.
func (msg *MsgRecvPacket) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewMsgAcknowledgement creates a new MsgAcknowledgement instance
func NewMsgAcknowledgement(packet Packet, acknowledgement Acknowledgement, proofAcked []byte, proofHeight clienttypes.Height, signer string) *MsgAcknowledgement {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic performs basic checks on a MsgAcknowledgement.
func (msg *MsgAcknowledgement) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewMsgTimeout creates a new MsgTimeout instance
func NewMsgTimeout(packet Packet, proofUnreceived []byte, proofHeight clienttypes.Height, signer string) *MsgTimeout {
	_ = "STUB: not implemented"
	return nil
}

// ValidateBasic performs basic checks on a MsgTimeout
func (msg *MsgTimeout) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
