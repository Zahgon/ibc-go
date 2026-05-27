package types

import (
	"time"
)

// DefaultTimePerBlock is the default value for maximum expected time per block (in nanoseconds).
const DefaultTimePerBlock = 30 * time.Second

// NewParams creates a new parameter configuration for the ibc connection module
func NewParams(timePerBlock uint64) Params { _ = "STUB: not implemented"; return *new(Params) }

// DefaultParams is the default parameter configuration for the ibc connection module
func DefaultParams() Params { _ = "STUB: not implemented"; return *new(Params) }

// Validate ensures MaxExpectedTimePerBlock is non-zero
func (p Params) Validate() error { _ = "STUB: not implemented"; return nil }
