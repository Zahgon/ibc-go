package types

import (
	"context"

	"github.com/cosmos/gogoproto/proto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
)

var _ authz.Authorization = (*TransferAuthorization)(nil)

const (
	allocationNotFound = -1
)

// NewTransferAuthorization creates a new TransferAuthorization object.
func NewTransferAuthorization(allocations ...Allocation) *TransferAuthorization {
	_ = "STUB: not implemented"
	return nil
}

// MsgTypeURL implements Authorization.MsgTypeURL.
func (*TransferAuthorization) MsgTypeURL() string { _ = "STUB: not implemented"; return "" }

// Accept implements Authorization.Accept.
func (a *TransferAuthorization) Accept(goCtx context.Context, msg proto.Message) (authz.AcceptResponse, error) {
	_ = "STUB: not implemented"
	return *new(authz.AcceptResponse), nil
}

// bool flag to see if we have updated any of the allocations

// update spend limit the token token in the MsgTransfer
// If the spend limit is set to the MaxUint256 sentinel value, do not subtract the amount from the spend limit.
// if there is no unlimited spend, then we need to subtract the amount from the spend limit to get the limit left

// modify the spend limit with the reduced amount.

// if the spend limit is zero of the associated allocation then we delete it.
// NOTE: SpendLimit is an array of coins, with each one representing the remaining spend limit for an
// individual denomination.

// ValidateBasic implements Authorization.ValidateBasic.
func (a *TransferAuthorization) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// isAllowedAddress returns a boolean indicating if the receiver address is valid for transfer.
// gasCostPerIteration gas is consumed for each iteration.
func isAllowedAddress(ctx sdk.Context, receiver string, allowedAddrs []string) bool {
	_ = "STUB: not implemented"
	return false
}

// validateMemo returns a nil error indicating if the memo is valid for transfer.
func validateMemo(ctx sdk.Context, memo string, allowedMemos []string) error {
	_ = "STUB: not implemented"
	// if the allow list is empty, then the memo must be an empty string
	return nil
}

// if allowedPacketDataList has only 1 element and it equals AllowAllPacketDataKeys
// then accept all the memo strings

// getAllocationIndex ranges through a set of allocations, and returns the index of the allocation if found. If not, returns -1.
func getAllocationIndex(msg MsgTransfer, allocations []Allocation) int {
	_ = "STUB: not implemented"
	return 0
}
