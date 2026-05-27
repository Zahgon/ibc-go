package gmp

import (
	"encoding/json"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	"cosmossdk.io/core/appmodule"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/cosmos/ibc-go/v11/modules/apps/27-gmp/keeper"
)

var (
	_ module.AppModule      = (*AppModule)(nil)
	_ module.AppModuleBasic = (*AppModule)(nil)
	// _ module.AppModuleSimulation = (*AppModule)(nil)
	_ module.HasGenesis          = (*AppModule)(nil)
	_ module.HasName             = (*AppModule)(nil)
	_ module.HasConsensusVersion = (*AppModule)(nil)
	_ module.HasServices         = (*AppModule)(nil)
	// _ module.HasProposalMsgs     = (*AppModule)(nil)
	_ appmodule.AppModule = (*AppModule)(nil)
)

// AppModule represents the AppModule for this module
type AppModule struct {
	keeper *keeper.Keeper
}

// NewAppModule creates a new 27-gmp module
func NewAppModule(k *keeper.Keeper) AppModule { _ = "STUB: not implemented"; return *new(AppModule) }

func NewAppModuleBasic(m AppModule) module.AppModuleBasic {
	_ = "STUB: not implemented"
	return *new(module.AppModuleBasic)
}

// Name implements AppModuleBasic interface
func (AppModule) Name() string { _ = "STUB: not implemented"; return "" }

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (AppModule) IsOnePerModuleType() {
	_ = "STUB: not implemented"

	// IsAppModule implements the appmodule.AppModule interface.
	return
}

func (AppModule) IsAppModule() {
	_ = "STUB: not implemented"

	// RegisterLegacyAminoCodec implements AppModuleBasic interface
	return
}

func (AppModule) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	_ = "STUB: not implemented"

	// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the ics27 module.
	return
}

func (AppModule) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	_ = "STUB: not implemented"
	return
}

// RegisterInterfaces registers module concrete types into protobuf Any.
func (AppModule) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	_ = "STUB: not implemented"
	return
}

// ConsensusVersion implements AppModule/ConsensusVersion defining the current version of gmp.
func (AppModule) ConsensusVersion() uint64 {
	_ = "STUB: not implemented"

	// DefaultGenesis returns default genesis state as raw bytes for the gmp module.
	return 0
}

func (AppModule) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

// RegisterServices registers module services.
func (am AppModule) RegisterServices(cfg module.Configurator) { _ = "STUB: not implemented"; return }

// ValidateGenesis performs genesis state validation for the ibc gmp module.
func (AppModule) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (AppModule) AutoCLIOptions() *autocliv1.ModuleOptions { _ = "STUB: not implemented"; return nil }

// InitGenesis performs genesis initialization for the gmp module. It returns no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) {
	_ = "STUB: not implemented"
	return
}

// ExportGenesis returns the exported genesis state as raw bytes for the gmp module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

/*
// AppModuleSimulation functions

// GenerateGenesisState creates a randomized GenState of the gmp module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	simulation.RandomizedGenState(simState)
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return simulation.ProposalMsgs()
}

// RegisterStoreDecoder registers a decoder for gmp module's types
func (AppModule) RegisterStoreDecoder(sdr simtypes.StoreDecoderRegistry) {
	sdr[types.StoreKey] = simulation.NewDecodeStore()
}

// WeightedOperations returns the all the gmp module operations with their respective weights.
func (AppModule) WeightedOperations(_ module.SimulationState) []simtypes.WeightedOperation {
	return nil
}
*/
