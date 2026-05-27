package relayer

import (
	"context"
	"testing"

	"github.com/cosmos/interchaintest/v11/ibc"
	"github.com/cosmos/interchaintest/v11/relayer/hermes"
	dockerclient "github.com/moby/moby/client"
	"go.uber.org/zap"
)

const (
	Rly    = "rly"
	Hermes = "hermes"

	HermesRelayerRepository = "ghcr.io/informalsystems/hermes"
	hermesRelayerUser       = "2000:2000"
	RlyRelayerRepository    = "ghcr.io/cosmos/relayer"
	rlyRelayerUser          = "100:1000"

	// relativeHermesConfigFilePath is the path to the hermes config file relative to the home directory within the container.
	relativeHermesConfigFilePath = ".hermes/config.toml"
)

// Config holds configuration values for the relayer used in the tests.
type Config struct {
	// Tag is the tag used for the relayer image.
	Tag string `yaml:"tag"`
	// ID specifies the type of relayer that this is.
	ID string `yaml:"id"`
	// Image is the image that should be used for the relayer.
	Image string `yaml:"image"`
}

// New returns an implementation of ibc.Relayer depending on the provided RelayerType.
func New(t *testing.T, cfg Config, logger *zap.Logger, dockerClient *dockerclient.Client, network string) ibc.Relayer {
	_ = "STUB: not implemented"
	return *new(ibc.Relayer)
}

// ApplyPacketFilter applies a packet filter to the hermes config file, which specifies a complete set of channels
// to watch for packets.
func ApplyPacketFilter(ctx context.Context, t *testing.T, r ibc.Relayer, chainID string, channels []ibc.ChannelOutput) error {
	_ = "STUB: not implemented"
	return nil
}

// [chains.packet_filter]
//	# policy = 'allow'
//	# list = [
//	#   ['ica*', '*'],
//	#   ['transfer', 'channel-0'],
//	# ]

// TODO(chatton): explicitly enable watching of ICA channels
// this will ensure the ICA tests pass, but this will need to be modified to make sure
// ICA tests will succeed in parallel.

// we explicitly override the full list, this allows this function to provide a complete set of channels to watch.

// modifyHermesConfigFile reads the hermes config file, applies a modification function and returns an error if any.
func modifyHermesConfigFile(ctx context.Context, h *hermes.Relayer, modificationFn func(map[string]any) error) error {
	_ = "STUB: not implemented"
	return nil
}

// newCosmosRelayer returns an instance of the go relayer.
// Options are used to allow for relayer version selection and specifying the default processing option.
func newCosmosRelayer(t *testing.T, tag string, logger *zap.Logger, dockerClient *dockerclient.Client, network, relayerImage string) ibc.Relayer {
	_ = "STUB: not implemented"
	return *new(ibc.Relayer)
}

// relayer processes via events

// newHermesRelayer returns an instance of the hermes relayer.
func newHermesRelayer(t *testing.T, tag string, logger *zap.Logger, dockerClient *dockerclient.Client, network, relayerImage string) ibc.Relayer {
	_ = "STUB: not implemented"
	return *new(ibc.Relayer)
}

// Map is a mapping from test names to a relayer set for that test.
type Map map[string]map[ibc.Wallet]bool

// AddRelayer adds the given relayer to the relayer set for the given test name.
func (r Map) AddRelayer(testName string, ibcrelayer ibc.Wallet) { _ = "STUB: not implemented"; return }

// ContainsRelayer returns true if the given relayer is in the relayer set for the given test name.
func (r Map) ContainsRelayer(testName string, wallet ibc.Wallet) bool {
	_ = "STUB: not implemented"
	return false
}
