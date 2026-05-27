package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v11/modules/core/05-port/types"
)

/*

ADR-8 implementation

The Memo is used to ensure that the callback is desired by the user. This allows a user to send a packet to an ADR-8 enabled contract.

The Memo format is defined like so:

```json
{
	// ... other memo fields we don't care about
	"src_callback": {
		"address": {stringCallbackAddress},

		// optional fields
		"gas_limit": {stringForCallback}
	},
	"dest_callback": {
		"address": {stringCallbackAddress},

		// optional fields
		"gas_limit": {stringForCallback}
	}
}
```

We will pass the packet sender info (if available) to the contract keeper for source callback executions. This will allow the contract
keeper to verify that the packet sender is the same as the callback address if desired.

*/

// CallbackData is the callback data parsed from the packet.
type CallbackData struct {
	// CallbackAddress is the address of the callback actor.
	CallbackAddress string
	// ExecutionGasLimit is the gas limit which will be used for the callback execution.
	ExecutionGasLimit uint64
	// SenderAddress is the sender of the packet. This is passed to the contract keeper
	// to verify that the packet sender is the same as the callback address if desired.
	// This address is empty during destination callback execution.
	// This address may be empty if the sender is unknown or undefined.
	SenderAddress string
	// CommitGasLimit is the gas needed to commit the callback even if the callback
	// execution fails due to out of gas.
	// This parameter is only used in event emissions, or logging.
	CommitGasLimit uint64
	// ApplicationVersion is the base application version.
	ApplicationVersion string
	// Calldata is the calldata to be passed to the callback actor.
	// This may be empty but if it is not empty, it should be the calldata sent to the callback actor.
	Calldata []byte
}

// GetSourceCallbackData parses the packet data and returns the source callback data.
func GetSourceCallbackData(
	ctx sdk.Context,
	packetDataUnmarshaler porttypes.PacketDataUnmarshaler,
	packet channeltypes.Packet,
	maxGas uint64,
) (CallbackData, bool, error) {
	_ = "STUB: not implemented"
	return *new(CallbackData), false, nil
}

// GetDestCallbackData parses the packet data and returns the destination callback data.
func GetDestCallbackData(
	ctx sdk.Context,
	packetDataUnmarshaler porttypes.PacketDataUnmarshaler,
	packet channeltypes.Packet, maxGas uint64,
) (CallbackData, bool, error) {
	_ = "STUB: not implemented"
	return *new(CallbackData), false, nil
}

// GetCallbackData parses the packet data and returns the callback data.
// If the packet data does not have callback data set, it will return false for `isCbPacket` and an error.
// If the packet data does have callback data set but the callback data is malformed,
// it will return true for `isCbPacket` and an error
// It also checks that the remaining gas is greater than the gas limit specified in the packet data.
// The addressGetter and gasLimitGetter functions are used to retrieve the callback
// address and gas limit from the callback data.
func GetCallbackData(
	packetData any,
	version, srcPortID string,
	remainingGas, maxGas uint64,
	callbackKey string,
) (CallbackData, bool, error) {
	_ = "STUB: not implemented"
	return *new(CallbackData), false, nil
}

// get the callback address from the callback data

// retrieve packet sender from packet data if possible and if needed

// get the gas limit from the callback data

func computeExecAndCommitGasLimit(callbackData map[string]any, remainingGas, maxGas uint64) (uint64, uint64, error) {
	_ = "STUB: not implemented"
	// get the gas limit from the callback data
	return 0, 0, nil
}

// ensure user defined gas limit does not exceed the max gas limit

// account for the remaining gas in the context being less than the desired gas limit for the callback execution
// in this case, the callback execution may be retried upon failure

// getUserDefinedGasLimit returns the custom gas limit provided for callbacks if it is
// in the callback data. It is assumed that callback data is not nil.
// If no gas limit is specified or the gas limit is improperly formatted, 0 is returned.
//
// The memo is expected to specify the user defined gas limit in the following format:
// { "{callbackKey}": { ... , "gas_limit": {stringForCallback} }
//
// Note: the user defined gas limit must be set as a string and not a json number.
func getUserDefinedGasLimit(callbackData map[string]any) (uint64, error) {
	_ = "STUB: not implemented"
	// the gas limit must be specified as a string and not a json number
	return 0, nil
}

// getCallbackAddress returns the callback address if it is specified in the callback data.
// It is assumed that callback data is not nil.
// If no callback address is specified or the memo is improperly formatted, an empty string is returned.
//
// The memo is expected to contain the callback address in the following format:
// { "{callbackKey}": { "address": {stringCallbackAddress}}
//
// ADR-8 middleware should callback on the returned address if it is a PacketActor
// (i.e. smart contract that accepts IBC callbacks).
func getCallbackAddress(callbackData map[string]any) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// getCalldata returns the calldata if it is specified in the callback data.
func getCalldata(callbackData map[string]any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AllowRetry returns true if the callback execution gas limit is less than the commit gas limit.
func (c CallbackData) AllowRetry() bool { _ = "STUB: not implemented"; return false }
