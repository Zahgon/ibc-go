package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewGenesisState creates a new ibc-transfer GenesisState instance.
func NewGenesisState(portID string, denoms Denoms, params Params, totalEscrowed sdk.Coins) *GenesisState {
	_ = "STUB: not implemented"
	return nil
}

// DefaultGenesisState returns a GenesisState with "transfer" as the default PortID.
func DefaultGenesisState() *GenesisState { _ = "STUB: not implemented"; return nil }

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error { _ = "STUB: not implemented"; return nil }

// will fail if there are duplicates for any denom
