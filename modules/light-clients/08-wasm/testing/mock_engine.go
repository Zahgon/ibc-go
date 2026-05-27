package testing

import (
	wasmvm "github.com/CosmWasm/wasmvm/v3"
	wasmvmtypes "github.com/CosmWasm/wasmvm/v3/types"

	"github.com/cosmos/ibc-go/modules/light-clients/08-wasm/v11/types"
)

const DefaultGasUsed = uint64(1)

var (
	_ types.WasmEngine = (*MockWasmEngine)(nil)

	// queryTypes contains all the possible query message types.
	queryTypes = [...]any{types.StatusMsg{}, types.TimestampAtHeightMsg{}, types.VerifyClientMessageMsg{}, types.CheckForMisbehaviourMsg{}}

	// sudoTypes contains all the possible sudo message types.
	sudoTypes = [...]any{types.UpdateStateMsg{}, types.UpdateStateOnMisbehaviourMsg{}, types.VerifyUpgradeAndUpdateStateMsg{}, types.VerifyMembershipMsg{}, types.VerifyNonMembershipMsg{}, types.MigrateClientStoreMsg{}}
)

type (
	// queryFn is a callback function that is invoked when a specific query message type is received.
	queryFn func(checksum wasmvm.Checksum, env wasmvmtypes.Env, queryMsg []byte, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) (*wasmvmtypes.QueryResult, uint64, error)

	// sudoFn is a callback function that is invoked when a specific sudo message type is received.
	sudoFn func(checksum wasmvm.Checksum, env wasmvmtypes.Env, sudoMsg []byte, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) (*wasmvmtypes.ContractResult, uint64, error)
)

// MockWasmEngine implements types.WasmEngine for testing purposes. One or multiple messages can be stubbed.
// Without a stub function a panic is thrown.
// ref: https://github.com/CosmWasm/wasmd/blob/v0.42.0/x/wasm/keeper/wasmtesting/mock_engine.go#L19
type MockWasmEngine struct {
	StoreCodeFn          func(code wasmvm.WasmCode, gasLimit uint64) (wasmvmtypes.Checksum, uint64, error)
	StoreCodeUncheckedFn func(code wasmvm.WasmCode) (wasmvm.Checksum, error)
	InstantiateFn        func(checksum wasmvm.Checksum, env wasmvmtypes.Env, info wasmvmtypes.MessageInfo, initMsg []byte, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) (*wasmvmtypes.ContractResult, uint64, error)
	MigrateFn            func(checksum wasmvm.Checksum, env wasmvmtypes.Env, migrateMsg []byte, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) (*wasmvmtypes.ContractResult, uint64, error)
	GetCodeFn            func(checksum wasmvm.Checksum) (wasmvm.WasmCode, error)
	PinFn                func(checksum wasmvm.Checksum) error
	UnpinFn              func(checksum wasmvm.Checksum) error

	// queryCallbacks contains a mapping of queryMsg field type name to callback function.
	queryCallbacks map[string]queryFn
	sudoCallbacks  map[string]sudoFn

	// contracts contains a mapping of checksum to code.
	storedContracts map[uint32][]byte
}

// NewMockWasmEngine creates and returns a new instance of the mock wasmvm for testing purposes.
// Each callback method of the mock wasmvm can be overridden to assign specific functionality.
// Default functionality is assigned for StoreCode, StoreCodeUnchecked and GetCode. Both Pin and Unpin are implemented as no-op methods.
// All other callbacks stored in the query and sudo callback maps panic. Use RegisterQueryCallback and RegisterSudoCallback methods
// to assign expected behaviour for test cases.
func NewMockWasmEngine() *MockWasmEngine { _ = "STUB: not implemented"; return nil }

// Set up default behavior for Store/Pin/Get

// RegisterQueryCallback registers a callback for a specific message type.
func (m *MockWasmEngine) RegisterQueryCallback(queryMessage any, fn queryFn) {
	_ = "STUB: not implemented"
	return
}

// RegisterSudoCallback registers a callback for a specific sudo message type.
func (m *MockWasmEngine) RegisterSudoCallback(sudoMessage any, fn sudoFn) {
	_ = "STUB: not implemented"
	return
}

// StoreCode implements the WasmEngine interface.
func (m *MockWasmEngine) StoreCode(code wasmvm.WasmCode, gasLimit uint64) (wasmvmtypes.Checksum, uint64, error) {
	_ = "STUB: not implemented"
	return *new(wasmvmtypes.Checksum), 0, nil
}

// StoreCodeUnchecked implements the WasmEngine interface.
func (m *MockWasmEngine) StoreCodeUnchecked(code wasmvm.WasmCode) (wasmvm.Checksum, error) {
	_ = "STUB: not implemented"
	return *new(wasmvm.Checksum), nil
}

// Instantiate implements the WasmEngine interface.
func (m *MockWasmEngine) Instantiate(checksum wasmvm.Checksum, env wasmvmtypes.Env, info wasmvmtypes.MessageInfo, initMsg []byte, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) (*wasmvmtypes.ContractResult, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// Query implements the WasmEngine interface.
func (m *MockWasmEngine) Query(checksum wasmvm.Checksum, env wasmvmtypes.Env, queryMsg []byte, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) (*wasmvmtypes.QueryResult, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// Migrate implements the WasmEngine interface.
func (m *MockWasmEngine) Migrate(checksum wasmvm.Checksum, env wasmvmtypes.Env, migrateMsg []byte, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) (*wasmvmtypes.ContractResult, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// Sudo implements the WasmEngine interface.
func (m *MockWasmEngine) Sudo(checksum wasmvm.Checksum, env wasmvmtypes.Env, sudoMsg []byte, store wasmvm.KVStore, goapi wasmvm.GoAPI, querier wasmvm.Querier, gasMeter wasmvm.GasMeter, gasLimit uint64, deserCost wasmvmtypes.UFraction) (*wasmvmtypes.ContractResult, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// GetCode implements the WasmEngine interface.
func (m *MockWasmEngine) GetCode(checksum wasmvm.Checksum) (wasmvm.WasmCode, error) {
	_ = "STUB: not implemented"
	return *new(wasmvm.WasmCode), nil
}

// Pin implements the WasmEngine interface.
func (m *MockWasmEngine) Pin(checksum wasmvm.Checksum) error { _ = "STUB: not implemented"; return nil }

// Unpin implements the WasmEngine interface.
func (m *MockWasmEngine) Unpin(checksum wasmvm.Checksum) error {
	_ = "STUB: not implemented"
	return nil
}

// getQueryMsgPayloadTypeName extracts the name of the struct that is populated.
// this value is used as a key to map to a callback function to handle that message type.
func getQueryMsgPayloadTypeName(queryMsgBz []byte) string { _ = "STUB: not implemented"; return "" }

// getSudoMsgPayloadTypeName extracts the name of the struct that is populated.
// this value is used as a key to map to a callback function to handle that message type.
func getSudoMsgPayloadTypeName(sudoMsgBz []byte) string { _ = "STUB: not implemented"; return "" }
