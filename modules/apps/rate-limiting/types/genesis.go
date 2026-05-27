package types

// Splits a pending send packet of the form {channelId}/{sequenceNumber} into the channel Id
// and sequence number respectively
func ParsePendingPacketID(pendingPacketID string) (string, uint64, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState { _ = "STUB: not implemented"; return nil }

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error { _ = "STUB: not implemented"; return nil }

// Verify the epoch hour duration is specified

// If the hour epoch has been initialized already (epoch number != 0), validate and then use it
