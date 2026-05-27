package transfer

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

	"github.com/cosmos/ibc-go/v11/modules/apps/transfer/keeper"
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

	_ porttypes.IBCModule = (*IBCModule)(nil)
)

// AppModuleBasic is the IBC Transfer AppModuleBasic
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

	// RegisterLegacyAminoCodec implements AppModuleBasic interface
	return
}

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	_ = "STUB: not implemented"
	return
}

// RegisterInterfaces registers module concrete types into protobuf Any.
func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	_ = "STUB: not implemented"
	return
}

// DefaultGenesis returns default genesis state as raw bytes for the ibc
// transfer module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

// ValidateGenesis performs genesis state validation for the ibc transfer module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the ibc-transfer module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	_ = "STUB: not implemented"
	return
}

// GetTxCmd implements AppModuleBasic interface
func (AppModuleBasic) GetTxCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetQueryCmd implements AppModuleBasic interface
func (AppModuleBasic) GetQueryCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// AppModule represents the AppModule for this module
type AppModule struct {
	AppModuleBasic
	keeper *keeper.Keeper
}

// NewAppModule creates a new 20-transfer module
func NewAppModule(k *keeper.Keeper) AppModule { _ = "STUB: not implemented"; return *new(AppModule) }

// RegisterServices registers module services.
func (am AppModule) RegisterServices(cfg module.Configurator) { _ = "STUB: not implemented"; return }

// InitGenesis performs genesis initialization for the ibc-transfer module. It returns
// no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) {
	_ = "STUB: not implemented"
	return
}

// ExportGenesis returns the exported genesis state as raw bytes for the ibc-transfer
// module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

// ConsensusVersion implements AppModule/ConsensusVersion defining the current version of transfer.
func (AppModule) ConsensusVersion() uint64 {
	_ = "STUB: not implemented"

	// AppModuleSimulation functions
	return 0
}

// GenerateGenesisState creates a randomized GenState of the transfer module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	_ = "STUB: not implemented"
	return
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	_ = "STUB: not implemented"
	return nil
}

// RegisterStoreDecoder registers a decoder for transfer module's types
func (AppModule) RegisterStoreDecoder(sdr simtypes.StoreDecoderRegistry) {
	_ = "STUB: not implemented"
	return
}

// WeightedOperations returns the all the transfer module operations with their respective weights.
func (AppModule) WeightedOperations(_ module.SimulationState) []simtypes.WeightedOperation {
	_ = "STUB: not implemented"
	return nil
}
