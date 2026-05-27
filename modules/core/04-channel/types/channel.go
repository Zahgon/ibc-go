package types

// NewChannel creates a new Channel instance
func NewChannel(
	state State, ordering Order, counterparty Counterparty,
	hops []string, version string,
) Channel {
	_ = "STUB: not implemented"
	return *new(Channel)
}

// ValidateBasic performs a basic validation of the channel fields
func (ch Channel) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewCounterparty returns a new Counterparty instance
func NewCounterparty(portID, channelID string) Counterparty {
	_ = "STUB: not implemented"
	return *new(Counterparty)
}

// ValidateBasic performs a basic validation check of the identifiers
func (c Counterparty) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// NewIdentifiedChannel creates a new IdentifiedChannel instance
func NewIdentifiedChannel(portID, channelID string, ch Channel) IdentifiedChannel {
	_ = "STUB: not implemented"
	return *new(IdentifiedChannel)
}

// ValidateBasic performs a basic validation of the identifiers and channel fields.
func (ic IdentifiedChannel) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
