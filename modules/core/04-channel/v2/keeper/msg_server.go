package keeper

import (
	"context"

	"github.com/cosmos/ibc-go/v11/modules/core/04-channel/v2/types"
)

var _ types.MsgServer = &Keeper{}

// SendPacket implements the PacketMsgServer SendPacket method.
func (k *Keeper) SendPacket(goCtx context.Context, msg *types.MsgSendPacket) (*types.MsgSendPacketResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RecvPacket implements the PacketMsgServer RecvPacket method.
func (k *Keeper) RecvPacket(goCtx context.Context, msg *types.MsgRecvPacket) (*types.MsgRecvPacketResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check if this client is allowed to update if v2 config are set

// Perform TAO verification
//
// If the packet was already received, perform a no-op
// Use a cached context to prevent accidental state changes

// build up the recv results for each application callback.

// construct acknowledgement with single app acknowledgement that is the sentinel error acknowledgement

// Modify events in cached context to reflect unsuccessful acknowledgement

// successful app acknowledgement cannot equal sentinel error acknowledgement

// append app acknowledgement to the overall acknowledgement

// Set packet acknowledgement to async if any of the acknowledgements are async.

// Return error if there is more than 1 payload
// TODO: Handle case where there are multiple payloads

// write application state changes for asynchronous and successful acknowledgements
// if any application returns a failure, then we discard all state changes
// to ensure an atomic execution of all payloads

// sanity check to ensure returned acknowledgement and calculated isSuccess boolean matches

// Set packet acknowledgement only if the acknowledgement is not async.
// NOTE: IBC applications modules may call the WriteAcknowledgement asynchronously if the
// acknowledgement is async.

// store the packet temporarily until the application returns an acknowledgement

// TODO: store the packet for async applications to access if required.

// Acknowledgement defines an rpc handler method for MsgAcknowledgement.
func (k *Keeper) Acknowledgement(goCtx context.Context, msg *types.MsgAcknowledgement) (*types.MsgAcknowledgementResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check if this client is allowed to update if v2 config are set

// if recv was successful, each payload should have its own acknowledgement so we send each individual acknowledgment to the application
// otherwise, the acknowledgement only contains the sentinel error acknowledgement which we send to the application. The application is responsible
// for knowing that this is an error acknowledgement and executing the appropriate logic.

// Timeout implements the PacketMsgServer Timeout method.
func (k *Keeper) Timeout(goCtx context.Context, timeout *types.MsgTimeout) (*types.MsgTimeoutResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check if this client is allowed to update if v2 config are set
