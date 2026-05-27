package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

// ChanOpenInit is called by a module to initiate a channel opening handshake with
// a module on another chain. The counterparty channel identifier is validated to be
// empty in msg validation.
func (k *Keeper) ChanOpenInit(
	ctx sdk.Context,
	order types.Order,
	connectionHops []string,
	portID string,
	counterparty types.Counterparty,
	version string,
) (string, error) {
	_ = "STUB: not implemented"
	// connection hop length checked on msg.ValidateBasic()
	return "", nil
}

// WriteOpenInitChannel writes a channel which has successfully passed the OpenInit handshake step.
// The channel is set in state and all the associated Send and Recv sequences are set to 1.
// An event is emitted for the handshake step.
func (k *Keeper) WriteOpenInitChannel(
	ctx sdk.Context,
	portID,
	channelID string,
	order types.Order,
	connectionHops []string,
	counterparty types.Counterparty,
	version string,
) {
	_ = "STUB: not implemented"
	return
}

// ChanOpenTry is called by a module to accept the first step of a channel opening
// handshake initiated by a module on another chain.
func (k *Keeper) ChanOpenTry(
	ctx sdk.Context,
	order types.Order,
	connectionHops []string,
	portID string,
	counterparty types.Counterparty,
	counterpartyVersion string,
	initProof []byte,
	proofHeight exported.Height,
) (string, error) {
	_ = "STUB: not implemented"
	// connection hops only supports a single connection
	return "", nil
}

// generate a new channel

// expectedCounterpaty is the counterparty of the counterparty's channel end
// (i.e self)

// WriteOpenTryChannel writes a channel which has successfully passed the OpenTry handshake step.
// The channel is set in state. If a previous channel state did not exist, all the Send and Recv
// sequences are set to 1. An event is emitted for the handshake step.
func (k *Keeper) WriteOpenTryChannel(
	ctx sdk.Context,
	portID,
	channelID string,
	order types.Order,
	connectionHops []string,
	counterparty types.Counterparty,
	version string,
) {
	_ = "STUB: not implemented"
	return
}

// ChanOpenAck is called by the handshake-originating module to acknowledge the
// acceptance of the initial request by the counterparty module on the other chain.
func (k *Keeper) ChanOpenAck(
	ctx sdk.Context,
	portID,
	channelID string,
	counterpartyVersion,
	counterpartyChannelID string,
	tryProof []byte,
	proofHeight exported.Height,
) error {
	_ = "STUB: not implemented"
	return nil
}

// counterparty of the counterparty channel end (i.e self)

// WriteOpenAckChannel writes an updated channel state for the successful OpenAck handshake step.
// An event is emitted for the handshake step.
func (k *Keeper) WriteOpenAckChannel(
	ctx sdk.Context,
	portID,
	channelID,
	counterpartyVersion,
	counterpartyChannelID string,
) {
	_ = "STUB: not implemented"
	return
}

// get the counterparty and set it in the client keeper v2 to support IBC v2 on this
// channel ID through aliasing
// NOTE: This should never error as the channel is set in the line above

// ChanOpenConfirm is called by the handshake-accepting module to confirm the acknowledgement
// of the handshake-originating module on the other chain and finish the channel opening handshake.
func (k *Keeper) ChanOpenConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
	ackProof []byte,
	proofHeight exported.Height,
) error {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: If the counterparty has initialized an upgrade in the same block as performing the
// ACK handshake step, this channel end will be incapable of opening.

// WriteOpenConfirmChannel writes an updated channel state for the successful OpenConfirm handshake step.
// An event is emitted for the handshake step.
func (k *Keeper) WriteOpenConfirmChannel(
	ctx sdk.Context,
	portID,
	channelID string,
) {
	_ = "STUB: not implemented"
	return
}

// get the counterparty and set it in the client keeper v2 to support IBC v2 on this
// channel ID through aliasing
// NOTE: This should never error as the channel is set in the line above

// Closing Handshake
//
// This section defines the set of functions required to close a channel handshake
// as defined in https://github.com/cosmos/ibc/tree/master/spec/core/ics-004-channel-and-packet-semantics#closing-handshake
//
// ChanCloseInit is called by either module to close their end of the channel. Once
// closed, channels cannot be reopened.
func (k *Keeper) ChanCloseInit(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// ChanCloseConfirm is called by the counterparty module to close their end of the
// channel, since the other end has been closed.
func (k *Keeper) ChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
	initProof []byte,
	proofHeight exported.Height,
) error {
	_ = "STUB: not implemented"
	return nil
}
