package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// SubModuleName defines the interchain accounts host module name
	SubModuleName = "icahost"

	// StoreKey is the store key string for the interchain accounts host module
	StoreKey = SubModuleName

	// ParamsKey is the key to use for the storing params.
	ParamsKey = "params"

	// AllowAllHostMsgs holds the string key that allows all message types on interchain accounts host module
	AllowAllHostMsgs = "*"
)

var (
	// KeyHostEnabled is the store key for HostEnabled Params
	KeyHostEnabled = []byte("HostEnabled")
	// KeyAllowMessages is the store key for the AllowMessages Params
	KeyAllowMessages = []byte("AllowMessages")
)

// ContainsMsgType returns true if the sdk.Msg TypeURL is present in allowMsgs, otherwise false
func ContainsMsgType(allowMsgs []string, msg sdk.Msg) bool {
	_ = "STUB: not implemented"
	// check that wildcard * option for allowing all message types is the only string in the array, if so, return true
	return false
}
