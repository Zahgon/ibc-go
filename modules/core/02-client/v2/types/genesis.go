package types

// DefaultGenesisState returns the ibc client submodule's default genesis state.
func DefaultGenesisState() GenesisState { _ = "STUB: not implemented"; return *new(GenesisState) }

// Validate checks the CounterpartyInfos added to the genesis for validity.
func (gs GenesisState) Validate() error { _ = "STUB: not implemented"; return nil }
