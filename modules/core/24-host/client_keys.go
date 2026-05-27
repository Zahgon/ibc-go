package host

import (
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

// KeyClientStorePrefix defines the KVStore key prefix for IBC clients
var KeyClientStorePrefix = []byte("clients")

const (
	KeyClientState          = "clientState"
	KeyConsensusStatePrefix = "consensusStates"
)

// FullClientKey returns the full path of specific client path in the format:
// "clients/{clientID}/{path}" as a byte array.
func FullClientKey(clientID string, path []byte) []byte { _ = "STUB: not implemented"; return nil }

// PrefixedClientStoreKey returns a key which can be used for prefixed
// key store iteration. The prefix may be a clientType, clientID, or any
// valid key prefix which may be concatenated with the client store constant.
func PrefixedClientStoreKey(prefix []byte) []byte { _ = "STUB: not implemented"; return nil }

// FullClientStateKey takes a client identifier and returns a Key under which to store a
// particular client state.
func FullClientStateKey(clientID string) []byte { _ = "STUB: not implemented"; return nil }

// ClientStateKey returns a store key under which a particular client state is stored
// in a client prefixed store
func ClientStateKey() []byte { _ = "STUB: not implemented"; return nil }

// FullConsensusStateKey returns the store key for the consensus state of a particular
// client.
func FullConsensusStateKey(clientID string, height exported.Height) []byte {
	_ = "STUB: not implemented"
	return nil
}

// ConsensusStateKey returns the store key for a the consensus state of a particular
// client stored in a client prefixed store.
func ConsensusStateKey(height exported.Height) []byte { _ = "STUB: not implemented"; return nil }
