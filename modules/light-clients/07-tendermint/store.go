package tendermint

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/v2/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

/*
This file contains the logic for storage and iteration over `IterationKey` metadata that is stored
for each consensus state. The consensus state key specified in ICS-24 and expected by counterparty chains
stores the consensus state under the key: `consensusStates/{revision_number}-{revision_height}`, with each number
represented as a string.
While this works fine for IBC proof verification, it makes efficient iteration difficult since the lexicographic order
of the consensus state keys do not match the height order of consensus states. This makes consensus state pruning and
monotonic time enforcement difficult since it is inefficient to find the earliest consensus state or to find the neighboring
consensus states given a consensus state height.
Changing the ICS-24 representation will be a major breaking change that requires counterparty chains to accept a new key format.
Thus to avoid breaking IBC, we can store a lookup from a more efficiently formatted key: `iterationKey` to the consensus state key which
stores the underlying consensus state. This efficient iteration key will be formatted like so: `iterateConsensusStates{BigEndianRevisionBytes}{BigEndianHeightBytes}`.
This ensures that the lexicographic order of iteration keys match the height order of the consensus states. Thus, we can use the SDK store's
Iterators to iterate over the consensus states in ascending/descending order by providing a mapping from `iterationKey -> consensusStateKey -> ConsensusState`.
A future version of IBC may choose to replace the ICS24 ConsensusState path with the more efficient format and make this indirection unnecessary.
*/

const KeyIterateConsensusStatePrefix = "iterateConsensusStates"

var (
	// KeyProcessedTime is appended to consensus state key to store the processed time
	KeyProcessedTime = []byte("/processedTime")
	// KeyProcessedHeight is appended to consensus state key to store the processed height
	KeyProcessedHeight = []byte("/processedHeight")
	// KeyIteration stores the key mapping to consensus state key for efficient iteration
	KeyIteration = []byte("/iterationKey")
)

// setClientState stores the client state
func setClientState(clientStore storetypes.KVStore, cdc codec.BinaryCodec, clientState *ClientState) {
	_ = "STUB: not implemented"
	return
}

