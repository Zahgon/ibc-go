package testsuite

import (
	"github.com/cosmos/interchaintest/v11"
	"github.com/cosmos/interchaintest/v11/ibc"

	"github.com/cosmos/ibc-go/e2e/relayer"
)

const (
	// ChainImageEnv specifies the image that the chains will use. If left unspecified, it will
	// default to being determined based on the specified binary. E.g. ghcr.io/cosmos/ibc-go-simd
	ChainImageEnv = "CHAIN_IMAGE"
	// ChainATagEnv specifies the tag that Chain A will use.
	ChainATagEnv = "CHAIN_A_TAG"
	// ChainBTagEnv specifies the tag that Chain B will use. If unspecified
	// the value will default to the same value as Chain A.
	ChainBTagEnv = "CHAIN_B_TAG"
	// ChainCTagEnv specifies the tag that Chain C will use.
	// the value will default to the same value as Chain A.
	ChainCTagEnv = "CHAIN_C_TAG"
	// ChainDTagEnv specifies the tag that Chain D will use. If unspecified
	// the value will default to the same value as Chain A.
	ChainDTagEnv = "CHAIN_D_TAG"
	// RelayerIDEnv specifies the ID of the relayer to use.
	RelayerIDEnv = "RELAYER_ID"
	// ChainBinaryEnv binary is the binary that will be used for both chains.
	ChainBinaryEnv = "CHAIN_BINARY"
	// ChainUpgradePlanEnv specifies the upgrade plan name
	ChainUpgradePlanEnv = "CHAIN_UPGRADE_PLAN"
	// E2EConfigFilePathEnv allows you to specify a custom path for the config file to be used. It can be relative
	// or absolute.
	E2EConfigFilePathEnv = "E2E_CONFIG_PATH"
	// KeepContainersEnv instructs interchaintest to not delete the containers after a test has run.
	// this ensures that chain containers are not deleted after a test suite is run if other tests
	// depend on those chains.
	KeepContainersEnv = "KEEP_CONTAINERS"

	// defaultBinary is the default binary that will be used by the chains.
	defaultBinary = "simd"
	// defaultRlyTag is the tag that will be used if no relayer tag is specified.
	// all images are here https://github.com/cosmos/relayer/pkgs/container/relayer/versions
	defaultRlyTag = "latest"

	// defaultHermesTag is the tag that will be used if no relayer tag is specified for hermes.
	defaultHermesTag = "1.13.1"
	// defaultChainTag is the tag that will be used for the chains if none is specified.
	defaultChainTag = "main"
	// defaultConfigFileName is the default filename for the config file that can be used to configure
	// e2e tests. See sample.config.yaml or sample.config.extended.yaml as an example for what this should look like.
	defaultConfigFileName = ".ibc-go-e2e-config.yaml"
	// defaultCIConfigFileName is the default filename for the config file that should be used for CI.
	defaultCIConfigFileName = "ci-e2e-config.yaml"
)

// defaultChainNames contains the default name for chainA, chainB, ChainC and ChainD.
var defaultChainNames = []string{"simapp-a", "simapp-b", "simapp-c", "simapp-d"}

func getChainImage(binary string) string { _ = "STUB: not implemented"; return "" }

// TestConfig holds configuration used throughout the different e2e tests.
type TestConfig struct {
	// ChainConfigs holds configuration values related to the chains used in the tests.
	ChainConfigs []ChainConfig `yaml:"chains"`
	// RelayerConfigs holds all known relayer configurations that can be used in the tests.
	RelayerConfigs []relayer.Config `yaml:"relayers"`
	// ActiveRelayer specifies the relayer that will be used. It must match the ID of one of the entries in RelayerConfigs.
	ActiveRelayer string `yaml:"activeRelayer"`
	// CometBFTConfig holds values for configuring CometBFT.
	CometBFTConfig CometBFTConfig `yaml:"cometbft"`
	// DebugConfig holds configuration for miscellaneous options.
	DebugConfig DebugConfig `yaml:"debug"`
	// UpgradePlanName specifies which upgrade plan to use. It must match a plan name for an entry in the
	// list of UpgradeConfigs.
	UpgradePlanName string `yaml:"upgradePlanName"`
	// UpgradeConfigs provides a list of all possible upgrades.
	UpgradeConfigs []UpgradeConfig `yaml:"upgrades"`
}

