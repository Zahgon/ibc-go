package testsuite

import (
	"context"
	"sync"

	"github.com/cosmos/interchaintest/v11"
	"github.com/cosmos/interchaintest/v11/ibc"
	"github.com/cosmos/interchaintest/v11/testreporter"
	dockerclient "github.com/moby/moby/client"
	testifysuite "github.com/stretchr/testify/suite"
	"go.uber.org/zap"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/e2e/relayer"
	transfertypes "github.com/cosmos/ibc-go/v11/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
)

const (
	// ChainARelayerName is the name given to the relayer wallet on ChainA
	ChainARelayerName = "rlyA"
	// ChainBRelayerName is the name given to the relayer wallet on ChainB
	ChainBRelayerName = "rlyB"
	// DefaultGasValue is the default gas value used to configure tx.Factory
	DefaultGasValue = 500_000_0000
)

// E2ETestSuite has methods and functionality which can be shared among all test suites.
type E2ETestSuite struct {
	testifysuite.Suite

	// proposalIDs keeps track of the active proposal ID for each chain.
	proposalIDs map[string]uint64

	// chains is a list of chains that are created for the test suite.
	// each test suite has a single slice of chains that are used for all individual test
	// cases.
	chains         []ibc.Chain
	relayerWallets relayer.Map
	logger         *zap.Logger
	DockerClient   *dockerclient.Client
	network        string

	// pathNameIndex is the latest index to be used for generating chains
	pathNameIndex int64

	// testSuiteName is the name of the test suite, used to store chains under the test suite name.
	testSuiteName       string
	testPathsByTestName map[string][]string
	testPathsByChains   map[ibc.Chain]map[ibc.Chain]string
	channelByChains     map[string]map[ibc.Chain]map[ibc.Chain]ibc.ChannelOutput

	// channelLock ensures concurrent tests are not creating and accessing channels as the same time.
	channelLock sync.Mutex
	// relayerLock ensures concurrent tests are not accessing the pool of relayers as the same time.
	relayerLock sync.Mutex
	// relayerPool is a pool of relayers that can be used in tests.
	relayerPool []ibc.Relayer
	// testRelayerMap is a map of test suite names to relayers that are used in the test suite.
	// this is used as a cache after a relayer has been assigned to a test suite.
	testRelayerMap map[string]ibc.Relayer
}

// initState populates variables that are used across the test suite.
// note: this should be called only from SetupSuite.
func (s *E2ETestSuite) initState() { _ = "STUB: not implemented"; return }

// testSuiteName gets populated in the context of SetupSuite and stored as s.T().Name()
// will return the name of the suite and test when called from SetupTest or within the body of tests.
// the chains will be stored under the test suite name, so we need to store this for future use.

// initDockerClient creates a docker client and populates the network to be used for the test.
func (s *E2ETestSuite) initDockerClient() { _ = "STUB: not implemented"; return }

// configureGenesisDebugExport sets, if needed, env variables to enable exporting of Genesis debug files.
func (s *E2ETestSuite) configureGenesisDebugExport() { _ = "STUB: not implemented"; return }

// Set the export path.

// If no path is provided, use the default (e2e/diagnostics/genesis.json).

// This env variables are set by the interchain test code:
// https://github.com/cosmos/interchaintest/blob/7aa0fd6487f76238ab44231fdaebc34627bc5990/chain/cosmos/cosmos_chain.go#L1007-L1008

// Interchaintest adds a suffix (https://github.com/cosmos/interchaintest/blob/a3f4c7bcccf1925ffa6dc793a298f15497919a38/chainspec.go#L125)
// to the chain name, so we need to do the same.

// initializeRelayerPool pre-loads the relayer pool with n relayers.
// this is a workaround due to the restriction on relayer creation during the test
// ref: https://github.com/cosmos/interchaintest/issues/1153
// if the above issue is resolved, it should be possible to lazily create relayers in each test.
func (s *E2ETestSuite) initializeRelayerPool(n int) []ibc.Relayer {
	_ = "STUB: not implemented"
	return nil
}

// SetupChains creates the chains for the test suite, and also a relayer that is wired up to establish
// connections and channels between the chains.
func (s *E2ETestSuite) SetupChains(ctx context.Context, chainCount int, channelOptionsModifier ChainOptionModifier, chainSpecOpts ...ChainOptionConfiguration) {
	_ = "STUB: not implemented"
	return
}

// we skip path creation because we are just creating the chains and not connections/channels

// setup query paths for GRPC queries:

// CreatePaths creates paths between the chains using the provided client and channel options.
// The paths are created such that ChainA is connected to ChainB, ChainB is connected to ChainC etc.
func (s *E2ETestSuite) CreatePaths(clientOpts ibc.CreateClientOptions, channelOpts ibc.CreateChannelOptions, testName string) {
	_ = "STUB: not implemented"
	return
}

