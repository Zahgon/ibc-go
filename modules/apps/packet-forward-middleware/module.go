package packetforward

import (
	"encoding/json"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/cosmos/ibc-go/v11/modules/apps/packet-forward-middleware/keeper"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic is the packetforward AppModuleBasic
type AppModuleBasic struct{}

// Name implements AppModuleBasic interface
func (AppModuleBasic) Name() string { _ = "STUB: not implemented"; return "" }

// RegisterLegacyAminoCodec implements AppModuleBasic interface
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	_ = "STUB: not implemented"
	return
}

// RegisterInterfaces registers module concrete types into protobuf Any.
func (AppModuleBasic) RegisterInterfaces(r codectypes.InterfaceRegistry) {
	_ = "STUB: not implemented"
	return
}

// DefaultGenesis returns default genesis state as raw bytes for the ibc
// packetforward module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

// ValidateGenesis performs genesis state validation for the packetforward module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the packetforward module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	_ = "STUB: not implemented"

	// GetTxCmd implements AppModuleBasic interface
	return
}

func (AppModuleBasic) GetTxCmd() *cobra.Command {
	_ = "STUB: not implemented"

	// GetQueryCmd implements AppModuleBasic interface
	return nil
}

func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	_ = "STUB: not implemented"

	// AppModule represents the AppModule for this module
	return nil
}

type AppModule struct {
	AppModuleBasic
	keeper *keeper.Keeper
}

// NewAppModule creates a new packetforward module
func NewAppModule(k *keeper.Keeper) AppModule { _ = "STUB: not implemented"; return *new(AppModule) }

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (AppModule) IsOnePerModuleType() {
	_ = "STUB: not implemented"

	// IsAppModule implements the appmodule.AppModule interface.
	return
}

func (AppModule) IsAppModule() {
	_ = "STUB: not implemented"

	// QuerierRoute implements the AppModule interface
	return
}

func (AppModule) QuerierRoute() string { _ = "STUB: not implemented"; return "" }

// RegisterServices registers module services.
func (am AppModule) RegisterServices(cfg module.Configurator) { _ = "STUB: not implemented"; return }

// InitGenesis performs genesis initialization for the packetforward module. It returns
// no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	_ = "STUB: not implemented"
	return nil
}

// ExportGenesis returns the exported genesis state as raw bytes for the packetforward
// module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { _ = "STUB: not implemented"; return 0 }
