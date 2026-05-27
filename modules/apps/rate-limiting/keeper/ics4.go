package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

func (k *Keeper) SendPacket(ctx sdk.Context, sourcePort string, sourceChannel string, timeoutHeight clienttypes.Height, timeoutTimestamp uint64, data []byte) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (k *Keeper) WriteAcknowledgement(ctx sdk.Context, packet exported.PacketI, ack exported.Acknowledgement) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Keeper) GetAppVersion(ctx sdk.Context, portID, channelID string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}
