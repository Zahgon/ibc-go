package types

// Maximum length of the allowed clients list
const MaxAllowedClientsLength = 200

// DefaultAllowedClients are the default clients for the AllowedClients parameter.
// By default it allows all client types.
var DefaultAllowedClients = []string{AllowAllClients}

// NewParams creates a new parameter configuration for the ibc client module
func NewParams(allowedClients ...string) Params { _ = "STUB: not implemented"; return *new(Params) }

// DefaultParams is the default parameter configuration for the ibc-client module.
func DefaultParams() Params { _ = "STUB: not implemented"; return *new(Params) }

// Validate all ibc-client module parameters
func (p Params) Validate() error { _ = "STUB: not implemented"; return nil }

// IsAllowedClient checks if the given client type is registered on the allowlist.
func (p Params) IsAllowedClient(clientType string) bool {
	_ = "STUB: not implemented"
	// Still need to check for blank client type
	return false
}

// Check for allow all client wildcard
// If exist then allow all type of client

// validateClients checks that the given clients are not blank and there are no duplicates.
// If AllowAllClients wildcard (*) is used, then there should no other client types in the allow list
func validateClients(clients []string) error { _ = "STUB: not implemented"; return nil }
