package types

import (
	ics23 "github.com/cosmos/ics23/go"

	v2 "github.com/cosmos/ibc-go/v11/modules/core/23-commitment/types/v2"
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

// var representing the proofspecs for an SDK chain
var sdkSpecs = []*ics23.ProofSpec{ics23.IavlSpec, ics23.TendermintSpec}

// ICS 023 Merkle Types Implementation
//
// This file defines Merkle commitment types that implements ICS 023.

// Merkle proof implementation of the Proof interface
// Applied on SDK-based IBC implementation
var _ exported.Root = (*MerkleRoot)(nil)

// GetSDKSpecs is a getter function for the proofspecs of an sdk chain
func GetSDKSpecs() []*ics23.ProofSpec {
	_ = "STUB: not implemented"

	// NewMerkleRoot constructs a new MerkleRoot
	return nil
}

func NewMerkleRoot(hash []byte) MerkleRoot { _ = "STUB: not implemented"; return *new(MerkleRoot) }

// GetHash implements RootI interface
func (mr MerkleRoot) GetHash() []byte {
	_ = "STUB: not implemented"

	// Empty returns true if the root is empty
	return nil
}

func (mr MerkleRoot) Empty() bool { _ = "STUB: not implemented"; return false }

var _ exported.Prefix = (*MerklePrefix)(nil)

// NewMerklePrefix constructs new MerklePrefix instance
func NewMerklePrefix(keyPrefix []byte) MerklePrefix {
	_ = "STUB: not implemented"
	return *new(MerklePrefix)
}

// Bytes returns the key prefix bytes
func (mp MerklePrefix) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// Empty returns true if the prefix is empty
func (mp MerklePrefix) Empty() bool { _ = "STUB: not implemented"; return false }

// NewMerklePath creates a new MerklePath instance
// The keys must be passed in from root-to-leaf order.
// NOTE: NewMerklePath returns a commitment/v2 MerklePath.
var NewMerklePath = v2.NewMerklePath

// ApplyPrefix constructs a new commitment path from the arguments. It prepends the prefix key
// with the given path.
func ApplyPrefix(prefix exported.Prefix, path v2.MerklePath) (v2.MerklePath, error) {
	_ = "STUB: not implemented"
	return *new(v2.MerklePath), nil
}

// VerifyMembership verifies the membership of a merkle proof against the given root, path, and value.
// Note that the path is expected as []string{<store key of module>, <key corresponding to requested value>}.
func (p MerkleProof) VerifyMembership(specs []*ics23.ProofSpec, root exported.Root, path exported.Path, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyMembership specific argument validation

// Since every proof in chain is a membership proof we can use verifyChainedMembershipProof from index 0
// to validate entire proof

// VerifyNonMembership verifies the absence of a merkle proof against the given root and path.
// VerifyNonMembership verifies a chained proof where the absence of a given path is proven
// at the lowest subtree and then each subtree's inclusion is proved up to the final root.
func (p MerkleProof) VerifyNonMembership(specs []*ics23.ProofSpec, root exported.Root, path exported.Path) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyNonMembership will verify the absence of key in lowest subtree, and then chain inclusion proofs
// of all subroots up to final root

// Verify chained membership proof starting from index 1 with value = subroot

// verifyChainedMembershipProof takes a list of proofs and specs and verifies each proof sequentially ensuring that the value is committed to
// by first proof and each subsequent subroot is committed to by the next subroot and checking that the final calculated root is equal to the given roothash.
// The proofs and specs are passed in from lowest subtree to the highest subtree, but the keys are passed in from highest subtree to lowest.
// The index specifies what index to start chaining the membership proofs, this is useful since the lowest proof may not be a membership proof, thus we
// will want to start the membership proof chaining from index 1 with value being the lowest subroot
func verifyChainedMembershipProof(root []byte, specs []*ics23.ProofSpec, proofs []*ics23.CommitmentProof, keys v2.MerklePath, value []byte, index int) error {
	_ = "STUB: not implemented"
	return nil
}

// Initialize subroot to value since the proofs list may be empty.
// This may happen if this call is verifying intermediate proofs after the lowest proof has been executed.
// In this case, there may be no intermediate proofs to verify and we just check that lowest proof root equals final root

// Since keys are passed in from highest to lowest, we must grab their indices in reverse order
// from the proofs and specs which are lowest to highest

// verify membership of the proof at this index with appropriate key and value

// Set value to subroot so that we verify next proof in chain commits to this subroot

// Check that chained proof root equals passed-in root

// validateVerificationArgs verifies the proof arguments are valid.
// The merkle path and merkle proof contain a list of keys and their proofs
// which correspond to individual trees. The length of these keys and their proofs
// must equal the length of the given specs. All arguments must be non-empty.
func validateVerificationArgs(proof MerkleProof, path v2.MerklePath, specs []*ics23.ProofSpec, root exported.Root) error {
	_ = "STUB: not implemented"
	return nil
}
