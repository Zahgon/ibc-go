package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

// RegisterInterfaces registers ibc types against interfaces using the global InterfaceRegistry.
// Note: The localhost client is created by ibc core and thus requires explicit type registration.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) { _ = "STUB: not implemented"; return }
