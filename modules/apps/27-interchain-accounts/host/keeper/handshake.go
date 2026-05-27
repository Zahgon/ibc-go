package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
)

// OnChanOpenTry performs basic validation of the ICA channel
// and registers a new interchain account (if it doesn't exist).
// The version returned will include the registered interchain
// account address.
func (k *Keeper) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID,
	channelID string,
	counterparty channeltypes.Counterparty,
	counterpartyVersion string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Propose the default metadata if the counterparty version is invalid

// set here the HostConnectionId in case the controller did not set it

// if a channel is being reopened, we allow the controller to propose new fields
// which are not exactly the same as the previous. The provided address will
// be overwritten with the correct one before the metadata is returned.

// reopening an interchain account

// OnChanOpenConfirm completes the handshake process by setting the active channel in state on the host chain
func (k *Keeper) OnChanOpenConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// It is assumed the controller chain will not allow multiple active channels to be created for the same connectionID/portID
// If the controller chain does allow multiple active channels to be created for the same connectionID/portID,
// disallowing overwriting the current active channel guarantees the channel can no longer be used as the controller
// and host will disagree on what the currently active channel is

// OnChanCloseConfirm removes the active channel stored in state
func (*Keeper) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	_ = "STUB: not implemented"
	return nil
}
