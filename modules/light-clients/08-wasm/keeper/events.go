package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/modules/light-clients/08-wasm/v11/types"
)

// emitStoreWasmCodeEvent emits a store wasm code event
func emitStoreWasmCodeEvent(ctx sdk.Context, checksum types.Checksum) {
	_ = "STUB: not implemented"
	return
}

// emitMigrateContractEvent emits a migrate contract event
func emitMigrateContractEvent(ctx sdk.Context, clientID string, checksum, newChecksum types.Checksum) {
	_ = "STUB: not implemented"
	return
}
