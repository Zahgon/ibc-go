package simapp

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

const (
	IBCWasmUpgrade = "ibcwasm-v8"
)

// registerUpgradeHandlers registers all supported upgrade handlers
func (app *SimApp) registerUpgradeHandlers() { _ = "STUB: not implemented"; return }

// configure store loader that checks if version == upgradeHeight and applies store upgrades

// createWasmStoreUpgradeHandler creates an upgrade handler for the 08-wasm ibc-go/v8 SimApp upgrade.
func createWasmStoreUpgradeHandler(mm *module.Manager, configurator module.Configurator) upgradetypes.UpgradeHandler {
	_ = "STUB: not implemented"
	return *new(upgradetypes.UpgradeHandler)
}
