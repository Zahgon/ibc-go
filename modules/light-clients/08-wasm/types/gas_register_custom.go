package types

import (
	wasmvm "github.com/CosmWasm/wasmvm/v3"
	wasmvmtypes "github.com/CosmWasm/wasmvm/v3/types"

	storetypes "github.com/cosmos/cosmos-sdk/store/v2/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// While gas_register.go is a direct copy of https://github.com/CosmWasm/wasmd/blob/main/x/wasm/types/gas_register.go
// This file contains additional constructs that can be maintained separately.
// Most of these functions are slight modifications of keeper function from wasmd, which act on `WasmGasRegister` instead of `Keeper`.
const (
	// DefaultDeserializationCostPerByte The formula should be `len(data) * deserializationCostPerByte`
	DefaultDeserializationCostPerByte = 1
)

var CostJSONDeserialization = wasmvmtypes.UFraction{
	Numerator:   DefaultDeserializationCostPerByte * DefaultGasMultiplier,
	Denominator: 1,
}

func (g WasmGasRegister) RuntimeGasForContract(ctx sdk.Context) uint64 {
	_ = "STUB: not implemented"
	return 0
}

// infinite gas meter with limit=0 or MaxUint64

func (g WasmGasRegister) ConsumeRuntimeGas(ctx sdk.Context, gas uint64) {
	_ = "STUB: not implemented"
	return
}

// throw OutOfGas error if we ran out (got exactly to zero due to better limit enforcing)

// MultipliedGasMeter wraps the GasMeter from context and multiplies all reads by out defined multiplier
type MultipliedGasMeter struct {
	originalMeter storetypes.GasMeter
	GasRegister   GasRegister
}

func NewMultipliedGasMeter(originalMeter storetypes.GasMeter, gr GasRegister) MultipliedGasMeter {
	_ = "STUB: not implemented"
	return *new(MultipliedGasMeter)
}

var _ wasmvm.GasMeter = MultipliedGasMeter{}

func (m MultipliedGasMeter) GasConsumed() storetypes.Gas {
	_ = "STUB: not implemented"
	return *new(storetypes.Gas)
}
