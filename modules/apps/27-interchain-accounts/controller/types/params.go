package types

const (
	// DefaultControllerEnabled is the default value for the controller param (set to true)
	DefaultControllerEnabled = true
)

// NewParams creates a new parameter configuration for the controller submodule
func NewParams(enableController bool) Params { _ = "STUB: not implemented"; return *new(Params) }

// DefaultParams is the default parameter configuration for the controller submodule
func DefaultParams() Params { _ = "STUB: not implemented"; return *new(Params) }
