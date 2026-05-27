package simapp

import (
	"testing"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	cmttypes "github.com/cometbft/cometbft/types"

	ibcwasmtypes "github.com/cosmos/ibc-go/modules/light-clients/08-wasm/v11/types"
)

func setup(tb testing.TB, chainID string, withGenesis bool, invCheckPeriod uint, mockVM ibcwasmtypes.WasmEngine) (*SimApp, GenesisState) {
	_ = "STUB: not implemented"
	return nil, *new(GenesisState)
}

// ensure unique folder

// SetupWithEmptyStore set up a simapp instance with empty DB
func SetupWithEmptyStore(tb testing.TB, mockVM ibcwasmtypes.WasmEngine) *SimApp {
	_ = "STUB: not implemented"
	return nil
}

// SetupWithGenesisValSet initializes a new SimApp with a validator set and genesis accounts
// that also act as delegators. For simplicity, each validator is bonded with a delegation
// of one consensus engine unit in the default token of the simapp from first genesis
// account. A Nop logger is set in SimApp.
func SetupWithGenesisValSetSnapshotter(t *testing.T, mockVM ibcwasmtypes.WasmEngine, valSet *cmttypes.ValidatorSet, genAccs []authtypes.GenesisAccount, balances ...banktypes.Balance) *SimApp {
	_ = "STUB: not implemented"
	return nil
}

// init chain will set the validator set and initialize the genesis accounts

// SetupWithSnapshotter initializes a new SimApp with a configured snapshot db. A Nop logger is set in SimApp.
func SetupWithSnapshotter(t *testing.T, mockVM ibcwasmtypes.WasmEngine) *SimApp {
	_ = "STUB: not implemented"
	return nil
}

// create validator set with single validator

// generate genesis account
