package ibctesting

import (
	abci "github.com/cometbft/cometbft/abci/types"

	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
	connectiontypes "github.com/cosmos/ibc-go/v11/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
	commitmenttypesv2 "github.com/cosmos/ibc-go/v11/modules/core/23-commitment/types/v2"
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

// Endpoint is a which represents a channel endpoint and its associated
// client and connections. It contains client, connection, and channel
// configuration parameters. Endpoint functions will utilize the parameters
// set in the configuration structs when executing IBC messages.
type Endpoint struct {
	Chain        *TestChain
	Counterparty *Endpoint
	ClientID     string
	ConnectionID string
	ChannelID    string

	ClientConfig     ClientConfig
	ConnectionConfig *ConnectionConfig
	ChannelConfig    *ChannelConfig

	MerklePathPrefix commitmenttypesv2.MerklePath
	// disableUniqueChannelIDs is used to enforce, in a test,
	// the old way to generate channel IDs (all channels are called channel-0)
	// It is used only by one test suite and should not be used for new tests.
	disableUniqueChannelIDs bool
}

// NewEndpoint constructs a new endpoint without the counterparty.
// CONTRACT: the counterparty endpoint must be set by the caller.
func NewEndpoint(
	chain *TestChain, clientConfig ClientConfig,
	connectionConfig *ConnectionConfig, channelConfig *ChannelConfig,
) *Endpoint {
	_ = "STUB: not implemented"
	return nil
}

// NewDefaultEndpoint constructs a new endpoint using default values.
// CONTRACT: the counterparty endpoitn must be set by the caller.
func NewDefaultEndpoint(chain *TestChain) *Endpoint { _ = "STUB: not implemented"; return nil }

// QueryProof queries proof associated with this endpoint using the latest client state
// height on the counterparty chain.
func (ep *Endpoint) QueryProof(key []byte) ([]byte, clienttypes.Height) {
	_ = "STUB: not implemented"
	// obtain the counterparty client height.
	return nil, *new(clienttypes.Height)
}

// query proof on the counterparty using the latest height of the IBC client

// QueryProofAtHeight queries proof associated with this endpoint using the proof height
// provided
func (ep *Endpoint) QueryProofAtHeight(key []byte, height uint64) ([]byte, clienttypes.Height) {
	_ = "STUB: not implemented"
	// query proof on the counterparty using the latest height of the IBC client
	return nil, *new(clienttypes.Height)
}

// CreateClient creates an IBC client on the ep. It will update the
// clientID for the endpoint if the message is successfully executed.
// NOTE: a solo machine client will be created with an empty diversifier.
func (ep *Endpoint) CreateClient() error {
	_ = "STUB: not implemented"
	// ensure counterparty has committed state
	return nil
}

// TODO
//		solo := NewSolomachine(ep.Chain.TB, ep.Chain.Codec, clientID, "", 1)
//		clientState = solo.ClientState()
//		consensusState = solo.ConsensusState()

// UpdateClient updates the IBC client associated with the ep.
func (ep *Endpoint) UpdateClient() error {
	_ = "STUB: not implemented"
	// ensure counterparty has committed state
	return nil
}

// FreezeClient freezes the IBC client associated with the ep.
func (ep *Endpoint) FreezeClient() { _ = "STUB: not implemented"; return }

// UpgradeChain will upgrade a chain's chainID to the next revision number.
// It will also update the counterparty client.
// TODO: implement actual upgrade chain functionality via scheduling an upgrade
// and upgrading the client via MsgUpgradeClient
// see reference https://github.com/cosmos/ibc-go/pull/1169
func (ep *Endpoint) UpgradeChain() error { _ = "STUB: not implemented"; return nil }

// increment revision number in chainID

// update chain

// commit changes

// update counterparty client manually

// ensure the next update isn't identical to the one set in state

// ConnOpenInit will construct and execute a MsgConnectionOpenInit on the associated ep.
func (ep *Endpoint) ConnOpenInit() error { _ = "STUB: not implemented"; return nil }

// ConnOpenTry will construct and execute a MsgConnectionOpenTry on the associated ep.
func (ep *Endpoint) ConnOpenTry() error { _ = "STUB: not implemented"; return nil }

// ConnOpenAck will construct and execute a MsgConnectionOpenAck on the associated ep.
func (ep *Endpoint) ConnOpenAck() error { _ = "STUB: not implemented"; return nil }

// testing doesn't use flexible selection

// ConnOpenConfirm will construct and execute a MsgConnectionOpenConfirm on the associated ep.
func (ep *Endpoint) ConnOpenConfirm() error { _ = "STUB: not implemented"; return nil }

// QueryConnectionHandshakeProof returns all the proofs necessary to execute OpenTry or Open Ack of
// the connection handshakes. It returns the proof of the counterparty connection and the proof height.
func (ep *Endpoint) QueryConnectionHandshakeProof() (
	[]byte, clienttypes.Height,
) {
	_ = "STUB: not implemented"
	// query proof for the connection on the counterparty
	return nil, *new(clienttypes.Height)
}

var sequenceNumber int

// IncrementNextChannelSequence incrementes the value "nextChannelSequence" in the store,
// which is used to determine the next channel ID.
// This guarantees that we'll have always different IDs while running tests.
func (ep *Endpoint) IncrementNextChannelSequence() { _ = "STUB: not implemented"; return }

// ChanOpenInit will construct and execute a MsgChannelOpenInit on the associated ep.
func (ep *Endpoint) ChanOpenInit() error { _ = "STUB: not implemented"; return nil }

// update version to selected app version
// NOTE: this update must be performed after SendMsgs()

// ChanOpenTry will construct and execute a MsgChannelOpenTry on the associated ep.
func (ep *Endpoint) ChanOpenTry() error { _ = "STUB: not implemented"; return nil }

// update version to selected app version
// NOTE: this update must be performed after the endpoint channelID is set

// ChanOpenAck will construct and execute a MsgChannelOpenAck on the associated ep.
func (ep *Endpoint) ChanOpenAck() error { _ = "STUB: not implemented"; return nil }

// testing doesn't use flexible selection

// ChanOpenConfirm will construct and execute a MsgChannelOpenConfirm on the associated ep.
func (ep *Endpoint) ChanOpenConfirm() error { _ = "STUB: not implemented"; return nil }

// ChanCloseInit will construct and execute a MsgChannelCloseInit on the associated ep.
//
// NOTE: does not work with ibc-transfer module
func (ep *Endpoint) ChanCloseInit() error { _ = "STUB: not implemented"; return nil }

// SendPacket sends a packet through the channel keeper using the associated endpoint
// The counterparty client is updated so proofs can be sent to the counterparty chain.
// The packet sequence generated for the packet to be sent is returned. An error
// is returned if one occurs.
func (ep *Endpoint) SendPacket(
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
	data []byte,
) (uint64, error) {
	_ = "STUB: not implemented"
	// no need to send message, acting as a module
	return 0, nil
}

// commit changes since no message was sent

// RecvPacket receives a packet on the associated ep.
// The counterparty client is updated.
func (ep *Endpoint) RecvPacket(packet channeltypes.Packet) error {
	_ = "STUB: not implemented"
	return nil
}

// RecvPacketWithResult receives a packet on the associated endpoint and the result
// of the transaction is returned. The counterparty client is updated.
func (ep *Endpoint) RecvPacketWithResult(packet channeltypes.Packet) (*abci.ExecTxResult, error) {
	_ = "STUB: not implemented"
	// get proof of packet commitment on source
	return nil, nil
}

// receive on counterparty and update source client

// WriteAcknowledgement writes an acknowledgement on the channel associated with the ep.
// The counterparty client is updated.
func (ep *Endpoint) WriteAcknowledgement(ack exported.Acknowledgement, packet exported.PacketI) error {
	_ = "STUB: not implemented"
	// no need to send message, acting as a handler
	return nil
}

// commit changes since no message was sent

// AcknowledgePacket sends a MsgAcknowledgement to the channel associated with the ep.
func (ep *Endpoint) AcknowledgePacket(packet channeltypes.Packet, ack []byte) error {
	_ = "STUB: not implemented"
	// get proof of acknowledgement on counterparty
	return nil
}

// AcknowledgePacketWithResult sends a MsgAcknowledgement to the channel associated with the endpoint and returns the result.
func (ep *Endpoint) AcknowledgePacketWithResult(packet channeltypes.Packet, ack []byte) (*abci.ExecTxResult, error) {
	_ = "STUB: not implemented"
	// get proof of acknowledgement on counterparty
	return nil, nil
}

// TimeoutPacketWithResult sends a MsgTimeout to the channel associated with the ep.
func (ep *Endpoint) TimeoutPacketWithResult(packet channeltypes.Packet) (*abci.ExecTxResult, error) {
	_ = "STUB: not implemented"
	// get proof for timeout based on channel order
	return nil, nil
}

// TimeoutPacket sends a MsgTimeout to the channel associated with the ep.
func (ep *Endpoint) TimeoutPacket(packet channeltypes.Packet) error {
	_ = "STUB: not implemented"
	return nil
}

// TimeoutOnClose sends a MsgTimeoutOnClose to the channel associated with the ep.
func (ep *Endpoint) TimeoutOnClose(packet channeltypes.Packet) error {
	_ = "STUB: not implemented"
	// get proof for timeout based on channel order
	return nil
}

// Deprecated: usage of this function should be replaced by `UpdateChannel`
// SetChannelState sets a channel state
func (ep *Endpoint) SetChannelState(state channeltypes.State) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateChannel updates the channel associated with the given ep. It accepts a
// closure which takes a channel allowing the caller to modify its fields.
func (ep *Endpoint) UpdateChannel(updater func(channel *channeltypes.Channel)) {
	_ = "STUB: not implemented"
	return
}

// GetClientLatestHeight returns the latest height for the client state for this ep.
// The client state is expected to exist otherwise testing will fail.
func (ep *Endpoint) GetClientLatestHeight() exported.Height {
	_ = "STUB: not implemented"
	return *new(exported.Height)
}

// GetClientState retrieves the client state for this ep. The
// client state is expected to exist otherwise testing will fail.
func (ep *Endpoint) GetClientState() exported.ClientState {
	_ = "STUB: not implemented"
	return *new(exported.ClientState)
}

// SetClientState sets the client state for this ep.
func (ep *Endpoint) SetClientState(clientState exported.ClientState) {
	_ = "STUB: not implemented"
	return
}

// GetConsensusState retrieves the Consensus State for this endpoint at the provided height.
// The consensus state is expected to exist otherwise testing will fail.
func (ep *Endpoint) GetConsensusState(height exported.Height) exported.ConsensusState {
	_ = "STUB: not implemented"
	return *new(exported.ConsensusState)
}

// SetConsensusState sets the consensus state for this ep.
func (ep *Endpoint) SetConsensusState(consensusState exported.ConsensusState, height exported.Height) {
	_ = "STUB: not implemented"
	return
}

// GetConnection retrieves an IBC Connection for the ep. The
// connection is expected to exist otherwise testing will fail.
func (ep *Endpoint) GetConnection() connectiontypes.ConnectionEnd {
	_ = "STUB: not implemented"
	return *new(connectiontypes.ConnectionEnd)
}

// SetConnection sets the connection for this ep.
func (ep *Endpoint) SetConnection(connection connectiontypes.ConnectionEnd) {
	_ = "STUB: not implemented"
	return
}

// GetChannel retrieves an IBC Channel for the ep. The channel
// is expected to exist otherwise testing will fail.
func (ep *Endpoint) GetChannel() channeltypes.Channel {
	_ = "STUB: not implemented"
	return *new(channeltypes.Channel)
}

// SetChannel sets the channel for this ep.
func (ep *Endpoint) SetChannel(channel channeltypes.Channel) { _ = "STUB: not implemented"; return }

// QueryClientStateProof performs and abci query for a client stat associated
// with this endpoint and returns the ClientState along with the proof.
func (ep *Endpoint) QueryClientStateProof() (exported.ClientState, []byte) {
	_ = "STUB: not implemented"
	// retrieve client state to provide proof for
	return *new(exported.ClientState), nil
}

// UpdateConnection updates the connection associated with the given ep. It accepts a
// closure which takes a connection allowing the caller to modify the connection fields.
func (ep *Endpoint) UpdateConnection(updater func(connection *connectiontypes.ConnectionEnd)) {
	_ = "STUB: not implemented"
	return
}
