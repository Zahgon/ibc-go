package diagnostics

import (
	"context"
	"testing"

	dockertypes "github.com/docker/docker/api/types"
	dockerclient "github.com/moby/moby/client"
)

const (
	dockerInspectFileName = "docker-inspect.json"
	defaultFilePerm       = 0o750
)

// Collect can be used in `t.Cleanup` and will copy all the of the container logs and relevant files
// into e2e/<test-suite>/<test-name>.log. These log files will be uploaded to GH upon test failure.
func Collect(t *testing.T, dc *dockerclient.Client, debugModeEnabled bool, suiteName string, chainNames ...string) {
	_ = "STUB: not implemented"
	return
}

// when we are not forcing log collection, we only upload upon test failing.

// getContainerName returns an either the ID of the container or stripped down human-readable
// version of the name if the name is non-empty.
//
// Note: You should still always use the ID  when interacting with the docker client.
func getContainerName(t *testing.T, container dockertypes.Container) string {
	_ = "STUB: not implemented"

	// container will always have an id, by may not have a name.
	return ""
}

// remove the test name from the container as the folder structure will provide this
// information already.

// fetchAndWriteDiagnosticsFile fetches the contents of a single file from the given container id and writes
// the contents of the file to a local path provided.
func fetchAndWriteDiagnosticsFile(ctx context.Context, dc *dockerclient.Client, containerID, localPath, absoluteFilePathInContainer string) error {
	_ = "STUB: not implemented"
	return nil
}

// fetchAndWriteDockerInspectOutput writes the contents of docker inspect to the specified file.
func fetchAndWriteDockerInspectOutput(ctx context.Context, dc *dockerclient.Client, containerID, localPath string) error {
	_ = "STUB: not implemented"
	return nil
}

// chainDiagnosticAbsoluteFilePaths returns a slice of absolute file paths (in the containers) which are the files that should be
// copied locally when fetching diagnostics.
func chainDiagnosticAbsoluteFilePaths(chainName string) []string {
	_ = "STUB: not implemented"
	return nil
}

// relayerDiagnosticAbsoluteFilePaths returns a slice of absolute file paths (in the containers) which are the files that should be
// copied locally when fetching diagnostics.
func relayerDiagnosticAbsoluteFilePaths() []string { _ = "STUB: not implemented"; return nil }
