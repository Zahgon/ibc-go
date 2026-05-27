package simapp

import (
	"encoding/json"
	"os"
	"path/filepath"

	dbm "github.com/cosmos/cosmos-db"

	"cosmossdk.io/client/v2/autocli"
	"cosmossdk.io/log/v2"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/v2/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	consensusparamkeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"

	abci "github.com/cometbft/cometbft/abci/types"

	gmpkeeper "github.com/cosmos/ibc-go/v11/modules/apps/27-gmp/keeper"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v11/modules/apps/27-interchain-accounts/controller/keeper"
	icahostkeeper "github.com/cosmos/ibc-go/v11/modules/apps/27-interchain-accounts/host/keeper"
	icatypes "github.com/cosmos/ibc-go/v11/modules/apps/27-interchain-accounts/types"
	packetforwardkeeper "github.com/cosmos/ibc-go/v11/modules/apps/packet-forward-middleware/keeper"
	ratelimitkeeper "github.com/cosmos/ibc-go/v11/modules/apps/rate-limiting/keeper"
	ibctransferkeeper "github.com/cosmos/ibc-go/v11/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v11/modules/apps/transfer/types"
	ibckeeper "github.com/cosmos/ibc-go/v11/modules/core/keeper"
)

const appName = "SimApp"

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:     nil,
		distrtypes.ModuleName:          nil,
		minttypes.ModuleName:           {authtypes.Minter},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:            {authtypes.Burner},
		ibctransfertypes.ModuleName:    {authtypes.Minter, authtypes.Burner},
		icatypes.ModuleName:            nil,
	}
)

var (
	_ runtime.AppI            = (*SimApp)(nil)
	_ servertypes.Application = (*SimApp)(nil)
)

// SimApp extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions.
type SimApp struct {
	*baseapp.BaseApp
	legacyAmino       *codec.LegacyAmino
	appCodec          codec.Codec
	txConfig          client.TxConfig
	interfaceRegistry types.InterfaceRegistry

	// keys to access the substores
	keys map[string]*storetypes.KVStoreKey

	// keepers
	AccountKeeper         authkeeper.AccountKeeper
	BankKeeper            bankkeeper.Keeper
	StakingKeeper         *stakingkeeper.Keeper
	SlashingKeeper        slashingkeeper.Keeper
	MintKeeper            mintkeeper.Keeper
	DistrKeeper           distrkeeper.Keeper
	GovKeeper             govkeeper.Keeper
	UpgradeKeeper         *upgradekeeper.Keeper
	AuthzKeeper           authzkeeper.Keeper
	IBCKeeper             *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	ICAControllerKeeper   *icacontrollerkeeper.Keeper
	ICAHostKeeper         *icahostkeeper.Keeper
	EvidenceKeeper        evidencekeeper.Keeper
	TransferKeeper        *ibctransferkeeper.Keeper
	FeeGrantKeeper        feegrantkeeper.Keeper
	ConsensusParamsKeeper consensusparamkeeper.Keeper
	PFMKeeper             *packetforwardkeeper.Keeper
	RateLimitKeeper       *ratelimitkeeper.Keeper
	GMPKeeper             *gmpkeeper.Keeper

	// the module manager
	ModuleManager      *module.Manager
	BasicModuleManager module.BasicManager

	// simulation manager
	simulationManager *module.SimulationManager

	// module configurator
	configurator module.Configurator
}

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, ".simapp")
}

