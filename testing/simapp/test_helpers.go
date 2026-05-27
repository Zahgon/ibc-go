package simapp

import (
	"testing"
	"time"

	dbm "github.com/cosmos/cosmos-db"

	"cosmossdk.io/log/v2"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	abci "github.com/cometbft/cometbft/abci/types"
	cmttypes "github.com/cometbft/cometbft/types"
)

// SetupOptions defines arguments that are passed into `Simapp` constructor.
type SetupOptions struct {
	Logger  log.Logger
	DB      *dbm.MemDB
	AppOpts servertypes.AppOptions
}

// initSetup initializes a new SimApp. A Nop logger is set in SimApp.
func initSetup(withGenesis bool, invCheckPeriod uint) (*SimApp, GenesisState) {
	_ = "STUB: not implemented"
	return nil, *new(GenesisState)
}

// Setup initializes a new SimApp. A Nop logger is set in SimApp.
func Setup(t *testing.T, isCheckTx bool) *SimApp { _ = "STUB: not implemented"; return nil }

// create validator set with single validator

// generate genesis account

// SetupWithGenesisValSet initializes a new SimApp with a validator set and genesis accounts
// that also act as delegators. For simplicity, each validator is bonded with a delegation
// of one consensus engine unit in the default token of the simapp from first genesis
// account. A Nop logger is set in SimApp.
func SetupWithGenesisValSet(t *testing.T, valSet *cmttypes.ValidatorSet, genAccs []authtypes.GenesisAccount, balances ...banktypes.Balance) *SimApp {
	_ = "STUB: not implemented"
	return nil
}

// init chain will set the validator set and initialize the genesis accounts

// SignAndDeliver signs and delivers a transaction. No simulation occurs as the
// ibc testing package causes checkState and deliverState to diverge in block time.
//
// CONTRACT: BeginBlock must be called before this function.
func SignAndDeliver(
	tb testing.TB, txCfg client.TxConfig, app *bam.BaseApp, msgs []sdk.Msg,
	chainID string, accNums, accSeqs []uint64, expPass bool, blockTime time.Time, nextValHash []byte, priv ...cryptotypes.PrivKey,
) (*abci.ResponseFinalizeBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
