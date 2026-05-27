package ibctesting

import (
	"encoding/json"
	"testing"

	sdkmath "cosmossdk.io/math"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/v2/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	cmttypes "github.com/cometbft/cometbft/types"

	"github.com/cosmos/ibc-go/v11/modules/core/keeper"
)

var DefaultTestingAppInit AppCreator = SetupTestingApp

type TestingApp interface {
	servertypes.ABCI

	// ibc-go additions
	GetBaseApp() *baseapp.BaseApp
	GetIBCKeeper() *keeper.Keeper
	GetTxConfig() client.TxConfig

	// Implemented by SimApp
	AppCodec() codec.Codec

	// Implemented by BaseApp
	LastCommitID() storetypes.CommitID
	LastBlockHeight() int64
}

func SetupTestingApp() (TestingApp, map[string]json.RawMessage) {
	_ = "STUB: not implemented"
	return *new(TestingApp), nil
}

// SetupWithGenesisValSet initializes a new SimApp with a validator set and genesis accounts
// that also act as delegators. For simplicity, each validator is bonded with a delegation
// of one consensus engine unit (10^6) in the default token of the simapp from first genesis
// account. A Nop logger is set in SimApp.
func SetupWithGenesisValSet(tb testing.TB, valSet *cmttypes.ValidatorSet, genAccs []authtypes.GenesisAccount, chainID string, powerReduction sdkmath.Int, balances ...banktypes.Balance) TestingApp {
	_ = "STUB: not implemented"
	return *new(TestingApp)
}

func setupWithGenesisValSet(tb testing.TB, valSet *cmttypes.ValidatorSet, genAccs []authtypes.GenesisAccount, chainID string, powerReduction sdkmath.Int, appCreator AppCreator, balances ...banktypes.Balance) TestingApp {
	_ = "STUB: not implemented"
	return *new(TestingApp)
}

// ensure baseapp has a chain-id set before running InitChain

// set genesis accounts

// set validators and delegations

// add bonded amount to bonded pool module account

// set validators and delegations

// update total supply

// init chain will set the validator set and initialize the genesis accounts
