package cli

import (
	"github.com/cosmos/gogoproto/proto"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/codec"
)

const (
	memoFlag     string = "memo"
	encodingFlag string = "encoding"
)

func generatePacketDataCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GeneratePacketData takes in message bytes and a memo and serializes the message into an
// instance of InterchainAccountPacketData which is returned as bytes.
func GeneratePacketData(cdc *codec.ProtoCodec, msgBytes []byte, memo string, encoding string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convertBytesIntoProtoMessages returns a list of proto messages from bytes. The bytes can be in the form of a single
// message, or a json array of messages.
func convertBytesIntoProtoMessages(cdc *codec.ProtoCodec, msgBytes []byte) ([]proto.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if we fail to unmarshal a list of messages, we assume we are just dealing with a single message.
// in this case we return a list of a single item.

// generateIcaPacketDataFromProtoMessages generates ica packet data as bytes from a given set of proto encoded sdk messages and a memo.
func generateIcaPacketDataFromProtoMessages(cdc *codec.ProtoCodec, sdkMessages []proto.Message, memo string, encoding string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
