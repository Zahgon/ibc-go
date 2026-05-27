package solomachine

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
)

// VerifySignature verifies if the provided public key generated the signature
// over the given data. Single and Multi signature public keys are supported.
// The signature data type must correspond to the public key type. An error is
// returned if signature verification fails or an invalid SignatureData type is
// provided.
func VerifySignature(pubKey cryptotypes.PubKey, signBytes []byte, sigData signing.SignatureData) error {
	_ = "STUB: not implemented"
	return nil
}

// The function supplied fulfills the VerifyMultisignature interface. No special
// adjustments need to be made to the sign bytes based on the sign mode.
