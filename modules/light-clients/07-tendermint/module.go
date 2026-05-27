package tendermint

import (
	"encoding/json"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"cosmossdk.io/core/appmodule"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

var (
	_ module.AppModuleBasic = (*AppModuleBasic)(nil)
	_ appmodule.AppModule   = (*AppModule)(nil)
)

// AppModuleBasic defines the basic application module used by the tendermint light client.
// Only the RegisterInterfaces function needs to be implemented. All other function perform
// a no-op.
type AppModuleBasic struct{}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (AppModuleBasic) IsOnePerModuleType() {
	_ = "STUB: not implemented"

	// IsAppModule implements the appmodule.AppModule interface.
	return
}

func (AppModuleBasic) IsAppModule() {
	_ = "STUB: not implemented"

	// Name returns the tendermint module name.
	return
}

func (AppModuleBasic) Name() string {
	_ = "STUB: not implemented"

	// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
	return ""
}

func (AppModule) IsOnePerModuleType() {
	_ = "STUB: not implemented"

	// IsAppModule implements the appmodule.AppModule interface.
	return
}

func (AppModule) IsAppModule() {
	_ = "STUB: not implemented"

	// RegisterLegacyAminoCodec performs a no-op. The Tendermint client does not support amino.
	return
}

func (AppModuleBasic) RegisterLegacyAminoCodec(*codec.LegacyAmino) {
	_ = "STUB: not implemented"

	// RegisterInterfaces registers module concrete types into protobuf Any. This allows core IBC
	// to unmarshal tendermint light client types.
	return
}

func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	_ = "STUB: not implemented"
	return
}

// DefaultGenesis performs a no-op. Genesis is not supported for the tendermint light client.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	_ = "STUB: not implemented"

	// ValidateGenesis performs a no-op. Genesis is not supported for the tendermint light client.
	return *new(json.RawMessage)
}

func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	_ = "STUB: not implemented"

	// RegisterGRPCGatewayRoutes performs a no-op.
	return nil
}

func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	_ = "STUB: not implemented"

	// GetTxCmd performs a no-op. Please see the 02-client cli commands.
	return
}

func (AppModuleBasic) GetTxCmd() *cobra.Command {
	_ = "STUB: not implemented"

	// GetQueryCmd performs a no-op. Please see the 02-client cli commands.
	return nil
}

func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	_ = "STUB: not implemented"

	// AppModule is the application module for the Tendermint client module
	return nil
}

type AppModule struct {
	AppModuleBasic
	lightClientModule LightClientModule
}

// NewAppModule creates a new Tendermint client module
func NewAppModule(lightClientModule LightClientModule) AppModule {
	_ = "STUB: not implemented"
	return *new(AppModule)
}
