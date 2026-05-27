package simapp

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/v2/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	callbacktypes "github.com/cosmos/ibc-go/v11/modules/apps/callbacks/types"
	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
	ibcexported "github.com/cosmos/ibc-go/v11/modules/core/exported"
)

// MockKeeper implements callbacktypes.ContractKeeper
var _ callbacktypes.ContractKeeper = (*ContractKeeper)(nil)

var StatefulCounterKey = "stateful-callback-counter"

const (
	// OogPanicContract is a contract address that will panic out of gas
	OogPanicContract = "panics out of gas"
	// OogErrorContract is a contract address that will error out of gas
	OogErrorContract = "errors out of gas"
	// PanicContract is a contract address that will panic
	PanicContract = "panics"
	// ErrorContract is a contract address that will return an error
	ErrorContract = "errors"
	// SuccessContract is a contract address that will return nil
	SuccessContract = "success"
)

// This is a mock contract keeper used for testing. It is not wired up to any modules.
// It implements the interface functions expected by the ibccallbacks middleware
// so that it can be tested with simapp. The keeper is responsible for tracking
// two metrics:
//   - number of callbacks called per callback type
//   - stateful entry attempts
//
// The counter for callbacks allows us to ensure the correct callbacks were routed to
// and the stateful entries allows us to track state reversals or reverted state upon
// contract execution failure or out of gas errors.
type ContractKeeper struct {
	key storetypes.StoreKey

	Counters map[callbacktypes.CallbackType]int

	IBCSendPacketCallbackFn func(
		cachedCtx sdk.Context,
		sourcePort string,
		sourceChannel string,
		timeoutHeight clienttypes.Height,
		timeoutTimestamp uint64,
		packetData []byte,
		contractAddress,
		packetSenderAddress string,
		version string,
	) error

	IBCOnAcknowledgementPacketCallbackFn func(
		cachedCtx sdk.Context,
		packet channeltypes.Packet,
		acknowledgement []byte,
		relayer sdk.AccAddress,
		contractAddress,
		packetSenderAddress string,
		version string,
	) error

	IBCOnTimeoutPacketCallbackFn func(
		cachedCtx sdk.Context,
		packet channeltypes.Packet,
		relayer sdk.AccAddress,
		contractAddress,
		packetSenderAddress string,
		version string,
	) error

	IBCReceivePacketCallbackFn func(
		cachedCtx sdk.Context,
		packet ibcexported.PacketI,
		ack ibcexported.Acknowledgement,
		contractAddress string,
		version string,
	) error
}

// SetStateEntryCounter sets state entry counter. The number of stateful
// entries is tracked as a uint8. This function is used to test state reversals.
func (k ContractKeeper) SetStateEntryCounter(ctx sdk.Context, count uint8) {
	_ = "STUB: not implemented"
	return
}

// GetStateEntryCounter returns the state entry counter stored in state.
func (k ContractKeeper) GetStateEntryCounter(ctx sdk.Context) uint8 {
	_ = "STUB: not implemented"
	return 0
}

// IncrementStateEntryCounter increments the stateful callback counter in state.
func (k ContractKeeper) IncrementStateEntryCounter(ctx sdk.Context) {
	_ = "STUB: not implemented"
	return
}

// NewContractKeeper creates a new mock ContractKeeper.
func NewContractKeeper(key storetypes.StoreKey) *ContractKeeper {
	_ = "STUB: not implemented"
	return nil
}

// IBCSendPacketCallback increments the stateful entry counter and the send_packet callback counter.
// This function:
//   - returns MockApplicationCallbackError and consumes half the remaining gas if the contract address is ErrorContract
//   - Oog panics and consumes all the remaining gas + 1 if the contract address is OogPanicContract
//   - returns MockApplicationCallbackError and consumes all the remaining gas + 1 if the contract address is OogErrorContract
//   - Panics and consumes half the remaining gas if the contract address is PanicContract
//   - returns nil and consumes half the remaining gas if the contract address is SuccessContract or any other value
func (k ContractKeeper) IBCSendPacketCallback(
	ctx sdk.Context,
	sourcePort string,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
	packetData []byte,
	contractAddress,
	packetSenderAddress,
	version string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// IBCOnAcknowledgementPacketCallback increments the stateful entry counter and the acknowledgement_packet callback counter.
// This function:
//   - returns MockApplicationCallbackError and consumes half the remaining gas if the contract address is ErrorContract
//   - Oog panics and consumes all the remaining gas + 1 if the contract address is OogPanicContract
//   - returns MockApplicationCallbackError and consumes all the remaining gas + 1 if the contract address is OogErrorContract
//   - Panics and consumes half the remaining gas if the contract address is PanicContract
//   - returns nil and consumes half the remaining gas if the contract address is SuccessContract or any other value
func (k ContractKeeper) IBCOnAcknowledgementPacketCallback(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
	contractAddress,
	packetSenderAddress,
	version string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// IBCOnTimeoutPacketCallback increments the stateful entry counter and the timeout_packet callback counter.
// This function:
//   - returns MockApplicationCallbackError and consumes half the remaining gas if the contract address is ErrorContract
//   - Oog panics and consumes all the remaining gas + 1 if the contract address is OogPanicContract
//   - returns MockApplicationCallbackError and consumes all the remaining gas + 1 if the contract address is OogErrorContract
//   - Panics and consumes half the remaining gas if the contract address is PanicContract
//   - returns nil and consumes half the remaining gas if the contract address is SuccessContract or any other value
func (k ContractKeeper) IBCOnTimeoutPacketCallback(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
	contractAddress,
	packetSenderAddress,
	version string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// IBCReceivePacketCallback increments the stateful entry counter and the receive_packet callback counter.
// This function:
//   - returns MockApplicationCallbackError and consumes half the remaining gas if the contract address is ErrorContract
//   - Oog panics and consumes all the remaining gas + 1 if the contract address is OogPanicContract
//   - returns MockApplicationCallbackError and consumes all the remaining gas + 1 if the contract address is OogErrorContract
//   - Panics and consumes half the remaining gas if the contract address is PanicContract
//   - returns nil and consumes half the remaining gas if the contract address is SuccessContract or any other value
func (k ContractKeeper) IBCReceivePacketCallback(
	ctx sdk.Context,
	packet ibcexported.PacketI,
	ack ibcexported.Acknowledgement,
	contractAddress,
	version string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// ProcessMockCallback processes a mock callback.
// It increments the stateful entry counter and the callback counter.
// This function:
//   - returns MockApplicationCallbackError and consumes half the remaining gas if the contract address is ErrorContract
//   - Oog panics and consumes all the remaining gas + 1 if the contract address is OogPanicContract
//   - returns MockApplicationCallbackError and consumes all the remaining gas + 1 if the contract address is OogErrorContract
//   - Panics and consumes half the remaining gas if the contract address is PanicContract
//   - returns nil and consumes half the remaining gas if the contract address is SuccessContract or any other value
func (k ContractKeeper) ProcessMockCallback(
	ctx sdk.Context,
	callbackType callbacktypes.CallbackType,
	contractAddress string,
) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// increment stateful entries, if the callbacks module handler
// reverts state, we can check by querying for the counter
// currently stored.

// increment callback execution attempts

// consume half of the remaining gas so that ConsumeGas cannot oog panic

// unreachable

// unreachable

// consume half of the remaining gas so that ConsumeGas cannot oog panic

// consume half of the remaining gas so that ConsumeGas cannot oog panic

// success
