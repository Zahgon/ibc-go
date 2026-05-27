package attestations

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
)

const (
	nanosPerSecond = 1_000_000_000
)

var (
	uint64Type, _  = abi.NewType("uint64", "", nil)
	bytes32Type, _ = abi.NewType("bytes32", "", nil)
	tupleArrayType = abi.Arguments{
		{Name: "path", Type: bytes32Type},
		{Name: "commitment", Type: bytes32Type},
	}

	stateAttestationArgs = abi.Arguments{
		{Name: "height", Type: uint64Type},
		{Name: "timestamp", Type: uint64Type},
	}

	packetAttestationType, _ = abi.NewType("tuple", "PacketAttestation", []abi.ArgumentMarshaling{
		{Name: "height", Type: "uint64"},
		{Name: "packets", Type: "tuple[]", Components: []abi.ArgumentMarshaling{
			{Name: "path", Type: "bytes32"},
			{Name: "commitment", Type: "bytes32"},
		}},
	})

	packetAttestationArgs = abi.Arguments{
		{Name: "attestation", Type: packetAttestationType},
	}
)

// ABIPacketCompact is the ABI-compatible representation with fixed-size arrays.
type ABIPacketCompact struct {
	Path       [32]byte
	Commitment [32]byte
}

// StateAttestation is used by client updates.
// This type uses ABI encoding (not Protobuf) for cross-platform compatibility.
type StateAttestation struct {
	Height    uint64
	Timestamp uint64
}

// PacketAttestation is used by membership queries.
// This type uses ABI encoding (not Protobuf) for cross-platform compatibility.
type PacketAttestation struct {
	Height  uint64
	Packets []PacketCompact
}

// PacketCompact represents a packet commitment.
// This type uses ABI encoding (not Protobuf) for cross-platform compatibility.
type PacketCompact struct {
	Path       []byte
	Commitment []byte
}

func (sa *StateAttestation) ABIEncode() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// ABIPacketAttestation is the ABI-compatible representation for tuple-wrapped encoding.
type ABIPacketAttestation struct {
	Height  uint64
	Packets []ABIPacketCompact
}

func (pa *PacketAttestation) ABIEncode() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Pack as tuple-wrapped struct to match Solidity's abi.encode(PacketAttestation)

func ABIDecodePacketAttestation(data []byte) (*PacketAttestation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Tuple-wrapped format: single element containing the struct

//nolint:revive // go-ethereum returns anonymous struct, cannot use named type

func (pc *PacketCompact) ABIEncode() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func ABIDecodeStateAttestation(data []byte) (*StateAttestation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func bytesToBytes32(b []byte) [32]byte { _ = "STUB: not implemented"; return nil }
