package mock

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"cosmossdk.io/core/appmodule"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	abci "github.com/cometbft/cometbft/abci/types"

	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v11/modules/core/05-port/types"
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

const (
	ModuleName = "mock"

	MemStoreKey = "memory:mock"

	PortID = ModuleName

	Version = "mock-version"
)

var (
	MockAcknowledgement     = channeltypes.NewResultAcknowledgement([]byte("mock acknowledgement"))
	MockFailAcknowledgement = channeltypes.NewErrorAcknowledgement(errors.New("mock failed acknowledgement"))
	MockPacketData          = []byte("mock packet data")
	MockFailPacketData      = []byte("mock failed packet data")
	MockAsyncPacketData     = []byte("mock async packet data")
	UpgradeVersion          = fmt.Sprintf("%s-v2", Version)
	// MockApplicationCallbackError should be returned when an application callback should fail. It is possible to
	// test that this error was returned using ErrorIs.
	MockApplicationCallbackError error = &applicationCallbackError{}
)

var (
	TestKey   = []byte("test-key")
	TestValue = []byte("test-value")
)

var (
	_ module.AppModuleBasic = (*AppModuleBasic)(nil)
	_ appmodule.AppModule   = (*AppModule)(nil)

	_ porttypes.IBCModule = (*IBCModule)(nil)
)

// AppModuleBasic is the mock AppModuleBasic.
type AppModuleBasic struct{}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (AppModuleBasic) IsOnePerModuleType() {
	_ = "STUB: not implemented"

	// IsAppModule implements the appmodule.AppModule interface.
	return
}

func (AppModuleBasic) IsAppModule() {
	_ = "STUB: not implemented"

	// Name implements AppModuleBasic interface.
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

	// RegisterLegacyAminoCodec implements AppModuleBasic interface.
	return
}

func (AppModuleBasic) RegisterLegacyAminoCodec(*codec.LegacyAmino) {
	_ = "STUB: not implemented"

	// RegisterInterfaces implements AppModuleBasic interface.
	return
}

func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	_ = "STUB: not implemented"

	// DefaultGenesis implements AppModuleBasic interface.
	return
}

func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	_ = "STUB: not implemented"

	// ValidateGenesis implements the AppModuleBasic interface.
	return *new(json.RawMessage)
}

func (AppModuleBasic) ValidateGenesis(codec.JSONCodec, client.TxEncodingConfig, json.RawMessage) error {
	_ = "STUB: not implemented"

	// RegisterGRPCGatewayRoutes implements AppModuleBasic interface.
	return nil
}

func (AppModuleBasic) RegisterGRPCGatewayRoutes(_ client.Context, _ *runtime.ServeMux) {
	_ = "STUB: not implemented"

	// GetTxCmd implements AppModuleBasic interface.
	return
}

func (AppModuleBasic) GetTxCmd() *cobra.Command {
	_ = "STUB: not implemented"

	// GetQueryCmd implements AppModuleBasic interface.
	return nil
}

func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	_ = "STUB: not implemented"

	// AppModule represents the AppModule for the mock module.
	return nil
}

type AppModule struct {
	AppModuleBasic
	ibcApps []*IBCApp
}

// NewAppModule returns a mock AppModule instance.
func NewAppModule() AppModule {
	_ = "STUB: not implemented"

	// RegisterInvariants implements the AppModule interface.
	return *new(AppModule)
}

func (AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {
	_ = "STUB: not implemented"

	// RegisterServices implements the AppModule interface.
	return
}

func (AppModule) RegisterServices(module.Configurator) {
	_ = "STUB: not implemented"

	// InitGenesis implements the AppModule interface.
	return
}

func (AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	_ = "STUB: not implemented"
	return nil
}

// ExportGenesis implements the AppModule interface.
func (AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	_ = "STUB: not implemented"

	// ConsensusVersion implements AppModule/ConsensusVersion.
	return *new(json.RawMessage)
}

func (AppModule) ConsensusVersion() uint64 { _ = "STUB: not implemented"; return 0 }

var _ exported.Path = KeyPath{}

// KeyPath defines a placeholder struct which implements the exported.Path interface
type KeyPath struct{}

// String implements the exported.Path interface
func (KeyPath) String() string {
	_ = "STUB: not implemented"

	// Empty implements the exported.Path interface
	return ""
}

func (KeyPath) Empty() bool { _ = "STUB: not implemented"; return false }

var _ exported.Height = Height{}

// Height defines a placeholder struct which implements the exported.Height interface
type Height struct {
	exported.Height
}
