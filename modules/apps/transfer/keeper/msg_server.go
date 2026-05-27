package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
)

var _ types.MsgServer = (*Keeper)(nil)

// Transfer defines an rpc handler method for MsgTransfer.
func (k *Keeper) Transfer(goCtx context.Context, msg *types.MsgTransfer) (*types.MsgTransferResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Using types.UnboundedSpendLimit allows us to send the entire balance of a given denom.

// if the channel does not exist, or we are using channel aliasing then use IBC V2 protocol
// otherwise use IBC V1 protocol

// otherwise try to send an IBC V2 packet, if the sourceChannel is not a IBC V2 client
// then core IBC will return a CounterpartyNotFound error

// if a V1 channel exists for the source channel, then use IBC V1 protocol

// telemetry for transfer occurs here, in IBC V2 this is done in the onSendPacket callback

func (k *Keeper) transferV1Packet(ctx sdk.Context, sourceChannel string, token types.Token, timeoutHeight clienttypes.Height, timeoutTimestamp uint64, sender sdk.AccAddress, packetData types.FungibleTokenPacketData) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (k *Keeper) transferV2Packet(ctx sdk.Context, encoding, sourceChannel string, timeoutTimestamp uint64, packetData types.FungibleTokenPacketData) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// NOTE: The sdk msg handler creates a new EventManager, so events must be correctly propagated back to the current context

// Each individual sdk.Result has exactly one Msg response. We aggregate here.

// UpdateParams defines an rpc handler method for MsgUpdateParams. Updates the ibc-transfer module's parameters.
func (k *Keeper) UpdateParams(goCtx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
