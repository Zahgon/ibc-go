package attestations

import (
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var _ exported.ClientMessage = (*AttestationProof)(nil)

// ClientType defines that the AttestationProof is for Attestations.
func (AttestationProof) ClientType() string { _ = "STUB: not implemented"; return "" }

// ValidateBasic ensures that the attestation data and signatures are initialized.
// Attestation data can be either a StateAttestation (for client updates) or
// a PacketAttestation (for packet membership/non-membership proofs).
func (ap AttestationProof) ValidateBasic() error { _ = "STUB: not implemented"; return nil }

// Try to decode as PacketAttestation first (used for membership/non-membership proofs)

// If that fails, try to decode as StateAttestation (used for client updates)
