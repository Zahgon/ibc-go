package host

const (
	KeyChannelEndPrefix = "channelEnds"
	KeyChannelPrefix    = "channels"
)

// ICS04
// The following paths are the keys to the store as defined in https://github.com/cosmos/ibc/tree/master/spec/core/ics-004-channel-and-packet-semantics#store-paths

// ChannelKey returns the store key for a particular channel
func ChannelKey(portID, channelID string) []byte { _ = "STUB: not implemented"; return nil }

func ChannelPath(portID, channelID string) string { _ = "STUB: not implemented"; return "" }
