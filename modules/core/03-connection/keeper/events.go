package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/core/03-connection/types"
)

// emitConnectionOpenInitEvent emits a connection open init event
func emitConnectionOpenInitEvent(ctx sdk.Context, connectionID string, clientID string, counterparty types.Counterparty) {
	_ = "STUB: not implemented"
	return
}

// emitConnectionOpenTryEvent emits a connection open try event
func emitConnectionOpenTryEvent(ctx sdk.Context, connectionID string, clientID string, counterparty types.Counterparty) {
	_ = "STUB: not implemented"
	return
}

// emitConnectionOpenAckEvent emits a connection open acknowledge event
func emitConnectionOpenAckEvent(ctx sdk.Context, connectionID string, connectionEnd types.ConnectionEnd) {
	_ = "STUB: not implemented"
	return
}

// emitConnectionOpenConfirmEvent emits a connection open confirm event
func emitConnectionOpenConfirmEvent(ctx sdk.Context, connectionID string, connectionEnd types.ConnectionEnd) {
	_ = "STUB: not implemented"
	return
}
