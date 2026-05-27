package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/apps/27-gmp/types"
)

var _ types.MsgServer = (*Keeper)(nil)

// SendCall defines the handler for the MsgSendCall message.
func (k *Keeper) SendCall(goCtx context.Context, msg *types.MsgSendCall) (*types.MsgSendCallResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Keeper) sendPacket(ctx sdk.Context, encoding, sourceClient string, timeoutTimestamp uint64, packetData types.GMPPacketData) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// NOTE: The sdk msg handler creates a new EventManager, so events must be correctly propagated back to the current context

// Each individual sdk.Result has exactly one Msg response. We aggregate here.
