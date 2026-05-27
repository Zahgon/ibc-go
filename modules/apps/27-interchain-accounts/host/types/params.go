package types

const (
	// DefaultHostEnabled is the default value for the host param (set to true)
	DefaultHostEnabled = true
	// Maximum length of the allowlist
	MaxAllowListLength = 500
)

// NewParams creates a new parameter configuration for the host submodule
func NewParams(enableHost bool, allowMsgs []string) Params {
	_ = "STUB: not implemented"
	return *new(Params)
}

// DefaultParams is the default parameter configuration for the host submodule
func DefaultParams() Params { _ = "STUB: not implemented"; return *new(Params) }

// Validate validates all host submodule parameters
func (p Params) Validate() error { _ = "STUB: not implemented"; return nil }

func validateAllowlist(allowMsgs []string) error { _ = "STUB: not implemented"; return nil }
