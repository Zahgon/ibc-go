package types

// NewHop creates a Hop with the given port ID and channel ID.
func NewHop(portID, channelID string) Hop { _ = "STUB: not implemented"; return *new(Hop) }

// Validate performs a basic validation of the Hop fields.
func (h Hop) Validate() error { _ = "STUB: not implemented"; return nil }

// String returns the Hop in the format:
// <portID>/<channelID>
func (h Hop) String() string { _ = "STUB: not implemented"; return "" }
