package ibctesting

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	transfertypes "github.com/cosmos/ibc-go/v11/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
	commitmenttypesv2 "github.com/cosmos/ibc-go/v11/modules/core/23-commitment/types/v2"
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
	solomachine "github.com/cosmos/ibc-go/v11/modules/light-clients/06-solomachine"
)

var (
	clientIDSolomachine     = "client-on-solomachine"     // clientID generated on solo machine side
	connectionIDSolomachine = "connection-on-solomachine" // connectionID generated on solo machine side
	channelIDSolomachine    = "channel-on-solomachine"    // channelID generated on solo machine side
)

// DefaultSolomachineClientID is the default solo machine client id used for testing
var DefaultSolomachineClientID = "06-solomachine-0"

// Solomachine is a testing helper used to simulate a counterparty
// solo machine client.
type Solomachine struct {
	t *testing.T

	cdc         codec.BinaryCodec
	ClientID    string
	PrivateKeys []cryptotypes.PrivKey // keys used for signing
	PublicKeys  []cryptotypes.PubKey  // keys used for generating solo machine pub key
	PublicKey   cryptotypes.PubKey    // key used for verification
	Sequence    uint64
	Time        uint64
	Diversifier string
}

// NewSolomachine returns a new solomachine instance with an `nKeys` amount of
// generated private/public key pairs and a sequence starting at 1. If nKeys
// is greater than 1 then a multisig public key is used.
func NewSolomachine(t *testing.T, cdc codec.BinaryCodec, clientID, diversifier string, nKeys uint64) *Solomachine {
	_ = "STUB: not implemented"
	return nil
}

// GenerateKeys generates a new set of secp256k1 private keys and public keys.
// If the number of keys is greater than one then the public key returned represents
// a multisig public key. The private keys are used for signing, the public
// keys are used for generating the public key and the public key is used for
// solo machine verification. The usage of secp256k1 is entirely arbitrary.
// The key type can be swapped for any key type supported by the PublicKey
// interface, if needed. The same is true for the amino based Multisignature
// public key.
func GenerateKeys(t *testing.T, n uint64) ([]cryptotypes.PrivKey, []cryptotypes.PubKey, cryptotypes.PubKey) {
	_ = "STUB: not implemented"
	return nil, nil, *new(cryptotypes.PubKey)
}

// generate multi sig pk

// ClientState returns a new solo machine ClientState instance.
func (solo *Solomachine) ClientState() *solomachine.ClientState {
	_ = "STUB: not implemented"
	return nil
}

// ConsensusState returns a new solo machine ConsensusState instance
func (solo *Solomachine) ConsensusState() *solomachine.ConsensusState {
	_ = "STUB: not implemented"
	return nil
}

// GetHeight returns an exported.Height with Sequence as RevisionHeight
func (solo *Solomachine) GetHeight() exported.Height {
	_ = "STUB: not implemented"
	return *new(exported.Height)
}

// CreateClient creates an on-chain client on the provided chain.
func (solo *Solomachine) CreateClient(chain *TestChain) string {
	_ = "STUB: not implemented"
	return ""
}

// UpdateClient sends a MsgUpdateClient to the provided chain and updates the given clientID.
func (solo *Solomachine) UpdateClient(chain *TestChain, clientID string) {
	_ = "STUB: not implemented"
	return
}

// CreateHeader generates a new private/public key pair and creates the
// necessary signature to construct a valid solo machine header.
// A new diversifier will be used as well
func (solo *Solomachine) CreateHeader(newDiversifier string) *solomachine.Header {
	_ = "STUB: not implemented"
	// generate new private keys and signature for header
	return nil
}

// assumes successful header update

// CreateMisbehaviour constructs testing misbehaviour for the solo machine client
// by signing over two different data bytes at the same sequence.
func (solo *Solomachine) CreateMisbehaviour() *solomachine.Misbehaviour {
	_ = "STUB: not implemented"
	return nil
}

// misbehaviour signaturess can have different timestamps

// ConnOpenInit initializes a connection on the provided chain given a solo machine clientID.
func (solo *Solomachine) ConnOpenInit(chain *TestChain, clientID string) string {
	_ = "STUB: not implemented"
	return ""
}

// clientID generated on solo machine side

// ConnOpenAck performs the connection open ack handshake step on the tendermint chain for the associated
// solo machine client.
func (solo *Solomachine) ConnOpenAck(chain *TestChain, clientID, connectionID string) {
	_ = "STUB: not implemented"
	return
}

// ChanOpenInit initializes a channel on the provided chain given a solo machine connectionID.
func (solo *Solomachine) ChanOpenInit(chain *TestChain, connectionID string) string {
	_ = "STUB: not implemented"
	return ""
}

// ChanOpenAck performs the channel open ack handshake step on the tendermint chain for the associated
// solo machine client.
func (solo *Solomachine) ChanOpenAck(chain *TestChain, channelID string) {
	_ = "STUB: not implemented"
	return
}

// ChanCloseConfirm performs the channel close confirm handshake step on the tendermint chain for the associated
// solo machine client.
func (solo *Solomachine) ChanCloseConfirm(chain *TestChain, portID, channelID string) {
	_ = "STUB: not implemented"
	return
}

// SendTransfer constructs a MsgTransfer and sends the message to the given chain. Any number of optional
// functions can be provided which will modify the MsgTransfer before SendMsgs is called.
func (solo *Solomachine) SendTransfer(chain *TestChain, portID, channelID string, fns ...func(*transfertypes.MsgTransfer)) channeltypes.Packet {
	_ = "STUB: not implemented"
	return *new(channeltypes.Packet)
}

