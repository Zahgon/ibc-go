package mock

// EmptyAcknowledgement implements the exported.Acknowledgement interface and always returns an empty byte string as Response
type EmptyAcknowledgement struct {
	Response []byte
}

// NewEmptyAcknowledgement returns a new instance of EmptyAcknowledgement
func NewEmptyAcknowledgement() EmptyAcknowledgement {
	_ = "STUB: not implemented"
	return *new(EmptyAcknowledgement)
}

// Success implements the Acknowledgement interface
func (EmptyAcknowledgement) Success() bool {
	_ = "STUB: not implemented"

	// Acknowledgement implements the Acknowledgement interface
	return false
}

func (EmptyAcknowledgement) Acknowledgement() []byte { _ = "STUB: not implemented"; return nil }
