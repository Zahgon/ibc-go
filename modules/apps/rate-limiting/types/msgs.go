package types

import (
	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = &MsgAddRateLimit{}
	_ sdk.Msg = &MsgUpdateRateLimit{}
	_ sdk.Msg = &MsgRemoveRateLimit{}
	_ sdk.Msg = &MsgResetRateLimit{}
)

// ----------------------------------------------
//               MsgAddRateLimit
// ----------------------------------------------

func NewMsgAddRateLimit(denom, channelOrClientID string, maxPercentSend sdkmath.Int, maxPercentRecv sdkmath.Int, durationHours uint64) *MsgAddRateLimit {
	_ = "STUB: not implemented"
	return nil
}

func (msg *MsgAddRateLimit) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// ----------------------------------------------
//               MsgUpdateRateLimit
// ----------------------------------------------

func NewMsgUpdateRateLimit(denom, channelOrClientID string, maxPercentSend sdkmath.Int, maxPercentRecv sdkmath.Int, durationHours uint64) *MsgUpdateRateLimit {
	_ = "STUB: not implemented"
	return nil
}

func (msg *MsgUpdateRateLimit) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// ----------------------------------------------
//               MsgRemoveRateLimit
// ----------------------------------------------

func NewMsgRemoveRateLimit(denom, channelOrClientID string) *MsgRemoveRateLimit {
	_ = "STUB: not implemented"
	return nil
}

func (msg *MsgRemoveRateLimit) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// ----------------------------------------------
//               MsgResetRateLimit
// ----------------------------------------------

func NewMsgResetRateLimit(denom, channelOrClientID string) *MsgResetRateLimit {
	_ = "STUB: not implemented"
	return nil
}

func (msg *MsgResetRateLimit) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