// RecvPacket creates a commitment proof and broadcasts a new MsgRecvPacket.
func (solo *Solomachine) RecvPacket(chain *TestChain, packet channeltypes.Packet) {
	_ = "STUB: not implemented"
	return
}

// AcknowledgePacket creates an acknowledgement proof and broadcasts a MsgAcknowledgement.
func (solo *Solomachine) AcknowledgePacket(chain *TestChain, packet channeltypes.Packet) {
	_ = "STUB: not implemented"
	return
}

// TimeoutPacket creates a unreceived packet proof and broadcasts a MsgTimeout.
func (solo *Solomachine) TimeoutPacket(chain *TestChain, packet channeltypes.Packet) {
	_ = "STUB: not implemented"
	return
}

// nextSequenceRecv is unused for UNORDERED channels

// TimeoutPacketOnClose creates a channel closed and unreceived packet proof and broadcasts a MsgTimeoutOnClose.
func (solo *Solomachine) TimeoutPacketOnClose(chain *TestChain, packet channeltypes.Packet, channelID string) {
	_ = "STUB: not implemented"
	return
}

// nextSequenceRecv is unused for UNORDERED channels

// GenerateSignature uses the stored private keys to generate a signature
// over the sign bytes with each key. If the amount of keys is greater than
// 1 then a multisig data type is returned.
func (solo *Solomachine) GenerateSignature(signBytes []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// single public key

// generate multi signature data

// GenerateProof takes in solo machine sign bytes, generates a signature and marshals it as a proof.
// The solo machine sequence is incremented.
func (solo *Solomachine) GenerateProof(signBytes *solomachine.SignBytes) []byte {
	_ = "STUB: not implemented"
	return nil
}

// GenerateConnOpenTryProof generates the proofTry required for the connection open ack handshake step.
// The clientID, connectionID provided represent the clientID and connectionID created on the counterparty chain, that is the tendermint chain.
func (solo *Solomachine) GenerateConnOpenTryProof(counterpartyClientID, counterpartyConnectionID string) []byte {
	_ = "STUB: not implemented"
	return nil
}

// GenerateChanOpenTryProof generates the proofTry required for the channel open ack handshake step.
// The channelID provided represents the channelID created on the counterparty chain, that is the tendermint chain.
func (solo *Solomachine) GenerateChanOpenTryProof(portID, version, counterpartyChannelID string) []byte {
	_ = "STUB: not implemented"
	return nil
}

// GenerateChanClosedProof generates a channel closed proof.
// The channelID provided represents the channelID created on the counterparty chain, that is the tendermint chain.
func (solo *Solomachine) GenerateChanClosedProof(portID, version, counterpartyChannelID string) []byte {
	_ = "STUB: not implemented"
	return nil
}

// GenerateCommitmentProof generates a commitment proof for the provided packet.
func (solo *Solomachine) GenerateCommitmentProof(packet channeltypes.Packet) []byte {
	_ = "STUB: not implemented"
	return nil
}

// GenerateAcknowledgementProof generates an acknowledgement proof.
func (solo *Solomachine) GenerateAcknowledgementProof(packet channeltypes.Packet) []byte {
	_ = "STUB: not implemented"
	return nil
}

// GenerateReceiptAbsenceProof generates a receipt absence proof for the provided packet.
func (solo *Solomachine) GenerateReceiptAbsenceProof(packet channeltypes.Packet) []byte {
	_ = "STUB: not implemented"
	return nil
}

// GetClientStatePath returns the commitment path for the client state.
func (solo *Solomachine) GetClientStatePath(counterpartyClientIdentifier string) commitmenttypesv2.MerklePath {
	_ = "STUB: not implemented"
	return *new(commitmenttypesv2.MerklePath)
}

// GetConsensusStatePath returns the commitment path for the consensus state.
func (solo *Solomachine) GetConsensusStatePath(counterpartyClientIdentifier string, consensusHeight exported.Height) commitmenttypesv2.MerklePath {
	_ = "STUB: not implemented"
	return *new(commitmenttypesv2.MerklePath)
}

// GetConnectionStatePath returns the commitment path for the connection state.
func (solo *Solomachine) GetConnectionStatePath(connID string) commitmenttypesv2.MerklePath {
	_ = "STUB: not implemented"
	return *new(commitmenttypesv2.MerklePath)
}

// GetChannelStatePath returns the commitment path for that channel state.
func (solo *Solomachine) GetChannelStatePath(portID, channelID string) commitmenttypesv2.MerklePath {
	_ = "STUB: not implemented"
	return *new(commitmenttypesv2.MerklePath)
}

// GetPacketCommitmentPath returns the commitment path for a packet commitment.
func (solo *Solomachine) GetPacketCommitmentPath(portID, channelID string, sequence uint64) commitmenttypesv2.MerklePath {
	_ = "STUB: not implemented"
	return *new(commitmenttypesv2.MerklePath)
}

// GetPacketAcknowledgementPath returns the commitment path for a packet acknowledgement.
func (solo *Solomachine) GetPacketAcknowledgementPath(portID, channelID string, sequence uint64) commitmenttypesv2.MerklePath {
	_ = "STUB: not implemented"
	return *new(commitmenttypesv2.MerklePath)
}

// GetPacketReceiptPath returns the commitment path for a packet receipt
// and an absent receipts.
func (solo *Solomachine) GetPacketReceiptPath(portID, channelID string, sequence uint64) commitmenttypesv2.MerklePath {
	_ = "STUB: not implemented"
	return *new(commitmenttypesv2.MerklePath)
}

// GetNextSequenceRecvPath returns the commitment path for the next sequence recv counter.
func (solo *Solomachine) GetNextSequenceRecvPath(portID, channelID string) commitmenttypesv2.MerklePath {
	_ = "STUB: not implemented"
	return *new(commitmenttypesv2.MerklePath)
}