// CreatePath creates a path between chainA and chainB using the provided client and channel options.
func (s *E2ETestSuite) CreatePath(
	ctx context.Context,
	r ibc.Relayer,
	chainA ibc.Chain,
	chainB ibc.Chain,
	clientOpts ibc.CreateClientOptions,
	channelOpts ibc.CreateChannelOptions,
	testName string,
) (ibc.ChannelOutput, ibc.ChannelOutput) {
	_ = "STUB: not implemented"
	return *new(ibc.ChannelOutput), *new(ibc.ChannelOutput)
}

// Create new clients

// createChannelWithLock creates a channel between the two provided chains for the given test name. This applies a lock
// to ensure that the channels that are created are correctly mapped to the test that created them.
func (s *E2ETestSuite) createChannelWithLock(ctx context.Context, r ibc.Relayer, pathName, testName string, channelOpts ibc.CreateChannelOptions, chainA, chainB ibc.Chain) (ibc.ChannelOutput, ibc.ChannelOutput) {
	_ = "STUB: not implemented"
	// NOTE: we need to lock the creation of channels and applying of packet filters, as if we don't, the result
	// of `r.GetChannels` may return channels created by other relayers in different tests.
	return *new(ibc.ChannelOutput), *new(ibc.ChannelOutput)
}

func (s *E2ETestSuite) fetchChannelsBetweenChains(ctx context.Context, r ibc.Relayer, chainA ibc.Chain, chainB ibc.Chain) []ibc.ChannelOutput {
	_ = "STUB: not implemented"
	return nil
}

// This code is taken from interchaintests own Hermes query code, but since we need it here, it is copied - for now.
// extractJsonResult:

func (s *E2ETestSuite) mapChannel(testName string, fromChain ibc.Chain, toChain ibc.Chain, channel ibc.ChannelOutput) {
	_ = "STUB: not implemented"
	return
}

// getLatestChannel returns the latest channel from the list of channels.
func getLatestChannel(channels []ibc.ChannelOutput) ibc.ChannelOutput {
	_ = "STUB: not implemented"
	return *new(ibc.ChannelOutput)
}

// GetChannelsBetweenChains returns the channels between the two provided chains for the specified test.
func (s *E2ETestSuite) GetChannelBetweenChains(testname string, chainA ibc.Chain, chainB ibc.Chain) ibc.ChannelOutput {
	_ = "STUB: not implemented"
	return *new(ibc.ChannelOutput)
}

// GetRelayerForTest returns the relayer for the current test from the available pool of relayers.
// once a relayer has been returned to a test, it is cached and will be reused for the duration of the test.
func (s *E2ETestSuite) GetRelayerForTest(testName string) ibc.Relayer {
	_ = "STUB: not implemented"
	return *new(ibc.Relayer)
}

// remove the relayer from the pool

// GetRelayerUsers returns two ibc.Wallet instances which can be used for the relayer users
// on the two chains.
func (s *E2ETestSuite) GetRelayerUsers(ctx context.Context, testName string) (ibc.Wallet, ibc.Wallet) {
	_ = "STUB: not implemented"
	return *new(ibc.Wallet), *new(ibc.Wallet)
}

func (s *E2ETestSuite) FlushPackets(ctx context.Context, ibcrelayer ibc.Relayer, orderedChains []ibc.Chain) {
	_ = "STUB: not implemented"
	return
}

// Then we flush back the acknowledgements

// ChainOptionModifier is a function which accepts 2 chains as inputs, and returns a channel creation modifier function
// in order to conditionally modify the channel options based on the chains being used.
type ChainOptionModifier func(chainA, chainB ibc.Chain) func(options *ibc.CreateChannelOptions)

// newInterchain constructs a new interchain instance that creates channels between the chains.
func (s *E2ETestSuite) newInterchain(relayers []ibc.Relayer, chains []ibc.Chain, modificationProvider ChainOptionModifier) *interchaintest.Interchain {
	_ = "STUB: not implemented"
	return nil
}

// iterate through all chains, and create links such that there is a channel between
// - chainA and chainB
// - chainB and chainC
// - chainC and chainD etc

// make a modification to the channel options based on the chains which are being used.

// generatePathName generates the path name using the test suites name
func (s *E2ETestSuite) generatePathName() string { _ = "STUB: not implemented"; return "" }

func (s *E2ETestSuite) GetPaths(testName string) []string { _ = "STUB: not implemented"; return nil }

func (s *E2ETestSuite) GetPathByChains(chainA ibc.Chain, chainB ibc.Chain) string {
	_ = "STUB: not implemented"
	return ""
}

