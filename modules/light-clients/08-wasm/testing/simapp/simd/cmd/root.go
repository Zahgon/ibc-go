package cmd

import (
	"os"

	dbm "github.com/cosmos/cosmos-db"
	"github.com/spf13/cobra"

	"cosmossdk.io/client/v2/autocli"
	"cosmossdk.io/log/v2"

	"github.com/cosmos/cosmos-sdk/client"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	cmtcfg "github.com/cometbft/cometbft/config"

	"github.com/cosmos/ibc-go/modules/light-clients/08-wasm/v11/testing/simapp"
	"github.com/cosmos/ibc-go/modules/light-clients/08-wasm/v11/testing/simapp/params"
)

// NewRootCmd creates a new root command for simd. It is called once in the
// main function.
func NewRootCmd() *cobra.Command {
	_ = "STUB: not implemented"
	// we "pre"-instantiate the application for getting the injected/configured encoding configuration
	// note, this is not necessary when using app wiring, as depinject can be directly used (see root_v2.go)
	return nil
}

// In simapp, we don't use any prefix for env variables.

// set the default command outputs

// This needs to go after ReadFromClientConfig, as that function
// sets the RPC client needed for SIGN_MODE_TEXTUAL.
//nolint:gocritic // we know we aren't appending to the same slice

func enrichAutoCliOpts(autoCliOpts autocli.AppOptions, clientCtx client.Context) (autocli.AppOptions, error) {
	_ = "STUB: not implemented"
	return *new(autocli.AppOptions), nil
}

// initCometBFTConfig helps to override default CometBFT Config values.
// return cmtcfg.DefaultConfig if no custom configuration is required for the application.
func initCometBFTConfig() *cmtcfg.Config { _ = "STUB: not implemented"; return nil }

// these values put a higher strain on node memory
// cfg.P2P.MaxNumInboundPeers = 100
// cfg.P2P.MaxNumOutboundPeers = 40

// initAppConfig helps to override default appConfig template and configs.
// return "", nil if no custom configuration is required for the application.
func initAppConfig() (string, any) {
	_ = "STUB: not implemented"
	// The following code snippet is just for reference.
	return "", *new(any)
}

// WASMConfig defines configuration for the wasm module.

// This is the maximum sdk gas (wasm and storage) that we allow for any x/wasm "smart" queries

// Address defines the gRPC-web server to listen on

// Optionally allow the chain developer to overwrite the SDK's default
// server config.

// The SDK's default minimum gas price is set to "" (empty value) inside
// app.toml. If left empty by validators, the node will halt on startup.
// However, the chain developer can set a default app.toml value for their
// validators here.
//
// In summary:
// - if you leave srvCfg.MinGasPrices = "", all validators MUST tweak their
//   own app.toml config,
// - if you set srvCfg.MinGasPrices non-empty, validators CAN tweak their
//   own app.toml to override, or use this default value.
//
// In simapp, we set the min gas prices to 0.

// srvCfg.BaseConfig.IAVLDisableFastNode = true // disable fastnode by default

func initRootCmd(rootCmd *cobra.Command, encodingConfig params.EncodingConfig, basicManager module.BasicManager) {
	_ = "STUB: not implemented"
	return
}

// add keybase, auxiliary RPC, query, genesis, and tx child commands

func addModuleInitFlags(startCmd *cobra.Command) { _ = "STUB: not implemented"; return }

func queryCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func txCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// genesisCommand builds genesis-related `simd genesis` command. Users may provide application specific commands as a parameter
func genesisCommand(encodingConfig params.EncodingConfig, basicManager module.BasicManager, cmds ...*cobra.Command) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

// newApp creates the application
func newApp(
	logger log.Logger,
	db dbm.DB,
	appOpts servertypes.AppOptions,
) servertypes.Application {
	_ = "STUB: not implemented"
	return *new(servertypes.Application)
}

// appExport creates a new simapp (optionally at a given height) and exports state.
func appExport(
	logger log.Logger,
	db dbm.DB,
	height int64,
	forZeroHeight bool,
	jailAllowedAddrs []string,
	appOpts servertypes.AppOptions,
	modulesToExport []string,
) (servertypes.ExportedApp, error) {
	_ = "STUB: not implemented"
	return *

	// this check is necessary as we use the flag in x/upgrade.
	// we can exit more gracefully by checking the flag here.
	new(servertypes.ExportedApp), nil
}

// overwrite the FlagInvCheckPeriod

var tempDir = func() string {
	dir, err := os.MkdirTemp("", "simapp")
	if err != nil {
		dir = simapp.DefaultNodeHome
	}
	defer os.RemoveAll(dir)

	return dir
}

// CheckLibwasmVersion ensures that the libwasmvm version loaded at runtime matches the version
// of the github.com/CosmWasm/wasmvm dependency in go.mod.
// Ref: https://github.com/cosmos/ibc-go/issues/4821#issuecomment-1747240445
func CheckLibwasmVersion(wasmExpectedVersion string) error { _ = "STUB: not implemented"; return nil }

type preRunFn func(cmd *cobra.Command, args []string) error

func chainPreRuns(pfns ...preRunFn) preRunFn { _ = "STUB: not implemented"; return *new(preRunFn) }

func getExpectedLibwasmVersion() string { _ = "STUB: not implemented"; return "" }
