package host

// ParseIdentifier parses the sequence from the identifier using the provided prefix. This function
// does not need to be used by counterparty chains. SDK generated connection and channel identifiers
// are required to use this format.
func ParseIdentifier(identifier, prefix string) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// sanity check

// MustParseClientStatePath returns the client ID from a client state path. It panics
// if the provided path is invalid or if the clientID is empty.
func MustParseClientStatePath(path string) string { _ = "STUB: not implemented"; return "" }

// parseClientStatePath returns the client ID from a client state path. It returns
// an error if the provided path is invalid.
func parseClientStatePath(path string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// ParseConnectionPath returns the connection ID from a full path. It returns
// an error if the provided path is invalid.
func ParseConnectionPath(path string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// ParseChannelPath returns the port and channel ID from a full path. It returns
// an error if the provided path is invalid.
func ParseChannelPath(path string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// MustParseConnectionPath returns the connection ID from a full path. Panics
// if the provided path is invalid.
func MustParseConnectionPath(path string) string { _ = "STUB: not implemented"; return "" }

// MustParseChannelPath returns the port and channel ID from a full path. Panics
// if the provided path is invalid.
func MustParseChannelPath(path string) (string, string) { _ = "STUB: not implemented"; return "", "" }
