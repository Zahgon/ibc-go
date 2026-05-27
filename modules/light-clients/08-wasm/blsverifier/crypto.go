package blsverifier

import (
	"github.com/OffchainLabs/prysm/v6/crypto/bls"
)

func AggregatePublicKeys(publicKeys [][]byte) (bls.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(bls.PublicKey), nil
}

func VerifySignature(signature []byte, message [32]byte, publicKeys [][]byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
