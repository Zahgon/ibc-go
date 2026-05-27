package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v11/modules/core/keeper"
)

type RedundantRelayDecorator struct {
	k *keeper.Keeper
}

func NewRedundantRelayDecorator(k *keeper.Keeper) RedundantRelayDecorator {
	_ = "STUB: not implemented"
	return *new(RedundantRelayDecorator)
}

// AnteHandle returns an error if a multiMsg tx only contains packet messages (Recv, Ack, Timeout) and additional update messages
// and all packet messages are redundant. If the transaction is just a single UpdateClient message, or the multimsg transaction
// contains some other message type, then the antedecorator returns no error and continues processing to ensure these transactions
// are included. This will ensure that relayers do not waste fees on multiMsg transactions when another relayer has already submitted
// all packets, by rejecting the tx at the mempool layer.
func (rrd RedundantRelayDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	_ = "STUB: not implemented"
	// do not run redundancy check on DeliverTx or simulate
	return *new(sdk.Context), nil
}

// keep track of total packet messages and number of redundancies across `RecvPacket`, `AcknowledgePacket`, and `TimeoutPacket/OnClose`

// when we are in ReCheckTx mode, ctx.IsCheckTx() will also return true
// therefore we must start the if statement on ctx.IsReCheckTx() to correctly
// determine which mode we are in

// if the multiMsg tx has a msg that is not a packet msg or update msg, then we will not return error
// regardless of if all packet messages are redundant. This ensures that non-packet messages get processed
// even if they get batched with redundant packet messages.

// only return error if all packet messages are redundant

// recvPacketCheckTx runs a subset of ibc recv packet logic to be used specifically within the RedundantRelayDecorator AnteHandler.
// It only performs core IBC receiving logic and skips any application logic.
func (rrd RedundantRelayDecorator) recvPacketCheckTx(ctx sdk.Context, msg *channeltypes.MsgRecvPacket) (*channeltypes.MsgRecvPacketResponse, error) {
	_ = "STUB: not implemented"
	// If the packet was already received, perform a no-op
	// Use a cached context to prevent accidental state changes
	return nil, nil
}

// recvPacketReCheckTx runs a subset of ibc recv packet logic to be used specifically within the RedundantRelayDecorator AnteHandler.
// It only performs core IBC receiving logic and skips any application logic.
func (rrd RedundantRelayDecorator) recvPacketReCheckTx(ctx sdk.Context, msg *channeltypes.MsgRecvPacket) (*channeltypes.MsgRecvPacketResponse, error) {
	_ = "STUB: not implemented"
	// If the packet was already received, perform a no-op
	// Use a cached context to prevent accidental state changes
	return nil, nil
}

// updateClientCheckTx runs a subset of ibc client update logic to be used specifically within the RedundantRelayDecorator AnteHandler.
// The following function performs ibc client message verification for CheckTx only and state updates in both CheckTx and ReCheckTx.
// Note that misbehaviour checks are omitted.
func (rrd RedundantRelayDecorator) updateClientCheckTx(ctx sdk.Context, msg *clienttypes.MsgUpdateClient) error {
	_ = "STUB: not implemented"
	return nil
}
