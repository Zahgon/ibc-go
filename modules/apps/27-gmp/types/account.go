package types

import (
	"github.com/cosmos/gogoproto/proto"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewAccountIdentifier creates a new AccountIdentifier with the given clientId, sender, and salt.
func NewAccountIdentifier(clientID, sender string, salt []byte) AccountIdentifier {
	_ = "STUB: not implemented"
	return *new(AccountIdentifier)
}

// NewICS27Account creates a new ICS27Account with the given address and accountId.
func NewICS27Account(addr string, accountID *AccountIdentifier) ICS27Account {
	_ = "STUB: not implemented"
	return *new(ICS27Account)
}

// BuildAddressPredictable generates an account address for the gmp module with len = types.AccountAddrLen using the
// Cosmos SDK address.Module function.
// Internally a key is built containing:
// (len(clientId) | clientId | len(sender) | sender | len(salt) | salt).
//
// All method parameter values must be valid and not nil.
//
// This function was copied from wasmd and modified.
// <https://github.com/CosmWasm/wasmd/blob/632fc333d01a84fa5426de6783f7797ad2825e25/x/wasm/keeper/addresses.go#L49>
func BuildAddressPredictable(accountID *AccountIdentifier) (sdk.AccAddress, error) {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress), nil
}

// uint64LengthPrefix prepend big endian encoded byte length
func uint64LengthPrefix(bz []byte) []byte { _ = "STUB: not implemented"; return nil }

// DeserializeCosmosTx unmarshals and unpacks a slice of transaction bytes into a slice of sdk.Msg's.
// The transaction bytes are unmarshaled depending on the encoding type passed in. The sdk.Msg's are
// unpacked from Any's and returned.
func DeserializeCosmosTx(cdc codec.Codec, data []byte) ([]sdk.Msg, error) {
	_ = "STUB: not implemented"
	// this is a defensive check to ensure only the ProtoCodec is used for message deserialization
	return nil, nil
}

// SerializeCosmosTx serializes a slice of sdk.Msg's using the CosmosTx type. The sdk.Msg's are
// packed into Any's and inserted into the Messages field of a CosmosTx. The CosmosTx is marshaled
// depending on the encoding type passed in. The marshaled bytes are returned.
func SerializeCosmosTx(cdc codec.BinaryCodec, msgs []proto.Message) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func equalJSON(a, b []byte) (bool, error) { _ = "STUB: not implemented"; return false, nil }
