package types

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
)

// getICS20ABI returns an abi.Arguments slice describing the Solidity types of the struct.
func getICS20ABI() abi.Arguments {
	_ = "STUB: not implemented"
	// Create the ABI types for each field.
	// The Solidity types used are:
	// - string for Denom, Sender, Receiver and Memo.
	// - uint256 for Amount.
	return *new(abi.Arguments)
}

// Create an ABI argument representing our struct as a single tuple argument.

// DecodeABIFungibleTokenPacketData decodes a solidity ABI encoded ics20lib.ICS20LibFungibleTokenPacketData
// and converts it into an ibc-go FungibleTokenPacketData.
func DecodeABIFungibleTokenPacketData(data []byte) (*FungibleTokenPacketData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func EncodeABIFungibleTokenPacketData(data *FungibleTokenPacketData) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Pack the values in the order defined in the ABI.
