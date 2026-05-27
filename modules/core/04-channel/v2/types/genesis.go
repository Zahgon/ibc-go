package types

// NewPacketState creates a new PacketState instance.
func NewPacketState(clientID string, sequence uint64, data []byte) PacketState {
	_ = "STUB: not implemented"
	return *new(PacketState)
}

// Validate performs basic validation of fields returning an error upon any failure.
func (ps PacketState) Validate() error { _ = "STUB: not implemented"; return nil }

// NewPacketSequence creates a new PacketSequences instance.
func NewPacketSequence(clientID string, sequence uint64) PacketSequence {
	_ = "STUB: not implemented"
	return *new(PacketSequence)
}

// Validate performs basic validation of fields returning an error upon any failure.
func (ps PacketSequence) Validate() error { _ = "STUB: not implemented"; return nil }

// NewGenesisState creates a GenesisState instance.
func NewGenesisState(
	acks, receipts, commitments, asyncPackets []PacketState,
	sendSeqs []PacketSequence,
) GenesisState {
	_ = "STUB: not implemented"
	return *new(GenesisState)
}

// DefaultGenesisState returns the ibc channel v2 submodule's default genesis state.
func DefaultGenesisState() GenesisState { _ = "STUB: not implemented"; return *new(GenesisState) }

// Validate performs basic genesis state validation returning an error upon any failure.
func (gs GenesisState) Validate() error { _ = "STUB: not implemented"; return nil }

func validateGenFields(clientID string, sequence uint64) error {
	_ = "STUB: not implemented"
	return nil
}
