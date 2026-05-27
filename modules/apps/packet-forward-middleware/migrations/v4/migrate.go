package v4

import (
	corestore "cosmossdk.io/core/store"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Migrate migrates the x/packetforward module state from consensus version 3 to version 4.
// It removes the deprecated nonrefundable field from stored in-flight packets and aborts if
// any packet has nonrefundable=true.
func Migrate(ctx sdk.Context, storeService corestore.KVStoreService, cdc codec.BinaryCodec) error {
	_ = "STUB: not implemented"
	return nil
}
