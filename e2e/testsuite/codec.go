package testsuite

import (
	"github.com/cosmos/gogoproto/proto"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module/testutil"
)

// Codec returns the global E2E protobuf codec.
func Codec() *codec.ProtoCodec { _ = "STUB: not implemented"; return nil }

// SDKEncodingConfig returns the global E2E encoding config.
func SDKEncodingConfig() *testutil.TestEncodingConfig { _ = "STUB: not implemented"; return nil }

// codecAndEncodingConfig returns the codec and encoding config used in the E2E tests.
// Note: any new types added to the codec must be added here.
func codecAndEncodingConfig() (*codec.ProtoCodec, testutil.TestEncodingConfig) {
	_ = "STUB: not implemented"
	return nil, *new(testutil.TestEncodingConfig)
}

// ibc types

// all other types

// UnmarshalMsgResponses attempts to unmarshal the tx msg responses into the provided message types.
func UnmarshalMsgResponses(txResp sdk.TxResponse, msgs ...proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// MustProtoMarshalJSON provides an auxiliary function to return Proto3 JSON encoded
// bytes of a message. This function should be used when marshalling a proto.Message
// from the e2e tests. This function strips out unknown fields. This is useful for
// backwards compatibility tests where the types imported by the e2e package have
// new fields that older versions do not recognize.
func MustProtoMarshalJSON(msg proto.Message) []byte { _ = "STUB: not implemented"; return nil }

// EmitDefaults is set to false to prevent marshalling of unpopulated fields (memo)
// OrigName and the anyResovler match the fields the original SDK function would expect
// in order to minimize changes.

// OrigName is true since there is no particular reason to use camel case
// The any resolver is empty, but provided anyways.
