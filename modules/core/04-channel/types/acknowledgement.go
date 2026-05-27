package types

const (
	// ackErrorString defines a string constant included in error acknowledgements
	// NOTE: Changing this const is state machine breaking as acknowledgements are written into state.
	ackErrorString = "error handling packet: see events for details"
)

// NewResultAcknowledgement returns a new instance of Acknowledgement using an Acknowledgement_Result
// type in the Response field.
func NewResultAcknowledgement(result []byte) Acknowledgement {
	_ = "STUB: not implemented"
	return *new(Acknowledgement)
}

// NewErrorAcknowledgementWithCodespace returns a new instance of Acknowledgement using an Acknowledgement_Error
// type in the Response field.
// NOTE: The error includes the ABCI codespace and code in the error string to provide more information about the module
// that generated the error. This is useful for debugging but can potentially introduce non-determinism if care is
// not taken to ensure the codespace doesn't change in non state-machine breaking versions.
func NewErrorAcknowledgementWithCodespace(err error) Acknowledgement {
	_ = "STUB: not implemented"
	// The ABCI code is included in the abcitypes.ResponseDeliverTx hash
	// constructed in Tendermint and is therefore deterministic.
	// However, a code without codespace is incomplete information (e.g. sdk/5 and wasm/5 are
	// different errors). We add this codespace here, in order to provide a meaningful error
	// identifier which means changing the codespace of an error becomes a consensus breaking change.
	return *new(Acknowledgement)
}

// NewErrorAcknowledgement returns a new instance of Acknowledgement using an Acknowledgement_Error
// type in the Response field.
// NOTE: Acknowledgements are written into state and thus, changes made to error strings included in packet acknowledgements
// risk an app hash divergence when nodes in a network are running different patch versions of software.
func NewErrorAcknowledgement(err error) Acknowledgement {
	_ = "STUB: not implemented"
	// the ABCI code is included in the abcitypes.ResponseDeliverTx hash
	// constructed in Tendermint and is therefore deterministic
	return *new(Acknowledgement)
}

// discard non-deterministic codespace and log values

// ValidateBasic performs a basic validation of the acknowledgement
func (ack Acknowledgement) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// Success implements the Acknowledgement interface. The acknowledgement is
// considered successful if it is a ResultAcknowledgement. Otherwise it is
// considered a failed acknowledgement.
func (ack Acknowledgement) Success() bool { _ = "STUB: not implemented"; return false }

// Acknowledgement implements the Acknowledgement interface. It returns the
// acknowledgement serialised using JSON.
func (ack Acknowledgement) Acknowledgement() []byte { _ = "STUB: not implemented"; return nil }