// NewSimApp returns a reference to an initialized SimApp.
func NewSimApp(
	logger log.Logger,
	db dbm.DB,
	loadLatest bool,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *SimApp {
	_ = "STUB: not implemented"
	return nil
}

// Below we could construct and set an application specific mempool and
// ABCI 1.0 PrepareProposal and ProcessProposal handlers. These defaults are
// already set in the SDK's BaseApp, this shows an example of how to override
// them.
//
// Example:
//
// bApp := baseapp.NewBaseApp(...)
// nonceMempool := mempool.NewSenderNonceMempool()
// abciPropHandler := NewDefaultProposalHandler(nonceMempool, bApp)
//
// bApp.SetMempool(nonceMempool)
// bApp.SetPrepareProposal(abciPropHandler.PrepareProposalHandler())
// bApp.SetProcessProposal(abciPropHandler.ProcessProposalHandler())
//
// Alternatively, you can construct BaseApp options, append those to
// baseAppOptions and pass them to NewBaseApp.
//
// Example:
//
// prepareOpt = func(app *baseapp.BaseApp) {
// 	abciPropHandler := baseapp.NewDefaultProposalHandler(nonceMempool, app)
// 	app.SetPrepareProposal(abciPropHandler.PrepareProposalHandler())
// }
// baseAppOptions = append(baseAppOptions, prepareOpt)

// register streaming services

// set the BaseApp's parameter store

// SDK module keepers

// add keepers

// register the staking hooks
// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks

// get skipUpgradeHeights from the app options

// set the governance module account as the authority for conducting upgrades

/*
	Example of setting gov params:
	govConfig.MaxMetadataLen = 10000
*/

// register the governance hooks

// ICA Controller keeper

// ICA Host keeper

// GMP Keeper

// Transfer Keeper

// Packet Forward Middleware keeper

// Create IBC Router

// Create Transfer Stack
// SendPacket, since it is originating from the application to core IBC:
// transferKeeper.SendPacket -> Pf.SendPacket -> RateLim.SendPacket -> channel.SendPacket

// RecvPacket, message that originates from core IBC and goes down to app, the flow is the other way
// channel.RecvPacket -> RateLim.OnRecvPacket -> Pf.OnRecvPacket -> transfer.OnRecvPacket

// transfer stack contains (from top to bottom):
// - RateLimit
// - PacketForward
// - Transfer

// create IBC module from bottom to top of stack

// Add transfer stack to IBC Router

// Packet Forward Middleware Stack.

// Create Interchain Accounts Stack
// SendPacket, since it is originating from the application to core IBC:
// icaControllerKeeper.SendTx -> channel.SendPacket

// RecvPacket, message that originates from core IBC and goes down to app, the flow is:
// channel.RecvPacket -> icaHost.OnRecvPacket

// Add host, controller & ica auth modules to IBC router

// register the transfer v2 module.

// register the gmp module.

// Set the IBC Routers

// create evidence keeper with router

// If evidence needs to be handled for the app, set routes in router here and seal

// ****  Module Options ****

// NOTE: Any module instantiated in the module manager that is later modified
// must be passed by reference here.

// IBC modules

// IBC light clients

// BasicModuleManager defines the module BasicManager is in charge of setting up basic,
// non-dependant module elements, such as codec registration and genesis verification.
// By default it is composed of all the module from the module manager.
// Additionally, app module basics can be overwritten by passing them as argument.

// NOTE: upgrade module is required to be prioritized

// During begin block slashing happens after distr.BeginBlocker so that
// there is nothing left over in the validator fee pool, so as to keep the
// CanWithdrawInvariant invariant.
// NOTE: staking module is required if HistoricalEntries param > 0

// NOTE: The genutils module must occur after staking so that pools are
// properly initialized with tokens from genesis accounts.
// NOTE: The genutils module must also occur after auth so that it can access the params from auth.

// Uncomment if you want to set a custom migration order here.
// app.ModuleManager.SetOrderMigrations(custom order)

// registerUpgradeHandlers is used for registering any on-chain upgrades.
// Make sure it's called after `app.ModuleManager` and `app.configurator` are set.

// add test gRPC service for testing gRPC queries in isolation

// create the simulation manager and define the order of the modules for deterministic simulations
//
// NOTE: this is not required apps that don't use the simulator for fuzz testing
// transactions

// initialize stores

// initialize BaseApp

// In v0.46, the SDK introduces _postHandlers_. PostHandlers are like
// antehandlers, but are run _after_ the `runMsgs` execution. They are also
// defined as a chain, and have the same signature as antehandlers.
//
// In baseapp, postHandlers are run in the same store branch as `runMsgs`,
// meaning that both `runMsgs` and `postHandler` state will be committed if
// both are successful, and both will be reverted if any of the two fails.
//
// The SDK exposes a default postHandlers chain, which is comprised of only
// one decorator: the Transaction Tips decorator. However, some chains do
// not need it by default, so feel free to comment the next line if you do
// not need tips.
// To read more about tips:
// https://docs.cosmos.network/main/core/tips.html
//
// Please note that changing any of the anteHandler or postHandler chain is
// likely to be a state-machine breaking change, which needs a coordinated
// upgrade.

// At startup, after all modules have been registered, check that all proto
// annotations are correct.

// Once we switch to using protoreflect-based antehandlers, we might
// want to panic here instead of logging a warning.

func (app *SimApp) setAnteHandler(txConfig client.TxConfig) { _ = "STUB: not implemented"; return }

// Set the AnteHandler for the app

func (app *SimApp) setPostHandler() { _ = "STUB: not implemented"; return }

// Name returns the name of the App
func (app *SimApp) Name() string { _ = "STUB: not implemented"; return "" }

// PreBlocker application updates every pre block
func (app *SimApp) PreBlocker(ctx sdk.Context, _ *abci.RequestFinalizeBlock) (*sdk.ResponsePreBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BeginBlocker application updates every begin block
func (app *SimApp) BeginBlocker(ctx sdk.Context) (sdk.BeginBlock, error) {
	_ = "STUB: not implemented"
	return *new(sdk.BeginBlock), nil
}

// EndBlocker application updates every end block
func (app *SimApp) EndBlocker(ctx sdk.Context) (sdk.EndBlock, error) {
	_ = "STUB: not implemented"
	return *new(sdk.EndBlock), nil
}

// Configurator returns the configurator for the app
func (app *SimApp) Configurator() module.Configurator {
	_ = "STUB: not implemented"
	return *

	// InitChainer application update at chain initialization
	new(module.Configurator)
}

func (app *SimApp) InitChainer(ctx sdk.Context, req *abci.RequestInitChain) (*abci.ResponseInitChain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LoadHeight loads a particular height
func (app *SimApp) LoadHeight(height int64) error { _ = "STUB: not implemented"; return nil }

// LegacyAmino returns SimApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *SimApp) LegacyAmino() *codec.LegacyAmino { _ = "STUB: not implemented"; return nil }

// AppCodec returns SimApp's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *SimApp) AppCodec() codec.Codec {
	_ = "STUB: not implemented"
	return *

	// InterfaceRegistry returns SimApp's InterfaceRegistry
	new(codec.Codec)
}

func (app *SimApp) InterfaceRegistry() types.InterfaceRegistry {
	_ = "STUB: not implemented"
	return *new(types.InterfaceRegistry)
}

// TxConfig returns SimApp's TxConfig
func (app *SimApp) TxConfig() client.TxConfig {
	_ = "STUB: not implemented"
	return *

	// AutoCliOpts returns the autocli options for the app.
	new(client.TxConfig)
}

func (app *SimApp) AutoCliOpts() autocli.AppOptions {
	_ = "STUB: not implemented"
	return *new(autocli.AppOptions)
}

// DefaultGenesis returns a default genesis from the registered AppModuleBasic's.
func (app *SimApp) DefaultGenesis() map[string]json.RawMessage {
	_ = "STUB: not implemented"
	return nil
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *SimApp) GetKey(storeKey string) *storetypes.KVStoreKey {
	_ = "STUB: not implemented"
	return nil

	// GetStoreKeys returns all the stored store keys.
}

func (app *SimApp) GetStoreKeys() []storetypes.StoreKey { _ = "STUB: not implemented"; return nil }

// SimulationManager implements the SimulationApp interface
func (app *SimApp) SimulationManager() *module.SimulationManager {
	_ = "STUB: not implemented"
	return nil
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *SimApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	_ = "STUB: not implemented"
	return
}

// Register new tx routes from grpc-gateway.

// Register new CometBFT queries routes from grpc-gateway.

// Register node gRPC service for grpc-gateway.

// Register grpc-gateway routes for all modules.

// register swagger API from root so that other applications can override easily

// RegisterTxService implements the Application.RegisterTxService method.
func (app *SimApp) RegisterTxService(clientCtx client.Context) { _ = "STUB: not implemented"; return }

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *SimApp) RegisterTendermintService(clientCtx client.Context) {
	_ = "STUB: not implemented"
	return
}

func (app *SimApp) RegisterNodeService(clientCtx client.Context, cfg config.Config) {
	_ = "STUB: not implemented"
	return
}

// GetMaccPerms returns a copy of the module account permissions
//
// NOTE: This is solely to be used for testing purposes.
func GetMaccPerms() map[string][]string { _ = "STUB: not implemented"; return nil }

// BlockedAddresses returns all the app's blocked account addresses.
func BlockedAddresses() map[string]bool { _ = "STUB: not implemented"; return nil }

// allow the following addresses to receive funds
