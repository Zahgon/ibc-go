package attestations

import (
	"crypto/sha256"
)

// AttestationType distinguishes attestation types to prevent cross-protocol signature replay.
type AttestationType byte

const (
	// AttestationTypeState is used for client update (state) attestations.
	AttestationTypeState AttestationType = 0x01
	// AttestationTypePacket is used for packet membership/non-membership attestations.
	AttestationTypePacket AttestationType = 0x02
)

const (
	// SignatureLength is the expected length of an ECDSA signature (r||s||v)
	SignatureLength = 65
	// recoveryIDIndex is the byte position of the recovery ID (v) in the signature
	recoveryIDIndex = 64
	// domainSeparatedPreimageLen is the length of the domain-separated signing preimage:
	// 1-byte type tag + 32-byte SHA-256 hash.
	domainSeparatedPreimageLen = 1 + sha256.Size
)

// TaggedSigningInput computes the domain-separated prehash: `sha256(type_tag || sha256(data))`.
func TaggedSigningInput(data []byte, attestationType AttestationType) [32]byte {
	_ = "STUB: not implemented"
	return nil
}

// verifySignatures verifies that the attestation proof has valid signatures from unique attestors
// meeting the quorum threshold. Signatures cover `sha256(type_tag || sha256(attestationData))`.
func (cs *ClientState) verifySignatures(proof *AttestationProof, attestationType AttestationType) error {
	_ = "STUB: not implemented"
	return nil
}

// normalizeSignature converts the ECDSA recovery ID (v) from Ethereum format (27/28)
// to raw format (0/1). go-ethereum's crypto.SigToPub expects raw format, while
// Solidity's ECDSA.recover and most signing libraries produce Ethereum format.
func normalizeSignature(sig []byte) []byte { _ = "STUB: not implemented"; return nil }

// Already in raw format (0/1) or unknown, leave unchanged
