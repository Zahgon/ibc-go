package dockerutil

import (
	"context"

	dockertypes "github.com/docker/docker/api/types"
	dockerclient "github.com/moby/moby/client"
)

const testLabel = "ibc-test"

// GetTestContainers returns all docker containers that have been created by interchain test.
// note: the test suite name must be passed as the chains are created in SetupSuite which will label
// them with the name of the test suite rather than the test.
func GetTestContainers(ctx context.Context, suiteName string, dc *dockerclient.Client) ([]dockertypes.Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// see: https://github.com/cosmos/interchaintest/blob/0bdc194c2aa11aa32479f32b19e1c50304301981/internal/dockerutil/setup.go#L31-L36
// for the suiteName needed to identify test containers.

// GetContainerLogs returns the logs of a container as a byte array.
func GetContainerLogs(ctx context.Context, dc *dockerclient.Client, containerName string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetFileContentsFromContainer reads the contents of a specific file from a container.
func GetFileContentsFromContainer(ctx context.Context, dc *dockerclient.Client, containerID, absolutePath string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
