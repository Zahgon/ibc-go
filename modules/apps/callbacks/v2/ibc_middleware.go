package v2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/apps/callbacks/types"
	channeltypesv2 "github.com/cosmos/ibc-go/v11/modules/core/04-channel/v2/types"
	"github.com/cosmos/ibc-go/v11/modules/core/api"
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var (
	_ api.IBCModule            = (*IBCMiddleware)(nil)
	_ exported.Acknowledgement = (*RecvAcknowledgement)(nil)
)

// Create internal implementation of exported.Acknowledgement
// to pass the app acknowledgement bytes to contractKeeper IBCReceivePacketCallback
// interface
type RecvAcknowledgement []byte

// RecvPacket only passes callback to contract if the acknowledgement
// is successful. Thus, we can just return true here.
func (rack RecvAcknowledgement) Success() bool { _ = "STUB: not implemented"; return false }

// RecvPacket passes the application acknowledgment directly to contract
func (rack RecvAcknowledgement) Acknowledgement() []byte {
	_ = "STUB: not implemented"

	// IBCMiddleware implements the IBC v2 middleware interface
	// with the underlying application.
	return nil
}

type IBCMiddleware struct {
	app             api.PacketUnmarshalerModuleV2
	writeAckWrapper api.WriteAcknowledgementWrapper

	contractKeeper types.ContractKeeper
	chanKeeperV2   types.ChannelKeeperV2

	// maxCallbackGas defines the maximum amount of gas that a callback actor can ask the
	// relayer to pay for. If a callback fails due to insufficient gas, the entire tx
	// is reverted if the relayer hadn't provided the minimum(userDefinedGas, maxCallbackGas).
	// If the actor hasn't defined a gas limit, then it is assumed to be the maxCallbackGas.
	maxCallbackGas uint64
}

// NewIBCMiddleware creates a new IBCMiddleware instance given the keeper and underlying application.
// The underlying application must implement the required callback interfaces.
func NewIBCMiddleware(
	app api.IBCModule, writeAckWrapper api.WriteAcknowledgementWrapper,
	contractKeeper types.ContractKeeper, chanKeeperV2 types.ChannelKeeperV2, maxCallbackGas uint64,
) *IBCMiddleware {
	_ = "STUB: not implemented"
	return nil
}

// WithWriteAckWrapper sets the WriteAcknowledgementWrapper for the middleware.
func (im *IBCMiddleware) WithWriteAckWrapper(writeAckWrapper api.WriteAcknowledgementWrapper) {
	_ = "STUB: not implemented"
	return
}

// GetWriteAckWrapper returns the WriteAckWrapper
func (im *IBCMiddleware) GetWriteAckWrapper() api.WriteAcknowledgementWrapper {
	_ = "STUB: not implemented"
	return *new(api.WriteAcknowledgementWrapper)
}

// OnSendPacket implements source callbacks for sending packets.
// It defers to the underlying application and then calls the contract callback.
// If the contract callback returns an error, panics, or runs out of gas, then
// the packet send is rejected.
func (im *IBCMiddleware) OnSendPacket(
	ctx sdk.Context,
	sourceClient string,
	destinationClient string,
	sequence uint64,
	payload channeltypesv2.Payload,
	signer sdk.AccAddress,
) error {
	_ = "STUB: not implemented"
	return nil
}

// OnSendPacket is not blocked if the packet does not opt-in to callbacks

// OnSendPacket is not blocked if the packet does not opt-in to callbacks

// OnSendPacket is blocked if the packet opts-in to callbacks but the callback data is invalid

// contract keeper is allowed to reject the packet send.

// OnRecvPacket implements the ReceivePacket destination callbacks for the ibc-callbacks middleware during
// synchronous packet acknowledgement.
// It defers to the underlying application and then calls the contract callback.
// If the contract callback runs out of gas and may be retried with a higher gas limit then the state changes are
// reverted via a panic.
func (im *IBCMiddleware) OnRecvPacket(
	ctx sdk.Context,
	sourceClient string,
	destinationClient string,
	sequence uint64,
	payload channeltypesv2.Payload,
	relayer sdk.AccAddress,
) channeltypesv2.RecvPacketResult {
	_ = "STUB: not implemented"
	return *new(channeltypesv2.RecvPacketResult)
}

// if ack is nil (asynchronous acknowledgements), then the callback will be handled in WriteAcknowledgement
// if ack is not successful, all state changes are reverted. If a packet cannot be received, then there is
// no need to execute a callback on the receiving chain.

// OnRecvPacket is not blocked if the packet does not opt-in to callbacks

