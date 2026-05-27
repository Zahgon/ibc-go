//go:build !test_amino

package params

// MakeTestEncodingConfig creates an EncodingConfig for a non-amino based test configuration.
// This function should be used only internally (in the SDK).
// App user shouldn't create new codecs - use the app.AppCodec instead.
// [DEPRECATED]
func MakeTestEncodingConfig() EncodingConfig {
	_ = "STUB: not implemented"
	return *new(EncodingConfig)
}
