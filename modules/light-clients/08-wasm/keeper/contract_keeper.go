package keeper

import (
	wasmvm "github.com/CosmWasm/wasmvm/v3"
	wasmvmtypes "github.com/CosmWasm/wasmvm/v3/types"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/v2/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/modules/light-clients/08-wasm/v11/types"
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

var (
	VMGasRegister = types.NewDefaultWasmGasRegister()
	// wasmvmAPI is a wasmvm.GoAPI implementation that is passed to the wasmvm, it
	// doesn't implement any functionality, directly returning an error.
	wasmvmAPI = wasmvm.GoAPI{
		HumanizeAddress:     humanizeAddress,
		CanonicalizeAddress: canonicalizeAddress,
		ValidateAddress:     validateAddress,
	}
)

// instantiateContract calls vm.Instantiate with appropriate arguments.
func (k *Keeper) instantiateContract(ctx sdk.Context, clientID string, clientStore storetypes.KVStore, checksum types.Checksum, msg []byte) (*wasmvmtypes.ContractResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// callContract calls vm.Sudo with internally constructed gas meter and environment.
func (k *Keeper) callContract(ctx sdk.Context, clientID string, clientStore storetypes.KVStore, checksum types.Checksum, msg []byte) (*wasmvmtypes.ContractResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// queryContract calls vm.Query.
func (k *Keeper) queryContract(ctx sdk.Context, clientID string, clientStore storetypes.KVStore, checksum types.Checksum, msg []byte) (*wasmvmtypes.QueryResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// migrateContract calls vm.Migrate with internally constructed gas meter and environment.
func (k *Keeper) migrateContract(ctx sdk.Context, clientID string, clientStore storetypes.KVStore, checksum types.Checksum, msg []byte) (*wasmvmtypes.ContractResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WasmInstantiate accepts a message to instantiate a wasm contract, JSON encodes it and calls instantiateContract.
func (k *Keeper) WasmInstantiate(ctx sdk.Context, clientID string, clientStore storetypes.KVStore, cs *types.ClientState, payload types.InstantiateMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// Checksum should only be able to be modified during migration.

// WasmSudo calls the contract with the given payload and returns the result.
// WasmSudo returns an error if:
// - the contract call returns an error
// - the response of the contract call contains non-empty messages
// - the response of the contract call contains non-empty events
// - the response of the contract call contains non-empty attributes
// - the data bytes of the response cannot be unmarshaled into the result type
func (k *Keeper) WasmSudo(ctx sdk.Context, clientID string, clientStore storetypes.KVStore, cs *types.ClientState, payload types.SudoMsg) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Checksum should only be able to be modified during migration.

// WasmMigrate migrate calls the migrate entry point of the contract with the given payload and returns the result.
// WasmMigrate returns an error if:
// - the contract migration returns an error
func (k *Keeper) WasmMigrate(ctx sdk.Context, clientStore storetypes.KVStore, cs *types.ClientState, clientID string, payload []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// WasmQuery queries the contract with the given payload and returns the result.
// WasmQuery returns an error if:
// - the contract query returns an error
// - the data bytes of the response cannot be unmarshal into the result type
func (k *Keeper) WasmQuery(ctx sdk.Context, clientID string, clientStore storetypes.KVStore, cs *types.ClientState, payload types.QueryMsg) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// validatePostExecutionClientState validates that the contract has not many any invalid modifications
// to the client state during execution. It ensures that
// - the client state is still present
// - the client state can be unmarshaled successfully.
// - the client state is of type *ClientState
func validatePostExecutionClientState(clientStore storetypes.KVStore, cdc codec.BinaryCodec) (*types.ClientState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// unmarshalClientState unmarshals the client state from the given bytes.
func unmarshalClientState(cdc codec.BinaryCodec, bz []byte) (exported.ClientState, error) {
	_ = "STUB: not implemented"
	return *new(exported.ClientState), nil
}

// getEnv returns the state of the blockchain environment the contract is running on
func getEnv(ctx sdk.Context, contractAddr string) wasmvmtypes.Env {
	_ = "STUB: not implemented"
	return *new(wasmvmtypes.Env)
}

// safety checks before casting below

func humanizeAddress(canon []byte) (string, uint64, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func canonicalizeAddress(human string) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func validateAddress(human string) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// checkResponse returns an error if the response from a sudo, instantiate or migrate call
// to the Wasm VM contains messages, events or attributes.
func checkResponse(response *wasmvmtypes.Response) error {
	_ = "STUB: not implemented"
	// Only allow Data to flow back to us. SubMessages, Events and Attributes are not allowed.
	return nil
}
