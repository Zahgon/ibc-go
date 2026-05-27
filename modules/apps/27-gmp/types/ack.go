package types

// NewAcknowledgement creates a new Acknowledgement
func NewAcknowledgement(result []byte) Acknowledgement {
	_ = "STUB: not implemented"
	return *new(Acknowledgement)
}

// ValidateBasic performs basic validation on the Acknowledgement
func (Acknowledgement) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

func MarshalAcknowledgement(data *Acknowledgement, ics27Version string, encoding string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func UnmarshalAcknowledgement(bz []byte, ics27Version string, encoding string) (*Acknowledgement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
