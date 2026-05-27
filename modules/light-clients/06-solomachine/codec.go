package solomachine

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
)

// RegisterInterfaces register the ibc channel submodule interfaces to protobuf
// Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) { _ = "STUB: not implemented"; return }

func UnmarshalSignatureData(cdc codec.BinaryCodec, data []byte) (signing.SignatureData, error) {
	_ = "STUB: not implemented"
	return *new(signing.SignatureData), nil
}
