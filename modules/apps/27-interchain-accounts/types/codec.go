package types

import (
	"github.com/cosmos/gogoproto/proto"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ModuleCdc references the global interchain accounts module codec. Note, the codec
// should ONLY be used in certain instances of tests and for JSON encoding.
//
// The actual codec used for serialization should be provided to interchain accounts and
// defined at the application level.
var ModuleCdc = codec.NewProtoCodec(codectypes.NewInterfaceRegistry())

// RegisterInterfaces registers the interchain accounts controller types and the concrete InterchainAccount implementation
// against the associated x/auth AccountI and GenesisAccount interfaces.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) { _ = "STUB: not implemented"; return }

// SerializeCosmosTx serializes a slice of sdk.Msg's using the CosmosTx type. The sdk.Msg's are
// packed into Any's and inserted into the Messages field of a CosmosTx. The CosmosTx is marshaled
// depending on the encoding type passed in. The marshaled bytes are returned. Only the ProtoCodec
// is supported for serializing messages. Both protobuf and proto3 JSON are supported.
func SerializeCosmosTx(cdc codec.Codec, msgs []proto.Message, encoding string) ([]byte, error) {
	_ = "STUB: not implemented"
	// this is a defensive check to ensure only the ProtoCodec is used for message serialization
	return nil, nil
}

// DeserializeCosmosTx unmarshals and unpacks a slice of transaction bytes into a slice of sdk.Msg's.
// The transaction bytes are unmarshaled depending on the encoding type passed in. The sdk.Msg's are
// unpacked from Any's and returned. Only the ProtoCodec is supported for serializing messages. Both
// protobuf and proto3 JSON are supported.
func DeserializeCosmosTx(cdc codec.Codec, data []byte, encoding string) ([]sdk.Msg, error) {
	_ = "STUB: not implemented"
	// this is a defensive check to ensure only the ProtoCodec is used for message deserialization
	return nil, nil
}

func equalJSON(a, b []byte) (bool, error) { _ = "STUB: not implemented"; return false, nil }