// Validate validates the test configuration is valid for use within the tests.
// this should be called before using the configuration.
func (tc TestConfig) Validate() error { _ = "STUB: not implemented"; return nil }

// validateChains validates the chain configurations.
func (tc TestConfig) validateChains() error { _ = "STUB: not implemented"; return nil }

// clienttypes.ParseChainID is used to determine revision heights. If the chainIDs are not in the expected format,
// tests can fail with timeout errors.

// validateRelayers validates relayer configuration.
func (tc TestConfig) validateRelayers() error { _ = "STUB: not implemented"; return nil }

// GetUpgradeConfig returns the upgrade configuration for the current test configuration.
func (tc TestConfig) GetUpgradeConfig() UpgradeConfig {
	_ = "STUB: not implemented"
	return *new(UpgradeConfig)
}

// GetChainIndex returns the index of the chain with the given name, if it
// exists.
func (tc TestConfig) GetChainIndex(name string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// validateGenesisDebugConfig validates configuration of Genesis debug options/
func (tc TestConfig) validateGenesisDebugConfig() error { _ = "STUB: not implemented"; return nil }

// Verify that the provided chain exists in our config

// validateUpgradeConfig ensures the upgrade configuration is valid.
func (tc TestConfig) validateUpgradeConfig() error { _ = "STUB: not implemented"; return nil }

// the upgrade plan name specified must match one of the upgrade plans in the upgrade configs.

// GetActiveRelayerConfig returns the currently specified relayer config.
func (tc TestConfig) GetActiveRelayerConfig() *relayer.Config {
	_ = "STUB: not implemented"
	return nil
}

// GetChainNumValidators returns the number of validators for the specific chain index.
// default 1
func (tc TestConfig) GetChainNumValidators(idx int) int { _ = "STUB: not implemented"; return 0 }

// GetChainNumFullNodes returns the number of full nodes for the specific chain index.
// default 0
func (tc TestConfig) GetChainNumFullNodes(idx int) int { _ = "STUB: not implemented"; return 0 }

// GetChainID returns the chain-id for i. Assumes indicies are correct.
func (tc TestConfig) GetChainID(i int) string { _ = "STUB: not implemented"; return "" }

// GetChainAID returns the chain-id for chain A.
// NOTE: the default return value will ensure that ParseChainID will return 1 as the revision number.
func (tc TestConfig) GetChainAID() string { _ = "STUB: not implemented"; return "" }

// GetChainBID returns the chain-id for chain B.
// NOTE: the default return value will ensure that ParseChainID will return 1 as the revision number.
func (tc TestConfig) GetChainBID() string { _ = "STUB: not implemented"; return "" }

// GetChainCID returns the chain-id for chain C.
// NOTE: the default return value will ensure that ParseChainID will return 1 as the revision number.
func (tc TestConfig) GetChainCID() string { _ = "STUB: not implemented"; return "" }

// GetChainDID returns the chain-id for chain D.
// NOTE: the default return value will ensure that ParseChainID will return 1 as the revision number.
func (tc TestConfig) GetChainDID() string { _ = "STUB: not implemented"; return "" }

// GetChainName returns the name of the chain given an index.
func (tc TestConfig) GetChainName(idx int) string {
	_ = "STUB: not implemented"
	// Assumes that only valid indices are provided. We do the same in several other places.
	return ""
}

// GetGenesisChainName returns the name of the chain for which to dump Genesis files.
// If no chain is provided, it uses the default one (chainA).
func (tc TestConfig) GetGenesisChainName() string { _ = "STUB: not implemented"; return "" }

// UpgradeConfig holds values relevant to upgrade tests.
type UpgradeConfig struct {
	PlanName string `yaml:"planName"`
	Tag      string `yaml:"tag"`
}

// ChainConfig holds information about an individual chain used in the tests.
type ChainConfig struct {
	ChainID       string `yaml:"chainId"`
	Name          string `yaml:"name"`
	Image         string `yaml:"image"`
	Tag           string `yaml:"tag"`
	Binary        string `yaml:"binary"`
	NumValidators int    `yaml:"numValidators"`
	NumFullNodes  int    `yaml:"numFullNodes"`
}

type CometBFTConfig struct {
	LogLevel string `yaml:"logLevel"`
}

type GenesisDebugConfig struct {
	// DumpGenesisDebugInfo enables the output of Genesis debug files.
	DumpGenesisDebugInfo bool `yaml:"dumpGenesisDebugInfo"`

	// ExportFilePath specifies which path to export Genesis debug files to.
	ExportFilePath string `yaml:"filePath"`

	// ChainName represent which chain to get Genesis debug info for.
	ChainName string `yaml:"chainName"`
}

type DebugConfig struct {
	// DumpLogs forces the logs to be collected before removing test containers.
	DumpLogs bool `yaml:"dumpLogs"`

	// GenesisDebug contains debug information specific to Genesis.
	GenesisDebug GenesisDebugConfig `yaml:"genesis"`

	// KeepContainers specifies if the containers should be kept after the test suite is done.
	// NOTE: when running a full test suite, this value should be set to true in order to preserve
	// shared resources.
	KeepContainers bool `yaml:"keepContainers"`
}

// LoadConfig attempts to load a test configuration from the default file path.
// if any environment variables are specified, they will take precedence over the individual configuration
// options.
func LoadConfig() TestConfig { _ = "STUB: not implemented"; return *new(TestConfig) }

// getConfig returns the TestConfig with any environment variable overrides.
func getConfig() TestConfig { _ = "STUB: not implemented"; return *new(TestConfig) }

// If tags for chain C and D are not present in the file, also not set in the CI, fallback to A

// fromFile returns a TestConfig from a json file and a boolean indicating if the file was found.
func fromFile() (TestConfig, bool) { _ = "STUB: not implemented"; return *new(TestConfig), false }

// populateDefaults populates default values for the test config if
// certain required fields are not specified.
func populateDefaults(tc TestConfig) TestConfig { _ = "STUB: not implemented"; return *new(TestConfig) }

// If tag not given for chain C and D, set to chain A' tag

// applyEnvironmentVariableOverrides applies all environment variable changes to the config
// loaded from a file.
func applyEnvironmentVariableOverrides(fromFile TestConfig) TestConfig {
	_ = "STUB: not implemented"
	return *new(TestConfig)
}

// fromEnv returns a TestConfig constructed from environment variables.
func fromEnv() TestConfig { _ = "STUB: not implemented"; return *new(TestConfig) }

// getChainConfigsFromEnv returns the chain configs from environment variables.
func getChainConfigsFromEnv() []ChainConfig { _ = "STUB: not implemented"; return nil }

// getConfigFilePath returns the absolute path where the e2e config file should be.
func getConfigFilePath() string { _ = "STUB: not implemented"; return "" }

// running locally.

// TODO: remove in https://github.com/cosmos/ibc-go/issues/4697
// getDefaultHermesRelayerConfig returns the default config for the hermes relayer.
func getDefaultHermesRelayerConfig() relayer.Config {
	_ = "STUB: not implemented"
	return *new(relayer.Config)
}

// TODO: remove in https://github.com/cosmos/ibc-go/issues/4697
// getDefaultRlyRelayerConfig returns the default config for the golang relayer.
func getDefaultRlyRelayerConfig() relayer.Config {
	_ = "STUB: not implemented"
	return *new(relayer.Config)
}

func GetChainATag() string { _ = "STUB: not implemented"; return "" }

func GetChainBTag() string { _ = "STUB: not implemented"; return "" }

// IsCI returns true if the tests are running in CI, false is returned
// if the tests are running locally.
// Note: github actions passes a CI env value of true by default to all runners.
func IsCI() bool { _ = "STUB: not implemented"; return false }

// IsFork returns true if the tests are running in fork mode, false is returned otherwise.
func IsFork() bool { _ = "STUB: not implemented"; return false }

// IsRunSuite returns true if the tests are running in suite mode, false is returned otherwise.
func IsRunSuite() bool { _ = "STUB: not implemented"; return false }

func isEnvTrue(env string) bool { _ = "STUB: not implemented"; return false }

// ChainOptions stores chain configurations for the chains that will be
// created for the tests. They can be modified by passing ChainOptionConfiguration
// to E2ETestSuite.GetChains.
type ChainOptions struct {
	ChainSpecs       []*interchaintest.ChainSpec
	SkipPathCreation bool
	RelayerCount     int
}

// ChainOptionConfiguration enables arbitrary configuration of ChainOptions.
type ChainOptionConfiguration func(options *ChainOptions)

// DefaultChainOptions returns the default configuration for required number of chains.
// These options can be configured by passing configuration functions to E2ETestSuite.GetChains.
func DefaultChainOptions(chainCount int) (ChainOptions, error) {
	_ = "STUB: not implemented"
	return *new(ChainOptions), nil
}

// if running a single test, only one relayer is needed.

// arbitrary number that will not be required if https://github.com/cosmos/interchaintest/issues/1153 is resolved.
// It can be overridden in individual test suites in SetupSuite if required.

// newDefaultSimappConfig creates an ibc configuration for simd.
func newDefaultSimappConfig(cc ChainConfig, name, chainID, denom string, cometCfg CometBFTConfig) ibc.ChainConfig {
	_ = "STUB: not implemented"
	return *new(ibc.ChainConfig)
}

// change to debug in the e2e test config to increase cometbft logging.

// getGenesisModificationFunction returns a genesis modification function that handles the GenesisState type
// correctly depending on if the govv1beta1 gov module is used or if govv1 is being used.
func getGenesisModificationFunction(cc ChainConfig) func(ibc.ChainConfig, []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Remove after we drop v7 support (this is only needed right now because of v6 -> v7 upgrade tests)

// defaultGovv1ModifyGenesis will only modify governance params to ensure the voting period and minimum deposit
// are functional for e2e testing purposes.
func defaultGovv1ModifyGenesis(version string) func(ibc.ChainConfig, []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil
}

// in older version < v8, tmjson marshal must be used.
// regular json marshalling must be used for v8 and above as the
// sdk is de-coupled from comet.

// defaultGovv1Beta1ModifyGenesis will only modify governance params to ensure the voting period and minimum deposit
// // are functional for e2e testing purposes.
func defaultGovv1Beta1ModifyGenesis(version string) func(ibc.ChainConfig, []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil
}

// modifyGovV1AppState takes the existing gov app state and marshals it to a govv1 GenesisState.
func modifyGovV1AppState(chainConfig ibc.ChainConfig, govAppState []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// modifyGovv1Beta1AppState takes the existing gov app state and marshals it to a govv1beta1 GenesisState.
func modifyGovv1Beta1AppState(chainConfig ibc.ChainConfig, govAppState []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// modifyClientGenesisAppState takes the existing ibc app state and marshals it to an ibc GenesisState.
func modifyClientGenesisAppState(ibcAppState []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// modifyChannelGenesisAppState takes the existing ibc app state, unmarshals it to a map and removes the `params` entry from ibc channel genesis.
// It marshals and returns the ibc GenesisState JSON map as bytes.
func modifyChannelGenesisAppState(ibcAppState []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// be ashamed, be very ashamed

func modifyChannelV2GenesisAppState(ibcAppState []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func modifyClientV2GenesisAppState(ibcAppState []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