// GetPathName returns the name of a path at a specific index. This can be used in tests
// when the path name is required.
func GetPathName(idx int64) string { _ = "STUB: not implemented"; return "" }

// generatePath generates the path name using the test suites name. The indices provided specify which chains should be
// used. E.g. to generate a path between chain A and B, you would use 0 and 1, to specify between A and C, you would
// use 0 and 2 etc.
func (s *E2ETestSuite) generatePath(ctx context.Context, ibcrelayer ibc.Relayer, chainAIdx, chainBIdx int) string {
	_ = "STUB: not implemented"
	return ""
}

// SetupClients creates clients on chainA and chainB using the provided create client options
func (s *E2ETestSuite) SetupClients(ctx context.Context, ibcrelayer ibc.Relayer, opts ibc.CreateClientOptions) {
	_ = "STUB: not implemented"
	return
}

// UpdateClients updates clients on chainA and chainB
func (s *E2ETestSuite) UpdateClients(ctx context.Context, ibcrelayer ibc.Relayer, pathName string) {
	_ = "STUB: not implemented"
	return
}

// GetChains returns two chains that can be used in a test. The pair returned
// is unique to the current test being run. Note: this function does not create containers.
func (s *E2ETestSuite) GetChains() (ibc.Chain, ibc.Chain) {
	_ = "STUB: not implemented"
	return *new(ibc.Chain), *new(ibc.Chain)
}

// GetAllChains returns all chains that can be used in a test. The chains returned
// are unique to the current test being run. Note: this function does not create containers.
func (s *E2ETestSuite) GetAllChains() []ibc.Chain {
	_ = "STUB: not implemented"
	// chains are stored on a per test suite level
	return nil
}

// GetRelayerWallets returns the ibcrelayer wallets associated with the chains.
func (s *E2ETestSuite) GetRelayerWallets(ibcrelayer ibc.Relayer) (ibc.Wallet, ibc.Wallet, error) {
	_ = "STUB: not implemented"
	return *new(ibc.Wallet), *new(ibc.Wallet), nil
}

// RecoverRelayerWallets adds the corresponding ibcrelayer address to the keychain of the chain.
// This is useful if commands executed on the chains expect the relayer information to present in the keychain.
func (s *E2ETestSuite) RecoverRelayerWallets(ctx context.Context, ibcrelayer ibc.Relayer, testName string) (ibc.Wallet, ibc.Wallet, error) {
	_ = "STUB: not implemented"
	return *new(ibc.Wallet), *new(ibc.Wallet), nil
}

// StartRelayer starts the given ibcrelayer.
func (s *E2ETestSuite) StartRelayer(r ibc.Relayer, testName string) {
	_ = "STUB: not implemented"
	return
}

// wait for every chain to produce some blocks before using the relayer.

// StopRelayer stops the given ibcrelayer.
func (s *E2ETestSuite) StopRelayer(ctx context.Context, ibcrelayer ibc.Relayer) {
	_ = "STUB: not implemented"
	return
}

// RestartRelayer restarts the given relayer.
func (s *E2ETestSuite) RestartRelayer(ctx context.Context, ibcrelayer ibc.Relayer, testName string) {
	_ = "STUB: not implemented"
	return
}

// CreateUserOnChainA creates a user with the given amount of funds on chain A.
func (s *E2ETestSuite) CreateUserOnChainA(ctx context.Context, amount int64) ibc.Wallet {
	_ = "STUB: not implemented"
	return *new(ibc.Wallet)
}

// CreateUserOnChainB creates a user with the given amount of funds on chain B.
func (s *E2ETestSuite) CreateUserOnChainB(ctx context.Context, amount int64) ibc.Wallet {
	_ = "STUB: not implemented"
	return *new(ibc.Wallet)
}

// CreateUserOnChainC creates a user with the given amount of funds on chain C.
func (s *E2ETestSuite) CreateUserOnChainC(ctx context.Context, amount int64) ibc.Wallet {
	_ = "STUB: not implemented"
	return *new(ibc.Wallet)
}

// CreateUserOnChainD creates a user with the given amount of funds on chain C.
func (s *E2ETestSuite) CreateUserOnChainD(ctx context.Context, amount int64) ibc.Wallet {
	_ = "STUB: not implemented"
	return *new(ibc.Wallet)
}

// createWalletOnChainIndex creates a wallet with the given amount of funds on the chain of the given index.
func (s *E2ETestSuite) createWalletOnChainIndex(ctx context.Context, amount, chainIndex int64) ibc.Wallet {
	_ = "STUB: not implemented"
	return *new(ibc.Wallet)
}

// note the GetAndFundTestUsers requires the caller to wait for some blocks before the funds are accessible.