// OnRecvPacket is not blocked if the packet does not opt-in to callbacks

// OnRecvPacket is blocked if the packet opts-in to callbacks but the callback data is invalid

// reconstruct a channel v1 packet from the v2 packet
// in order to preserve the same interface for the contract keeper

// wrap the individual acknowledgement into the RecvAcknowledgement since it implements the exported.Acknowledgement interface
// since we return early on failure, we are guaranteed that the ack is a successful acknowledgement

// callback execution errors are not allowed to block the packet lifecycle, they are only used in event emissions

// OnAcknowledgementPacket implements source callbacks for acknowledgement packets.
// It defers to the underlying application and then calls the contract callback.
// If the contract callback runs out of gas and may be retried with a higher gas limit then the state changes are
// reverted via a panic.
func (im *IBCMiddleware) OnAcknowledgementPacket(
	ctx sdk.Context,
	sourceClient string,
	destinationClient string,
	sequence uint64,
	acknowledgement []byte,
	payload channeltypesv2.Payload,
	relayer sdk.AccAddress,
) error {
	_ = "STUB: not implemented"
	// we first call the underlying app to handle the acknowledgement
	return nil
}

// OnAcknowledgementPacket is not blocked if the packet does not opt-in to callbacks

// OnAcknowledgementPacket is not blocked if the packet does not opt-in to callbacks

// OnAcknowledgementPacket is blocked if the packet opts-in to callbacks but the callback data is invalid
// This should never occur since this error is already checked `OnSendPacket`

// reconstruct a channel v1 packet from the v2 packet
// in order to preserve the same interface for the contract keeper

// NOTE: The callback is receiving the acknowledgement that the application received for its particular payload.
// In the case of a successful acknowledgement, this will be the acknowledgement sent by the counterparty application for the given payload
// In the case of an error acknowledgement, this will be the sentinel error acknowledgement bytes defined by IBC v2 protocol.
// Thus, the contract must be aware that the sentinel error acknowledgement signals a failed receive
// and the contract must handle this error case and the corresponding success case (ie ack != ErrorAcknowledgement) accordingly.

// callback execution errors are not allowed to block the packet lifecycle, they are only used in event emissions

// OnTimeoutPacket implements timeout source callbacks for the ibc-callbacks middleware.
// It defers to the underlying application and then calls the contract callback.
// If the contract callback runs out of gas and may be retried with a higher gas limit then the state changes are
// reverted via a panic.
// OnTimeoutPacket is executed when a packet has timed out on the receiving chain.
func (im *IBCMiddleware) OnTimeoutPacket(
	ctx sdk.Context,
	sourceClient string,
	destinationClient string,
	sequence uint64,
	payload channeltypesv2.Payload,
	relayer sdk.AccAddress,
) error {
	_ = "STUB: not implemented"
	return nil
}

// OnTimeoutPacket is not blocked if the packet does not opt-in to callbacks

// OnTimeoutPacket is not blocked if the packet does not opt-in to callbacks

// OnTimeoutPacket is blocked if the packet opts-in to callbacks but the callback data is invalid
// This should never occur since this error is already checked `OnSendPacket`

// reconstruct a channel v1 packet from the v2 packet
// in order to preserve the same interface for the contract keeper

// callback execution errors are not allowed to block the packet lifecycle, they are only used in event emissions

// WriteAcknowledgement implements the ReceivePacket destination callbacks for the ibc-callbacks middleware
// during asynchronous packet acknowledgement.
// It defers to the underlying application and then calls the contract callback.
// If the contract callback runs out of gas and may be retried with a higher gas limit then the state changes are
// reverted via a panic.
func (im *IBCMiddleware) WriteAcknowledgement(
	ctx sdk.Context,
	clientID string,
	sequence uint64,
	ack channeltypesv2.Acknowledgement,
) error {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: use first payload as the payload that is being handled by callbacks middleware
// must reconsider if multipacket data gets supported with async packets
// TRACKING ISSUE: https://github.com/cosmos/ibc-go/issues/7950

// WriteAcknowledgement is not blocked if the packet does not opt-in to callbacks

// WriteAcknowledgement is blocked if the packet opts-in to callbacks but the callback data is invalid
// This should never occur since this error is already checked `OnRecvPacket`

// reconstruct a channel v1 packet from the v2 packet
// in order to preserve the same interface for the contract keeper

// wrap the individual acknowledgement into the channeltypesv2.Acknowledgement since it implements the exported.Acknowledgement interface

// callback execution errors are not allowed to block the packet lifecycle, they are only used in event emissions
