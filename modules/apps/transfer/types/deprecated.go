package types

import (
	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Deprecated: usage of this function should be replaced by `Denom.hasPrefix`
// SenderChainIsSource returns false if the denomination originally came
// from the receiving chain and true otherwise.
func SenderChainIsSource(sourcePort, sourceChannel, denom string) bool {
	_ = "STUB: not implemented"
	// This is the prefix that would have been prefixed to the denomination
	// on sender chain IF and only if the token originally came from the
	// receiving chain.
	return false
}

// Deprecated: usage of this function should be replaced by `Denom.hasPrefix`
// ReceiverChainIsSource returns true if the denomination originally came
// from the receiving chain and false otherwise.
func ReceiverChainIsSource(sourcePort, sourceChannel, denom string) bool {
	_ = "STUB: not implemented"
	// The prefix passed in should contain the SourcePort and SourceChannel.
	// If  the receiver chain originally sent the token to the sender chain
	// the denom will have the sender's SourcePort and SourceChannel as the
	// prefix.
	return false
}

// Deprecated: usage of this function should be replaced by `NewHop`
// GetDenomPrefix returns the receiving denomination prefix
func GetDenomPrefix(portID, channelID string) string { _ = "STUB: not implemented"; return "" }

// Deprecated: usage of this function should be replaced by `NewDenom`
// GetPrefixedDenom returns the denomination with the portID and channelID prefixed
func GetPrefixedDenom(portID, channelID, baseDenom string) string {
	_ = "STUB: not implemented"
	return ""
}

// Deprecated: usage of this function should be replaced by `Token.ToCoin`
// GetTransferCoin creates a transfer coin with the port ID and channel ID
// prefixed to the base denom.
func GetTransferCoin(portID, channelID, baseDenom string, amount sdkmath.Int) sdk.Coin {
	_ = "STUB: not implemented"
	return *new(sdk.Coin)
}

// Deprecated: usage of this function should be replaced by `ExtractDenomFromPath`
// ExtractDenomFromPath returns the denom from the full path.
func ParseDenomTrace(rawDenom string) Denom { _ = "STUB: not implemented"; return *new(Denom) }
