package types

const (
	// DefaultSendEnabled enabled
	DefaultSendEnabled = true
	// DefaultReceiveEnabled enabled
	DefaultReceiveEnabled = true
)

// NewParams creates a new parameter configuration for the ibc transfer module
func NewParams(enableSend, enableReceive bool) Params {
	_ = "STUB: not implemented"
	return *new(Params)
}

// DefaultParams is the default parameter configuration for the ibc-transfer module
func DefaultParams() Params { _ = "STUB: not implemented"; return *new(Params) }
