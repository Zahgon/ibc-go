package keeper

import (
	corestore "cosmossdk.io/core/store"

	"github.com/cosmos/cosmos-sdk/codec"

	clientkeeper "github.com/cosmos/ibc-go/v11/modules/core/02-client/keeper"
	clienttypes "github.com/cosmos/ibc-go/v11/modules/core/02-client/types"
	clientv2keeper "github.com/cosmos/ibc-go/v11/modules/core/02-client/v2/keeper"
	connectionkeeper "github.com/cosmos/ibc-go/v11/modules/core/03-connection/keeper"
	channelkeeper "github.com/cosmos/ibc-go/v11/modules/core/04-channel/keeper"
	channelkeeperv2 "github.com/cosmos/ibc-go/v11/modules/core/04-channel/v2/keeper"
	portkeeper "github.com/cosmos/ibc-go/v11/modules/core/05-port/keeper"
	porttypes "github.com/cosmos/ibc-go/v11/modules/core/05-port/types"
	"github.com/cosmos/ibc-go/v11/modules/core/api"
)

// Keeper defines each ICS keeper for IBC
type Keeper struct {
	ClientKeeper     *clientkeeper.Keeper
	ClientV2Keeper   *clientv2keeper.Keeper
	ConnectionKeeper *connectionkeeper.Keeper
	ChannelKeeper    *channelkeeper.Keeper
	ChannelKeeperV2  *channelkeeperv2.Keeper
	PortKeeper       *portkeeper.Keeper

	cdc codec.BinaryCodec

	authority string
}

// NewKeeper creates a new ibc Keeper
func NewKeeper(
	cdc codec.BinaryCodec, storeService corestore.KVStoreService,
	upgradeKeeper clienttypes.UpgradeKeeper, authority string,
) *Keeper {
	_ = "STUB: not implemented"
	// panic if any of the keepers passed in is empty
	return nil
}

// Codec returns the IBC module codec.
func (k *Keeper) Codec() codec.BinaryCodec {
	_ = "STUB: not implemented"

	// SetRouter sets the Router in IBC Keeper and seals it. The method panics if
	// there is an existing router that's already sealed.
	return *new(codec.BinaryCodec)
}

func (k *Keeper) SetRouter(rtr *porttypes.Router) { _ = "STUB: not implemented"; return }

// SetRouterV2 sets the v2 router for the IBC Keeper.
func (k *Keeper) SetRouterV2(rtr *api.Router) { _ = "STUB: not implemented"; return }

// GetAuthority returns the ibc module's authority.
func (k *Keeper) GetAuthority() string {
	_ = "STUB: not implemented"

	// isEmpty checks if the interface is an empty struct or a pointer pointing
	// to an empty struct
	return ""
}

func isEmpty(keeper any) bool { _ = "STUB: not implemented"; return false }
