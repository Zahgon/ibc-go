package types

import (
	"crypto/sha256"

	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var (
	ErrorAcknowledgement                          = sha256.Sum256([]byte("UNIVERSAL_ERROR_ACKNOWLEDGEMENT"))
	_                    exported.Acknowledgement = &Acknowledgement{}
)

// NewAcknowledgement creates a new Acknowledgement containing the provided app acknowledgements.
func NewAcknowledgement(appAcknowledgements ...[]byte) Acknowledgement {
	_ = "STUB: not implemented"
	return *new(Acknowledgement)
}

// Validate performs a basic validation of the acknowledgement
func (ack Acknowledgement) Validate() error {
	_ = "STUB: not implemented"
	// acknowledgement list should be non-empty
	return nil
}

// Each app acknowledgement should be non-empty

// Ensure that the app acknowledgement contains ErrorAcknowledgement
// **if and only if** the app acknowledgement list has a single element

// Success returns true if the acknowledgement is successful
// it implements the exported.Acknowledgement interface
func (ack Acknowledgement) Success() bool { _ = "STUB: not implemented"; return false }

// Acknowledgement returns the acknowledgement bytes to implement the acknowledgement interface
func (ack Acknowledgement) Acknowledgement() []byte { _ = "STUB: not implemented"; return nil }
