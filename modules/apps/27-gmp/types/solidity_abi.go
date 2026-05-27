package types

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
)

// getICS27PacketABI returns an abi.Arguments slice describing the Solidity types of the struct.
func getICS27PacketABI() abi.Arguments {
	_ = "STUB: not implemented"
	// Create the ABI types for each field.
	// The Solidity types used are:
	// - string for Sender, Receiver and Memo.
	// - bytes for Salt and Payload.
	return *new(abi.Arguments)
}

// Create an ABI argument representing our struct as a single tuple argument.

// getICS27AckABI returns an abi.Arguments slice describing the Solidity types of the struct.
func getICS27AckABI() abi.Arguments {
	_ = "STUB: not implemented"
	// Create the ABI types for each field.
	// The Solidity types used are:
	// - bytes for Result.
	return *new(abi.Arguments)
}

// Create an ABI argument representing our struct as a single tuple argument.

// DecodeABIGMPPacketData decodes a solidity ABI encoded ics27lib.GMPPacketData and converts it into an ibc-go GMPPacketData.
func DecodeABIGMPPacketData(data []byte) (*GMPPacketData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EncodeABIGMPPacketData encodes a GMPPacketData into a solidity ABI encoded byte array.
func EncodeABIGMPPacketData(data *GMPPacketData) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Pack the values in the order defined in the ABI.

// DecodeABIAcknowledgement decodes a solidity ABI encoded ics27lib.Acknowledgement and converts it into an ibc-go Acknowledgement
func DecodeABIAcknowledgement(data []byte) (*Acknowledgement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EncodeABIAcknowledgement encodes an Acknowledgement into a solidity ABI encoded byte array
func EncodeABIAcknowledgement(data *Acknowledgement) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Pack the values in the order defined in the ABI.
