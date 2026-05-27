package simapp

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"

	"github.com/cosmos/ibc-go/v11/modules/core/keeper"
)

// HandlerOptions extend the SDK's AnteHandler options by requiring the IBC keeper.
type HandlerOptions struct {
	ante.HandlerOptions

	IBCKeeper *keeper.Keeper
}

// NewAnteHandler creates a new ante handler
func NewAnteHandler(options HandlerOptions) (sdk.AnteHandler, error) {
	_ = "STUB: not implemented"
	return *new(sdk.AnteHandler), nil
}

// outermost AnteDecorator. SetUpContext must be called first

// SetPubKeyDecorator must be called before all signature verification decorators
