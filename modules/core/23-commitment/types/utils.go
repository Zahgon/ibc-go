package types

import (
	"github.com/cometbft/cometbft/proto/tendermint/crypto"
)

// ConvertProofs converts crypto.ProofOps into MerkleProof
func ConvertProofs(tmProof *crypto.ProofOps) (MerkleProof, error) {
	_ = "STUB: not implemented"
	return *new(MerkleProof), nil
}

// Unmarshal all proof ops to CommitmentProof