// GetChainANativeBalance gets the balance of a given user on chain A.
func (s *E2ETestSuite) GetChainANativeBalance(ctx context.Context, user ibc.Wallet) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GetChainBNativeBalance gets the balance of a given user on chain B.
func (s *E2ETestSuite) GetChainBNativeBalance(ctx context.Context, user ibc.Wallet) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GetChainBalanceForDenom returns the balance for a given denom given a chain.
func GetChainBalanceForDenom(ctx context.Context, chain ibc.Chain, denom string, user ibc.Wallet) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// AssertPacketRelayed asserts that the packet commitment does not exist on the sending chain.
// The packet commitment will be deleted upon a packet acknowledgement or timeout.
func (s *E2ETestSuite) AssertPacketRelayed(ctx context.Context, chain ibc.Chain, portID, channelID string, sequence uint64) {
	_ = "STUB: not implemented"
	return
}

// AssertPacketAcknowledged asserts that the packet has been acknowledged on the specified chain.
func (s *E2ETestSuite) AssertPacketAcknowledged(ctx context.Context, chain ibc.Chain, portID, channelID string, sequence uint64) {
	_ = "STUB: not implemented"
	return
}

// AssertHumanReadableDenom asserts that a human readable denom is present for a given chain.
func (s *E2ETestSuite) AssertHumanReadableDenom(ctx context.Context, chain ibc.Chain, counterpartyNativeDenom string, counterpartyChannel ibc.ChannelOutput) {
	_ = "STUB: not implemented"
	return
}

// createChains creates two separate chains in docker containers.
// test and can be retrieved with GetChains.
func (s *E2ETestSuite) createChains(chainCount int, chainOptions ChainOptions) []ibc.Chain {
	_ = "STUB: not implemented"
	return nil
}

// Take the first N specs

// this is intentionally called after the interchaintest.DockerSetup function. The above function registers a
// cleanup task which deletes all containers. By registering a cleanup function afterwards, it is executed first
// this allows us to process the logs before the containers are removed.

// initialise proposal ids for all chains.

// GetRelayerExecReporter returns a testreporter.RelayerExecReporter instances
// using the current test's testing.T.
func (s *E2ETestSuite) GetRelayerExecReporter() *testreporter.RelayerExecReporter {
	_ = "STUB: not implemented"
	return nil
}

// TransferChannelOptions configures both of the chains to have non-incentivized transfer channels.
func (*E2ETestSuite) TransferChannelOptions() ibc.CreateChannelOptions {
	_ = "STUB: not implemented"
	return *new(ibc.CreateChannelOptions)
}

// GetTimeoutHeight returns a timeout height of 1000 blocks above the current block height.
// This function should be used when the timeout is never expected to be reached
func (s *E2ETestSuite) GetTimeoutHeight(ctx context.Context, chain ibc.Chain) clienttypes.Height {
	_ = "STUB: not implemented"
	return *new(clienttypes.Height)
}

// GetIBCToken returns the denomination of the full token denom sent to the receiving channel
func GetIBCToken(fullTokenDenom string, portID, channelID string) transfertypes.Denom {
	_ = "STUB: not implemented"
	return *new(transfertypes.Denom)
}

// getValidatorsAndFullNodes returns the number of validators and full nodes respectively that should be used for
// the test. If the test is running in CI, more nodes are used, when running locally a single node is used by default to
// use less resources and allow the tests to run faster.
// both the number of validators and full nodes can be overwritten in a config file.
func getValidatorsAndFullNodes(chainIdx int) (int, int) { _ = "STUB: not implemented"; return 0, 0 }

// GetMsgTransfer returns a MsgTransfer that is constructed based on the channel version
func GetMsgTransfer(portID, channelID, version string, token sdk.Coin, sender, receiver string, timeoutHeight clienttypes.Height, timeoutTimestamp uint64, memo string) *transfertypes.MsgTransfer {
	_ = "STUB: not implemented"
	return nil
}

// SuiteName returns the name of the test suite.
func (s *E2ETestSuite) SuiteName() string { _ = "STUB: not implemented"; return "" }

// ThreeChainSetup provides the default behaviour to wire up 3 chains in the tests.
func ThreeChainSetup() ChainOptionConfiguration {
	_ = "STUB: not implemented"
	// copy all values of existing chains and tweak to make unique to new chain.
	return *new(ChainOptionConfiguration)
}

// nolint

// DefaultChannelOpts returns the default chain options for the test suite based on the provided chains.
func DefaultChannelOpts(chains []ibc.Chain) ibc.CreateChannelOptions {
	_ = "STUB: not implemented"
	return *new(ibc.CreateChannelOptions)
}
