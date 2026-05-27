package simapp

import (
	simappparams "github.com/cosmos/ibc-go/modules/light-clients/08-wasm/v11/testing/simapp/params"
)

// MakeTestEncodingConfig creates an EncodingConfig for testing. This function
// should be used only in tests or when creating a new app instance (NewApp*()).
// App user shouldn't create new codecs - use the app.AppCodec instead.
// [DEPRECATED]
func MakeTestEncodingConfig() simappparams.EncodingConfig {
	_ = "STUB: not implemented"
	return *new(simappparams.EncodingConfig)
}
