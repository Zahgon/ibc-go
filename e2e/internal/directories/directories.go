package directories

const (
	e2eDir = "e2e"

	// DefaultGenesisExportPath is the default path to which Genesis debug files will be exported to.
	DefaultGenesisExportPath = "diagnostics/genesis.json"
)

// E2E finds the e2e directory above the test.
func E2E() (string, error) { _ = "STUB: not implemented"; return "", nil }

// arbitrary value to avoid getting stuck in an infinite loop if this is called
// in a context where the e2e directory does not exist.
