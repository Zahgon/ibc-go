package testsuite

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	abci "github.com/cometbft/cometbft/abci/types"
)

// ABCIToSDKEvents converts a list of ABCI events to Cosmos SDK events.
func ABCIToSDKEvents(abciEvents []abci.Event) sdk.Events {
	_ = "STUB: not implemented"
	return *new(sdk.Events)
}
