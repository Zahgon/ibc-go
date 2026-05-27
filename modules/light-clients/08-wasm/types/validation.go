package types

// MaxWasmSize denotes the maximum size (in bytes) a contract is allowed to be.
const MaxWasmSize uint64 = 3 * 1024 * 1024

// ValidateWasmCode valides that the size of the wasm code is in the allowed range
// and that the contents are of a wasm binary.
func ValidateWasmCode(code []byte) error { _ = "STUB: not implemented"; return nil }

// ValidateWasmChecksum validates that the checksum is of the correct length
func ValidateWasmChecksum(checksum Checksum) error { _ = "STUB: not implemented"; return nil }

// sha256 output is 256 bits long

// ValidateClientID validates the client identifier by ensuring that it conforms
// to the 02-client identifier format and that it is a 08-wasm clientID.
func ValidateClientID(clientID string) error { _ = "STUB: not implemented"; return nil }
