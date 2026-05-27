package keeper

import (
	"context"

	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
	clientv2types "github.com/cosmos/ibc-go/v11/modules/core/02-client/v2/types"
	connectiontypes "github.com/cosmos/ibc-go/v11/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
)

var (
	_ clienttypes.MsgServer     = (*Keeper)(nil)
	_ clientv2types.MsgServer   = (*Keeper)(nil)
	_ connectiontypes.MsgServer = (*Keeper)(nil)
	_ channeltypes.MsgServer    = (*Keeper)(nil)
)

// CreateClient defines a rpc handler method for MsgCreateClient.
// NOTE: The raw bytes of the concrete types encoded into protobuf.Any is passed to the client keeper.
// The 02-client handler will route to the appropriate light client module based on client type and it is the responsibility
// of the light client module to unmarshal and interpret the proto encoded bytes.
// Backwards compatibility with older versions of ibc-go is maintained through the light client module reconstructing and encoding
// the expected concrete type to the protobuf.Any for proof verification.
func (k *Keeper) CreateClient(goCtx context.Context, msg *clienttypes.MsgCreateClient) (*clienttypes.MsgCreateClientResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// set the client creator so that IBC v2 counterparty can be set by same relayer

// RegisterCounterparty will register the IBC v2 counterparty info for the given client id
// it must be called by the same relayer that called CreateClient
func (k *Keeper) RegisterCounterparty(goCtx context.Context, msg *clientv2types.MsgRegisterCounterparty) (*clientv2types.MsgRegisterCounterpartyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// initialize next sequence send to enable packet flow

// UpdateClient defines a rpc handler method for MsgUpdateClient.
func (k *Keeper) UpdateClient(goCtx context.Context, msg *clienttypes.MsgUpdateClient) (*clienttypes.MsgUpdateClientResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// only check v2 params if this chain is setup with v2 clientKeepr

// check if this relayer is allowed to update if v2 configuration are set

// UpgradeClient defines a rpc handler method for MsgUpgradeClient.
// NOTE: The raw bytes of the concrete types encoded into protobuf.Any is passed to the client keeper.
// The 02-client handler will route to the appropriate light client module based on client identifier and it is the responsibility
// of the light client module to unmarshal and interpret the proto encoded bytes.
// Backwards compatibility with older versions of ibc-go is maintained through the light client module reconstructing and encoding
// the expected concrete type to the protobuf.Any for proof verification.
func (k *Keeper) UpgradeClient(goCtx context.Context, msg *clienttypes.MsgUpgradeClient) (*clienttypes.MsgUpgradeClientResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RecoverClient defines a rpc handler method for MsgRecoverClient.
func (k *Keeper) RecoverClient(goCtx context.Context, msg *clienttypes.MsgRecoverClient) (*clienttypes.MsgRecoverClientResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IBCSoftwareUpgrade defines a rpc handler method for MsgIBCSoftwareUpgrade.
func (k *Keeper) IBCSoftwareUpgrade(goCtx context.Context, msg *clienttypes.MsgIBCSoftwareUpgrade) (*clienttypes.MsgIBCSoftwareUpgradeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConnectionOpenInit defines a rpc handler method for MsgConnectionOpenInit.
func (k *Keeper) ConnectionOpenInit(goCtx context.Context, msg *connectiontypes.MsgConnectionOpenInit) (*connectiontypes.MsgConnectionOpenInitResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConnectionOpenTry defines a rpc handler method for MsgConnectionOpenTry.
func (k *Keeper) ConnectionOpenTry(goCtx context.Context, msg *connectiontypes.MsgConnectionOpenTry) (*connectiontypes.MsgConnectionOpenTryResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConnectionOpenAck defines a rpc handler method for MsgConnectionOpenAck.
func (k *Keeper) ConnectionOpenAck(goCtx context.Context, msg *connectiontypes.MsgConnectionOpenAck) (*connectiontypes.MsgConnectionOpenAckResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConnectionOpenConfirm defines a rpc handler method for MsgConnectionOpenConfirm.
func (k *Keeper) ConnectionOpenConfirm(goCtx context.Context, msg *connectiontypes.MsgConnectionOpenConfirm) (*connectiontypes.MsgConnectionOpenConfirmResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ChannelOpenInit defines a rpc handler method for MsgChannelOpenInit.
// ChannelOpenInit will perform 04-channel checks, route to the application
// callback, and write an OpenInit channel into state upon successful execution.
func (k *Keeper) ChannelOpenInit(goCtx context.Context, msg *channeltypes.MsgChannelOpenInit) (*channeltypes.MsgChannelOpenInitResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve application callbacks from router

// Perform 04-channel verification

// Perform application logic callback

// Write channel into state

// ChannelOpenTry defines a rpc handler method for MsgChannelOpenTry.
// ChannelOpenTry will perform 04-channel checks, route to the application
// callback, and write an OpenTry channel into state upon successful execution.
func (k *Keeper) ChannelOpenTry(goCtx context.Context, msg *channeltypes.MsgChannelOpenTry) (*channeltypes.MsgChannelOpenTryResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve application callbacks from router

// Perform 04-channel verification

// Perform application logic callback

// Write channel into state

// ChannelOpenAck defines a rpc handler method for MsgChannelOpenAck.
// ChannelOpenAck will perform 04-channel checks, route to the application
// callback, and write an OpenAck channel into state upon successful execution.
func (k *Keeper) ChannelOpenAck(goCtx context.Context, msg *channeltypes.MsgChannelOpenAck) (*channeltypes.MsgChannelOpenAckResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve application callbacks from router

// Perform 04-channel verification

// Write channel into state

// Perform application logic callback

// ChannelOpenConfirm defines a rpc handler method for MsgChannelOpenConfirm.
// ChannelOpenConfirm will perform 04-channel checks, route to the application
// callback, and write an OpenConfirm channel into state upon successful execution.
func (k *Keeper) ChannelOpenConfirm(goCtx context.Context, msg *channeltypes.MsgChannelOpenConfirm) (*channeltypes.MsgChannelOpenConfirmResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve application callbacks from router

// Perform 04-channel verification

// Write channel into state

// Perform application logic callback

// ChannelCloseInit defines a rpc handler method for MsgChannelCloseInit.
func (k *Keeper) ChannelCloseInit(goCtx context.Context, msg *channeltypes.MsgChannelCloseInit) (*channeltypes.MsgChannelCloseInitResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve callbacks from router

// ChannelCloseConfirm defines a rpc handler method for MsgChannelCloseConfirm.
func (k *Keeper) ChannelCloseConfirm(goCtx context.Context, msg *channeltypes.MsgChannelCloseConfirm) (*channeltypes.MsgChannelCloseConfirmResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve callbacks from router

// RecvPacket defines a rpc handler method for MsgRecvPacket.
func (k *Keeper) RecvPacket(goCtx context.Context, msg *channeltypes.MsgRecvPacket) (*channeltypes.MsgRecvPacketResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve callbacks from router

// Perform TAO verification
//
// If the packet was already received, perform a no-op
// Use a cached context to prevent accidental state changes

// Perform application logic callback
//
// Cache context so that we may discard state changes from callback if the acknowledgement is unsuccessful.

// write application state changes for asynchronous and successful acknowledgements

// Modify events in cached context to reflect unsuccessful acknowledgement

// Set packet acknowledgement only if the acknowledgement is not nil.
// NOTE: IBC applications modules may call the WriteAcknowledgement asynchronously if the
// acknowledgement is nil.

// Timeout defines a rpc handler method for MsgTimeout.
func (k *Keeper) Timeout(goCtx context.Context, msg *channeltypes.MsgTimeout) (*channeltypes.MsgTimeoutResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve callbacks from router

// Perform TAO verification
//
// If the timeout was already received, perform a no-op
// Use a cached context to prevent accidental state changes

// Perform application logic callback

// TimeoutOnClose defines a rpc handler method for MsgTimeoutOnClose.
func (k *Keeper) TimeoutOnClose(goCtx context.Context, msg *channeltypes.MsgTimeoutOnClose) (*channeltypes.MsgTimeoutOnCloseResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Perform TAO verification
//
// If the timeout was already received, perform a no-op
// Use a cached context to prevent accidental state changes

// Perform application logic callback
//
// NOTE: MsgTimeout and MsgTimeoutOnClose use the same "OnTimeoutPacket"
// application logic callback.

// Acknowledgement defines a rpc handler method for MsgAcknowledgement.
func (k *Keeper) Acknowledgement(goCtx context.Context, msg *channeltypes.MsgAcknowledgement) (*channeltypes.MsgAcknowledgementResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve callbacks from router

// Perform TAO verification
//
// If the acknowledgement was already received, perform a no-op
// Use a cached context to prevent accidental state changes

// Perform application logic callback

// UpdateClientParams defines a rpc handler method for MsgUpdateParams.
func (k *Keeper) UpdateClientParams(goCtx context.Context, msg *clienttypes.MsgUpdateParams) (*clienttypes.MsgUpdateParamsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateConnectionParams defines a rpc handler method for MsgUpdateParams for the 03-connection submodule.
func (k *Keeper) UpdateConnectionParams(goCtx context.Context, msg *connectiontypes.MsgUpdateParams) (*connectiontypes.MsgUpdateParamsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateClientConfig defines an rpc handler method for MsgUpdateClientConfig for the 02-client v2 submodule.
func (k *Keeper) UpdateClientConfig(goCtx context.Context, msg *clientv2types.MsgUpdateClientConfig) (*clientv2types.MsgUpdateClientConfigResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteClientCreator defines an rpc handler method for MsgDeleteClientCreator for the 02-client v1 submodule.
func (k *Keeper) DeleteClientCreator(goCtx context.Context, msg *clienttypes.MsgDeleteClientCreator) (*clienttypes.MsgDeleteClientCreatorResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check authorization
