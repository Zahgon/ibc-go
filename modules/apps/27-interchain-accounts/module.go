package ica

import (
	"encoding/json"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"cosmossdk.io/core/appmodule"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	controllerkeeper "github.com/cosmos/ibc-go/v11/modules/apps/27-interchain-accounts/controller/keeper"
	"github.com/cosmos/ibc-go/v11/modules/apps/27-interchain-accounts/host"
	hostkeeper "github.com/cosmos/ibc-go/v11/modules/apps/27-interchain-accounts/host/keeper"
	porttypes "github.com/cosmos/ibc-go/v11/modules/core/05-port/types"
)

var (
	_ module.AppModule           = (*AppModule)(nil)
	_ module.AppModuleBasic      = (*AppModuleBasic)(nil)
	_ module.AppModuleSimulation = (*AppModule)(nil)
	_ module.HasGenesis          = (*AppModule)(nil)
	_ module.HasName             = (*AppModule)(nil)
	_ module.HasConsensusVersion = (*AppModule)(nil)
	_ module.HasServices         = (*AppModule)(nil)
	_ module.HasProposalMsgs     = (*AppModule)(nil)
	_ appmodule.AppModule        = (*AppModule)(nil)

	_ porttypes.IBCModule = (*host.IBCModule)(nil)
)

// AppModuleBasic is the IBC interchain accounts AppModuleBasic
type AppModuleBasic struct{}

// Name implements AppModuleBasic interface
func (AppModuleBasic) Name() string { _ = "STUB: not implemented"; return "" }

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (AppModule) IsOnePerModuleType() {
	_ = "STUB: not implemented"

	// IsAppModule implements the appmodule.AppModule interface.
	return
}

func (AppModule) IsAppModule() {
	_ = "STUB: not implemented"

	// RegisterLegacyAminoCodec implements AppModuleBasic.
	return
}

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	_ = "STUB: not implemented"

	// RegisterInterfaces registers module concrete types into protobuf Any
	return
}

func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	_ = "STUB: not implemented"
	return
}

// DefaultGenesis returns default genesis state as raw bytes for the IBC
// interchain accounts module
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

// ValidateGenesis performs genesis state validation for the IBC interchain accounts module
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the interchain accounts module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	_ = "STUB: not implemented"
	return
}

// GetTxCmd implements AppModuleBasic interface
func (AppModuleBasic) GetTxCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetQueryCmd implements AppModuleBasic interface
func (AppModuleBasic) GetQueryCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// AppModule is the application module for the IBC interchain accounts module
type AppModule struct {
	AppModuleBasic
	controllerKeeper *controllerkeeper.Keeper
	hostKeeper       *hostkeeper.Keeper
}

// NewAppModule creates a new IBC interchain accounts module
func NewAppModule(controllerKeeper *controllerkeeper.Keeper, hostKeeper *hostkeeper.Keeper) AppModule {
	_ = "STUB: not implemented"
	return *new(AppModule)
}

// RegisterServices registers module services
func (am AppModule) RegisterServices(cfg module.Configurator) { _ = "STUB: not implemented"; return }

// InitGenesis performs genesis initialization for the interchain accounts module.
// It returns no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) {
	_ = "STUB: not implemented"
	return
}

// ExportGenesis returns the exported genesis state as raw bytes for the interchain accounts module
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 {
	_ = "STUB: not implemented"

	// AppModuleSimulation functions
	return 0
}

// GenerateGenesisState creates a randomized GenState of the ics27 module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	_ = "STUB: not implemented"
	return
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	_ = "STUB: not implemented"
	return nil
}

// WeightedOperations is unimplemented.
func (AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	_ = "STUB: not implemented"

	// RegisterStoreDecoder registers a decoder for interchain accounts module's types
	return nil
}

func (AppModule) RegisterStoreDecoder(sdr simtypes.StoreDecoderRegistry) {
	_ = "STUB: not implemented"
	return
}
