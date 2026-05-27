package types

const (
	// SubModuleName defines the channelv2 module name.
	SubModuleName = "channelv2"

	// KeyAsyncPacket defines the key to store the async packet.
	KeyAsyncPacket = "async_packet"

	// KeyAlias defines the key to store the alias to base client mapping.
	KeyAlias = "alias"
)

// AsyncPacketKey returns the key under which the packet is stored
// if the receiving application returns an async acknowledgement.
func AsyncPacketKey(clientID string, sequence uint64) []byte { _ = "STUB: not implemented"; return nil }

// AsyncPacketPrefixKey returns the prefix key under which all async packets are stored
// for a given clientID.
func AsyncPacketPrefixKey(clientID string) []byte { _ = "STUB: not implemented"; return nil }

// AliasKey returns the key under which the base clientID will be stored
// for an alias (original v1 channelID)
func AliasKey(alias string) []byte { _ = "STUB: not implemented"; return nil }
