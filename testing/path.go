package ibctesting

import (
	abci "github.com/cometbft/cometbft/abci/types"

	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
)

// Path contains two endpoints representing two chains connected over IBC
type Path struct {
	EndpointA *Endpoint
	EndpointB *Endpoint
}

// NewPath constructs an endpoint for each chain using the default values
// for the endpoints. Each endpoint is updated to have a pointer to the
// counterparty endpoint.
func NewPath(chainA, chainB *TestChain) *Path { _ = "STUB: not implemented"; return nil }

// NewTransferPath constructs a new path between each chain suitable for use with
// the transfer module.
func NewTransferPath(chainA, chainB *TestChain) *Path { _ = "STUB: not implemented"; return nil }

// SetChannelOrdered sets the channel order for both endpoints to ORDERED.
func (path *Path) SetChannelOrdered() { _ = "STUB: not implemented"; return }

// DisableUniqueChannelIDs provides an opt-out way to not have all channel IDs be different
// while testing.
func (path *Path) DisableUniqueChannelIDs() *Path { _ = "STUB: not implemented"; return nil }

// RelayPacket attempts to relay the packet first on EndpointA and then on EndpointB
// if EndpointA does not contain a packet commitment for that packet. An error is returned
// if a relay step fails or the packet commitment does not exist on either endpoint.
func (path *Path) RelayPacket(packet channeltypes.Packet) error {
	_ = "STUB: not implemented"
	return nil
}

// RelayPacketWithResults attempts to relay the packet first on EndpointA and then on EndpointB
// if EndpointA does not contain a packet commitment for that packet. The function returns:
// - The result of the packet receive transaction.
// - The acknowledgement written on the receiving chain.
// - An error if a relay step fails or the packet commitment does not exist on either endpoint.
func (path *Path) RelayPacketWithResults(packet channeltypes.Packet) (*abci.ExecTxResult, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// packet found, relay from A to B

// packet found, relay B to A

// Reversed returns a new path with endpoints reversed.
func (path *Path) Reversed() *Path { _ = "STUB: not implemented"; return nil }

// Setup constructs a TM client, connection, and channel on both chains provided. It will
// fail if any error occurs.
func (path *Path) Setup() { _ = "STUB: not implemented"; return }

// channels can also be referenced through the returned connections

// SetupV2 constructs clients on both sides and then provides the counterparties for both sides
// This is all that is necessary for path setup with the IBC v2 protocol
func (path *Path) SetupV2() { _ = "STUB: not implemented"; return }

// SetupClients is a helper function to create clients on both chains. It assumes the
// caller does not anticipate any errors.
func (path *Path) SetupClients() { _ = "STUB: not implemented"; return }

// SetupCounterparties is a helper function to set the counterparties supporting IBC v2 on both
// chains. It assumes the caller does not anticipate any errors.
func (path *Path) SetupCounterparties() { _ = "STUB: not implemented"; return }

// SetupConnections is a helper function to create clients and the appropriate
// connections on both the source and counterparty chain. It assumes the caller does not
// anticipate any errors.
func (path *Path) SetupConnections() { _ = "STUB: not implemented"; return }

// CreateConnections constructs and executes connection handshake messages in order to create
// OPEN connections on chainA and chainB. The function expects the connections to be
// successfully opened otherwise testing will fail.
func (path *Path) CreateConnections() { _ = "STUB: not implemented"; return }

// ensure counterparty is up to date

// CreateChannels constructs and executes channel handshake messages in order to create
// OPEN channels on chainA and chainB. The function expects the channels to be successfully
// opened otherwise testing will fail.
func (path *Path) CreateChannels() { _ = "STUB: not implemented"; return }

// ensure counterparty is up to date
