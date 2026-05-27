package v2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/apps/rate-limiting/keeper"
	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
	channeltypesv2 "github.com/cosmos/ibc-go/v11/modules/core/04-channel/v2/types"
	"github.com/cosmos/ibc-go/v11/modules/core/api"
)

var _ api.IBCModule = (*IBCMiddleware)(nil)

type IBCMiddleware struct {
	app    api.IBCModule
	keeper keeper.Keeper
}

func NewIBCMiddleware(k keeper.Keeper, app api.IBCModule) IBCMiddleware {
	_ = "STUB: not implemented"
	return *new(IBCMiddleware)
}

func (im IBCMiddleware) OnSendPacket(ctx sdk.Context, sourceClient string, destinationClient string, sequence uint64, payload channeltypesv2.Payload, signer sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}

func (im IBCMiddleware) OnRecvPacket(ctx sdk.Context, sourceClient string, destinationClient string, sequence uint64, payload channeltypesv2.Payload, relayer sdk.AccAddress) channeltypesv2.RecvPacketResult {
	_ = "STUB: not implemented"
	return *new(channeltypesv2.RecvPacketResult)
}

// Check if the packet would cause the rate limit to be exceeded,
// and if so, return an ack error

// If the packet was not rate-limited, pass it down to the Transfer OnRecvPacket callback

func (im IBCMiddleware) OnTimeoutPacket(ctx sdk.Context, sourceClient string, destinationClient string, sequence uint64, payload channeltypesv2.Payload, relayer sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}

func (im IBCMiddleware) OnAcknowledgementPacket(ctx sdk.Context, sourceClient string, destinationClient string, sequence uint64, acknowledgement []byte, payload channeltypesv2.Payload, relayer sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Something looks off about this, please review carefully
func v2ToV1Packet(payload channeltypesv2.Payload, sourceClient, destinationClient string, sequence uint64) (channeltypes.Packet, error) {
	_ = "STUB: not implemented"
	return *new(channeltypes.Packet), nil
}
