package ratelimiting

import (
	"context"
	"encoding/json"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"cosmossdk.io/core/appmodule"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/cosmos/ibc-go/v11/modules/apps/rate-limiting/keeper"
)

var (
	_ module.AppModule           = (*AppModule)(nil)
	_ module.AppModuleBasic      = (*AppModuleBasic)(nil)
	_ module.HasGenesis          = (*AppModule)(nil)
	_ module.HasName             = (*AppModule)(nil)
	_ module.HasConsensusVersion = (*AppModule)(nil)
	_ module.HasServices         = (*AppModule)(nil)
	_ appmodule.AppModule        = (*AppModule)(nil)
	_ appmodule.HasBeginBlocker  = (*AppModule)(nil)

	// Note: IBCMiddleware implements porttypes.Middleware and porttypes.ICS4Wrapper
)

// AppModuleBasic is the rate-limiting AppModuleBasic
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

// DefaultGenesis returns default genesis state as raw bytes for the rate-limiting
// module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

// ValidateGenesis performs genesis state validation for the rate-limiting module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the rate-limiting module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	_ = "STUB: not implemented"
	return
}

// GetTxCmd implements AppModuleBasic interface
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	_ = "STUB: not implemented"

	// GetQueryCmd implements AppModuleBasic interface
	return nil
}

func (AppModuleBasic) GetQueryCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// AppModule represents the AppModule for this module
type AppModule struct {
	AppModuleBasic
	keeper *keeper.Keeper
}

// NewAppModule creates a new rate-limiting module
func NewAppModule(k *keeper.Keeper) AppModule { _ = "STUB: not implemented"; return *new(AppModule) }

// RegisterServices registers module services.
func (am AppModule) RegisterServices(cfg module.Configurator) { _ = "STUB: not implemented"; return }

// Use the msgServer implementation

// InitGenesis performs genesis initialization for the rate-limiting module. It returns
// no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) {
	_ = "STUB: not implemented"
	return
}

// ExportGenesis returns the exported genesis state as raw bytes for the rate-limiting
// module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

// ConsensusVersion implements AppModule/ConsensusVersion defining the current version of rate-limiting.
func (AppModule) ConsensusVersion() uint64 {
	_ = "STUB: not implemented"

	// BeginBlock implements the AppModule interface
	return 0
}

func (am AppModule) BeginBlock(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// we do not want to raise an error in block processing if rate limit reset fails
