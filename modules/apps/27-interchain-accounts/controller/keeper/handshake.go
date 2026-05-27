package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
)

// OnChanOpenInit performs basic validation of channel initialization.
// The counterparty port identifier must be the host chain representation as defined in the types package,
// the channel version must be equal to the version in the types package,
// there must not be an active channel for the specified port identifier.
func (k *Keeper) OnChanOpenInit(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	counterparty channeltypes.Counterparty,
	version string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// OnChanOpenAck sets the active channel for the interchain account/owner pair
// and stores the associated interchain account address in state keyed by it's corresponding port identifier
func (k *Keeper) OnChanOpenAck(
	ctx sdk.Context,
	portID,
	channelID string,
	counterpartyVersion string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// OnChanCloseConfirm removes the active channel stored in state
func (*Keeper) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	_ = "STUB: not implemented"

	// OnChanUpgradeInit performs the upgrade init step of the channel upgrade handshake.
	// The upgrade init callback must verify the proposed changes to the order, connectionHops, and version.
	// Within the version we have the tx type, encoding, interchain account address, host/controller connectionID's
	// and the ICS27 protocol version.
	//
	// The following may be changed:
	// - tx type (must be supported)
	// - encoding (must be supported)
	// - order
	//
	// The following may not be changed:
	// - connectionHops (and subsequently host/controller connectionIDs)
	// - interchain account address
	// - ICS27 protocol version
	return nil
}

func (k *Keeper) OnChanUpgradeInit(ctx sdk.Context, portID, channelID string, proposedOrder channeltypes.Order, proposedConnectionHops []string, proposedversion string) (string, error) {
	_ = "STUB: not implemented"
	// verify connection hops has not changed
	return "", nil
}

// verify proposed version only modifies tx type or encoding

// ValidateControllerMetadata will ensure the ICS27 protocol version has not changed and that the
// tx type and encoding are supported

// the interchain account address on the host chain
// must remain the same after the upgrade.

// OnChanUpgradeAck implements the ack setup of the channel upgrade handshake.
// The upgrade ack callback must verify the proposed changes to the channel version.
// Within the channel version we have the tx type, encoding, interchain account address, host/controller connectionID's
// and the ICS27 protocol version.
//
// The following may be changed:
// - tx type (must be supported)
// - encoding (must be supported)
//
// The following may not be changed:
// - controller connectionID
// - host connectionID
// - interchain account address
// - ICS27 protocol version
func (k *Keeper) OnChanUpgradeAck(ctx sdk.Context, portID, channelID, counterpartyVersion string) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateControllerMetadata will ensure the ICS27 protocol version has not changed and that the
// tx type and encoding are supported. Note, we pass in the current channel connection hops. The upgrade init
// step will verify that the proposed connection hops will not change.

// the interchain account address on the host chain
// must remain the same after the upgrade.
