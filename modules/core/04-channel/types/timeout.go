package types

import (
	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
)

// NewTimeout returns a new Timeout instance.
func NewTimeout(height clienttypes.Height, timestamp uint64) Timeout {
	_ = "STUB: not implemented"
	return *new(Timeout)
}

// NewTimeoutWithTimestamp creates a new Timeout with only the timestamp set.
func NewTimeoutWithTimestamp(timestamp uint64) Timeout {
	_ = "STUB: not implemented"
	return *new(Timeout)
}

// IsValid returns true if either the height or timestamp is non-zero.
func (t Timeout) IsValid() bool { _ = "STUB: not implemented"; return false }

// Elapsed returns true if either the provided height or timestamp is past the
// respective absolute timeout values.
func (t Timeout) Elapsed(height clienttypes.Height, timestamp uint64) bool {
	_ = "STUB: not implemented"
	return false
}

// TimestampElapsed returns true if the provided timestamp is past the timeout timestamp.
func (t Timeout) TimestampElapsed(timestamp uint64) bool { _ = "STUB: not implemented"; return false }

// ErrTimeoutElapsed returns a timeout elapsed error indicating which timeout value
// has elapsed.
func (t Timeout) ErrTimeoutElapsed(height clienttypes.Height, timestamp uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// ErrTimeoutNotReached returns a timeout not reached error indicating which timeout value
// has not been reached.
func (t Timeout) ErrTimeoutNotReached(height clienttypes.Height, timestamp uint64) error {
	_ = "STUB: not implemented"
	// only return height information if the height is set
	// t.heightElapsed() will return false when it is empty
	return nil
}

// heightElapsed returns true if the timeout height is non empty
// and the timeout height is greater than or equal to the relative height.
func (t Timeout) heightElapsed(height clienttypes.Height) bool {
	_ = "STUB: not implemented"
	return false
}

// timestampElapsed returns true if the timeout timestamp is non empty
// and the timeout timestamp is greater than or equal to the relative timestamp.
func (t Timeout) timestampElapsed(timestamp uint64) bool { _ = "STUB: not implemented"; return false }