// getClientState retrieves the client state from the store using the provided KVStore and codec.
// It returns the unmarshaled ClientState and a boolean indicating if the state was found.
func getClientState(store storetypes.KVStore, cdc codec.BinaryCodec) (*ClientState, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// setConsensusState stores the consensus state at the given height.
func setConsensusState(clientStore storetypes.KVStore, cdc codec.BinaryCodec, consensusState *ConsensusState, height exported.Height) {
	_ = "STUB: not implemented"
	return
}

// GetConsensusState retrieves the consensus state from the client prefixed store.
// If the ConsensusState does not exist in state for the provided height a nil value and false boolean flag is returned
func GetConsensusState(store storetypes.KVStore, cdc codec.BinaryCodec, height exported.Height) (*ConsensusState, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// deleteConsensusState deletes the consensus state at the given height
func deleteConsensusState(clientStore storetypes.KVStore, height exported.Height) {
	_ = "STUB: not implemented"
	return
}

// ProcessedTimeKey returns the key under which the processed time will be stored in the client store.
func ProcessedTimeKey(height exported.Height) []byte { _ = "STUB: not implemented"; return nil }

// SetProcessedTime stores the time at which a header was processed and the corresponding consensus state was created.
// This is useful when validating whether a packet has reached the time specified delay period in the tendermint client's
// verification functions
func SetProcessedTime(clientStore storetypes.KVStore, height exported.Height, timeNs uint64) {
	_ = "STUB: not implemented"
	return
}

// GetProcessedTime gets the time (in nanoseconds) at which this chain received and processed a tendermint header.
// This is used to validate that a received packet has passed the time delay period.
func GetProcessedTime(clientStore storetypes.KVStore, height exported.Height) (uint64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// deleteProcessedTime deletes the processedTime for a given height
func deleteProcessedTime(clientStore storetypes.KVStore, height exported.Height) {
	_ = "STUB: not implemented"
	return
}

// ProcessedHeightKey returns the key under which the processed height will be stored in the client store.
func ProcessedHeightKey(height exported.Height) []byte { _ = "STUB: not implemented"; return nil }

// SetProcessedHeight stores the height at which a header was processed and the corresponding consensus state was created.
// This is useful when validating whether a packet has reached the specified block delay period in the tendermint client's
// verification functions
func SetProcessedHeight(clientStore storetypes.KVStore, consHeight, processedHeight exported.Height) {
	_ = "STUB: not implemented"
	return
}

// GetProcessedHeight gets the height at which this chain received and processed a tendermint header.
// This is used to validate that a received packet has passed the block delay period.
func GetProcessedHeight(clientStore storetypes.KVStore, height exported.Height) (exported.Height, bool) {
	_ = "STUB: not implemented"
	return *new(exported.Height), false
}

// deleteProcessedHeight deletes the processedHeight for a given height
func deleteProcessedHeight(clientStore storetypes.KVStore, height exported.Height) {
	_ = "STUB: not implemented"
	return
}

// IterationKey returns the key under which the consensus state key will be stored.
// The iteration key is a BigEndian representation of the consensus state key to support efficient iteration.
func IterationKey(height exported.Height) []byte { _ = "STUB: not implemented"; return nil }

// SetIterationKey stores the consensus state key under a key that is more efficient for ordered iteration
func SetIterationKey(clientStore storetypes.KVStore, height exported.Height) {
	_ = "STUB: not implemented"
	return
}

// GetIterationKey returns the consensus state key stored under the efficient iteration key.
// NOTE: This function is currently only used for testing purposes
func GetIterationKey(clientStore storetypes.KVStore, height exported.Height) []byte {
	_ = "STUB: not implemented"
	return nil
}

// deleteIterationKey deletes the iteration key for a given height
func deleteIterationKey(clientStore storetypes.KVStore, height exported.Height) {
	_ = "STUB: not implemented"
	return
}

// GetHeightFromIterationKey takes an iteration key and returns the height that it references
func GetHeightFromIterationKey(iterKey []byte) exported.Height {
	_ = "STUB: not implemented"
	return *new(exported.Height)
}

// IterateConsensusStateAscending iterates through the consensus states in ascending order. It calls the provided
// callback on each height, until stop=true is returned.
func IterateConsensusStateAscending(clientStore storetypes.KVStore, cb func(height exported.Height) (stop bool)) {
	_ = "STUB: not implemented"
	return
}

// GetNextConsensusState returns the lowest consensus state that is larger than the given height.
// The Iterator returns a storetypes.Iterator which iterates from start (inclusive) to end (exclusive).
// If the starting height exists in store, we need to call iterator.Next() to get the next consensus state.
// Otherwise, the iterator is already at the next consensus state so we can call iterator.Value() immediately.
func GetNextConsensusState(clientStore storetypes.KVStore, cdc codec.BinaryCodec, height exported.Height) (*ConsensusState, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// if iterator is at current height, ignore the consensus state at current height and get next height
// if iterator value is not at current height, it is already at next height.

// GetPreviousConsensusState returns the highest consensus state that is lower than the given height.
// The Iterator returns a storetypes.Iterator which iterates from the end (exclusive) to start (inclusive).
// Thus to get previous consensus state we call iterator.Value() immediately.
func GetPreviousConsensusState(clientStore storetypes.KVStore, cdc codec.BinaryCodec, height exported.Height) (*ConsensusState, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// PruneAllExpiredConsensusStates iterates over all consensus states for a given
// client store. If a consensus state is expired, it is deleted and its metadata
// is deleted. The number of consensus states pruned is returned.
func PruneAllExpiredConsensusStates(
	ctx sdk.Context, clientStore storetypes.KVStore,
	cdc codec.BinaryCodec, clientState *ClientState,
) int {
	_ = "STUB: not implemented"
	return 0
}

// consensus state should always be found

// Helper function for GetNextConsensusState and GetPreviousConsensusState
func getTmConsensusState(clientStore storetypes.KVStore, cdc codec.BinaryCodec, key []byte) (*ConsensusState, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func bigEndianHeightBytes(height exported.Height) []byte { _ = "STUB: not implemented"; return nil }

// setConsensusMetadata sets context time as processed time and set context height as processed height
// as this is internal tendermint light client logic.
// client state and consensus state will be set by client keeper
// set iteration key to provide ability for efficient ordered iteration of consensus states.
func setConsensusMetadata(ctx sdk.Context, clientStore storetypes.KVStore, height exported.Height) {
	_ = "STUB: not implemented"
	return
}

// setConsensusMetadataWithValues sets the consensus metadata with the provided values
func setConsensusMetadataWithValues(
	clientStore storetypes.KVStore, height,
	processedHeight exported.Height,
	processedTime uint64,
) {
	_ = "STUB: not implemented"
	return
}

// deleteConsensusMetadata deletes the metadata stored for a particular consensus state.
func deleteConsensusMetadata(clientStore storetypes.KVStore, height exported.Height) {
	_ = "STUB: not implemented"
	return
}
