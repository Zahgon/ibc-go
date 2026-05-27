package types

import (
	"io"
)

// Copied gzip feature from wasmd
// https://github.com/CosmWasm/wasmd/blob/v0.31.0/x/wasm/ioutils/utils.go

// Note: []byte can never be const as they are inherently mutable

// magic bytes to identify gzip.
// See https://www.ietf.org/rfc/rfc1952.txt
// and https://github.com/golang/go/blob/master/src/net/http/sniff.go#L186
var gzipIdent = []byte("\x1F\x8B\x08")

// IsGzip returns checks if the file contents are gzip compressed
func IsGzip(input []byte) bool { _ = "STUB: not implemented"; return false }

// Uncompress expects a valid gzip source to unpack or fails. See IsGzip
func Uncompress(gzipSrc []byte, limit uint64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// limitReader returns a Reader that reads from r
// but stops with types.ErrLimit after n bytes.
// The underlying implementation is a *io.LimitedReader.
func limitReader(r io.Reader, n int64) io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

type limitedReader struct {
	r *io.LimitedReader
}

func (l *limitedReader) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// GzipIt compresses the input ([]byte)
func GzipIt(input []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	// Create gzip writer.
	return nil, nil
}

// You must close this first to flush the bytes to the buffer.
