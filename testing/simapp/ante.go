package simapp

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"

	"github.com/cosmos/ibc-go/v11/modules/core/keeper"
)

// HandlerOptions are the options required for constructing a default SDK AnteHandler.
type HandlerOptions struct {
	ante.HandlerOptions
	IBCKeeper *keeper.Keeper
}

// NewAnteHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers, and deducts fees from the first
// signer.
func NewAnteHandler(options HandlerOptions) (sdk.AnteHandler, error) {
	_ = "STUB: not implemented"
	return *new(sdk.AnteHandler), nil
}

// outermost AnteDecorator. SetUpContext must be called first

// SetPubKeyDecorator must be called before all signature verification decorators
