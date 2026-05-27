package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

// RegisterLegacyAminoCodec registers the necessary interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) { _ = "STUB: not implemented"; return }

// RegisterInterfaces registers the client interfaces to protobuf Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) { _ = "STUB: not implemented"; return }

// PackClientState constructs a new Any packed with the given client state value. It returns
// an error if the client state can't be casted to a protobuf message or if the concrete
// implementation is not registered to the protobuf codec.
func PackClientState(clientState exported.ClientState) (*codectypes.Any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnpackClientState unpacks an Any into a ClientState. It returns an error if the
// client state can't be unpacked into a ClientState.
func UnpackClientState(protoAny *codectypes.Any) (exported.ClientState, error) {
	_ = "STUB: not implemented"
	return *new(exported.ClientState), nil
}

// PackConsensusState constructs a new Any packed with the given consensus state value. It returns
// an error if the consensus state can't be casted to a protobuf message or if the concrete
// implementation is not registered to the protobuf codec.
func PackConsensusState(consensusState exported.ConsensusState) (*codectypes.Any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MustPackConsensusState calls PackConsensusState and panics on error.
func MustPackConsensusState(consensusState exported.ConsensusState) *codectypes.Any {
	_ = "STUB: not implemented"
	return nil
}

// UnpackConsensusState unpacks an Any into a ConsensusState. It returns an error if the
// consensus state can't be unpacked into a ConsensusState.
func UnpackConsensusState(protoAny *codectypes.Any) (exported.ConsensusState, error) {
	_ = "STUB: not implemented"
	return *new(exported.ConsensusState), nil
}

// PackClientMessage constructs a new Any packed with the given value. It returns
// an error if the value can't be casted to a protobuf message or if the concrete
// implementation is not registered to the protobuf codec.
func PackClientMessage(clientMessage exported.ClientMessage) (*codectypes.Any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnpackClientMessage unpacks an Any into a ClientMessage. It returns an error if the
// consensus state can't be unpacked into a ClientMessage.
func UnpackClientMessage(protoAny *codectypes.Any) (exported.ClientMessage, error) {
	_ = "STUB: not implemented"
	return *new(exported.ClientMessage), nil
}
