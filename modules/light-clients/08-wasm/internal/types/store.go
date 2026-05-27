package types

import (
	wasmvmtypes "github.com/CosmWasm/wasmvm/v3/types"

	storetypes "github.com/cosmos/cosmos-sdk/store/v2/types"
)

var (
	_ wasmvmtypes.KVStore = &StoreAdapter{}
	_ storetypes.KVStore  = &ClientRecoveryStore{}

	SubjectPrefix    = []byte("subject/")
	SubstitutePrefix = []byte("substitute/")
)

// ClientRecoveryStore combines two KVStores into one.
//
// Both stores are used for reads, but only the subjectStore is used for writes. For all operations, the key
// is checked to determine which types to use and must be prefixed with either "subject/" or "substitute/" accordingly.
// If the key is not prefixed with either "subject/" or "substitute/", a default action is taken (e.g. no-op for Set/Delete).
type ClientRecoveryStore struct {
	subjectStore    storetypes.KVStore
	substituteStore storetypes.KVStore
}

// NewClientRecoveryStore returns a new instance of a ClientRecoveryStore
func NewClientRecoveryStore(subjectStore, substituteStore storetypes.KVStore) ClientRecoveryStore {
	_ = "STUB: not implemented"
	return *new(ClientRecoveryStore)
}

// Get implements the storetypes.KVStore interface. It allows reads from both the subjectStore and substituteStore.
//
// Get will return an empty byte slice if the key is not prefixed with either "subject/" or "substitute/".
func (s ClientRecoveryStore) Get(key []byte) []byte { _ = "STUB: not implemented"; return nil }

// return a nil byte slice as KVStore.Get() does by default

// Has implements the storetypes.KVStore interface. It allows reads from both the subjectStore and substituteStore.
//
// Note: contracts do not have access to the Has method, it is only implemented here to satisfy the storetypes.KVStore interface.
func (s ClientRecoveryStore) Has(key []byte) bool { _ = "STUB: not implemented"; return false }

// return false as value when types is not found

// Set implements the storetypes.KVStore interface. It allows writes solely to the subjectStore.
//
// Set will no-op if the key is not prefixed with "subject/".
func (s ClientRecoveryStore) Set(key, value []byte) { _ = "STUB: not implemented"; return }

// no-op

// Delete implements the storetypes.KVStore interface. It allows deletions solely to the subjectStore.
//
// Delete will no-op if the key is not prefixed with "subject/".
func (s ClientRecoveryStore) Delete(key []byte) { _ = "STUB: not implemented"; return }

// no-op

// Iterator implements the storetypes.KVStore interface. It allows iteration over both the subjectStore and substituteStore.
//
// Iterator will return a closed iterator if the start or end keys are not prefixed with either "subject/" or "substitute/".
func (s ClientRecoveryStore) Iterator(start, end []byte) storetypes.Iterator {
	_ = "STUB: not implemented"
	return *new(storetypes.Iterator)
}

// ReverseIterator implements the storetypes.KVStore interface. It allows iteration over both the subjectStore and substituteStore.
//
// ReverseIterator will return a closed iterator if the start or end keys are not prefixed with either "subject/" or "substitute/".
func (s ClientRecoveryStore) ReverseIterator(start, end []byte) storetypes.Iterator {
	_ = "STUB: not implemented"
	return *new(storetypes.Iterator)
}

// GetStoreType implements the storetypes.KVStore interface, it is implemented solely to satisfy the interface.
func (s ClientRecoveryStore) GetStoreType() storetypes.StoreType {
	_ = "STUB: not implemented"
	return *new(storetypes.StoreType)
}

// CacheWrap implements the storetypes.KVStore interface, it is implemented solely to satisfy the interface.
func (s ClientRecoveryStore) CacheWrap() storetypes.CacheWrap {
	_ = "STUB: not implemented"
	return *new(storetypes.CacheWrap)
}

// GetStore returns the types to be used for the given key and a boolean flag indicating if that type was found.
// If the key is prefixed with "subject/", the subjectStore is returned. If the key is prefixed with "substitute/",
// the substituteStore is returned.
//
// If the key is not prefixed with either "subject/" or "substitute/", a nil types is returned and the boolean flag is false.
func (s ClientRecoveryStore) GetStore(prefix []byte) (storetypes.KVStore, bool) {
	_ = "STUB: not implemented"
	return *new(storetypes.KVStore), false
}

// closedIterator returns an iterator that is always closed, used when Iterator() or ReverseIterator() is called
// with an invalid prefix or start/end key.
func (s ClientRecoveryStore) closedIterator() storetypes.Iterator {
	_ = "STUB: not implemented"
	// Create a dummy iterator that is always closed right away.
	return *new(storetypes.Iterator)
}

// SplitPrefix splits the key into the prefix and the key itself, if the key is prefixed with either "subject/" or "substitute/".
// If the key is not prefixed with either "subject/" or "substitute/", the prefix is nil.
func SplitPrefix(key []byte) ([]byte, []byte) { _ = "STUB: not implemented"; return nil, nil }

// StoreAdapter bridges the SDK types implementation to wasmvm one. It implements the wasmvmtypes.KVStore interface.
type StoreAdapter struct {
	parent storetypes.KVStore
}

// NewStoreAdapter constructor
func NewStoreAdapter(s storetypes.KVStore) *StoreAdapter { _ = "STUB: not implemented"; return nil }

// Get implements the wasmvmtypes.KVStore interface.
func (s StoreAdapter) Get(key []byte) []byte { _ = "STUB: not implemented"; return nil }

// Set implements the wasmvmtypes.KVStore interface.
func (s StoreAdapter) Set(key, value []byte) { _ = "STUB: not implemented"; return }

// Delete implements the wasmvmtypes.KVStore interface.
func (s StoreAdapter) Delete(key []byte) { _ = "STUB: not implemented"; return }

// Iterator implements the wasmvmtypes.KVStore interface.
func (s StoreAdapter) Iterator(start, end []byte) wasmvmtypes.Iterator {
	_ = "STUB: not implemented"
	return *new(wasmvmtypes.Iterator)
}

// ReverseIterator implements the wasmvmtypes.KVStore interface.
func (s StoreAdapter) ReverseIterator(start, end []byte) wasmvmtypes.Iterator {
	_ = "STUB: not implemented"
	return *new(wasmvmtypes.Iterator)
}
