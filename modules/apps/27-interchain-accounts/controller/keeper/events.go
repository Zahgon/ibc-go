package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

// EmitAcknowledgementEvent emits an event signalling a successful or failed acknowledgement and including the error
// details if any.
func EmitAcknowledgementEvent(ctx sdk.Context, packet channeltypes.Packet, ack exported.Acknowledgement,
	err error,
) {
	_ = "STUB: not implemented"
	return
}
