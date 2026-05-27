package types

// UnmarshalJSON implements the Unmarshaller interface for FungibleTokenPacketData.
func (ftpd *FungibleTokenPacketData) UnmarshalJSON(bz []byte) error {
	_ = "STUB: not implemented"
	// Recursion protection. We cannot unmarshal into FungibleTokenPacketData directly
	// else UnmarshalJSON is going to get invoked again, ad infinum. Create an alias
	// and unmarshal into that, instead.
	return nil
}

// Raise errors during decoding if unknown fields are encountered.
