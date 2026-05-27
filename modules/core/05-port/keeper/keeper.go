package keeper

import (
	"cosmossdk.io/log/v2"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/core/05-port/types"
	"github.com/cosmos/ibc-go/v11/modules/core/api"
)

// Keeper defines the IBC connection keeper
type Keeper struct {
	Router   *types.Router
	RouterV2 *api.Router
}

// NewKeeper creates a new IBC connection Keeper instance
func NewKeeper() *Keeper {
	_ = "STUB: not implemented"

	// Logger returns a module-specific logger.
	return nil
}

func (*Keeper) Logger(ctx sdk.Context) log.Logger {
	_ = "STUB: not implemented"
	return *new(log.Logger)
}

// Route returns a IBCModule for a given module, and a boolean indicating
// whether or not the route is present.
func (k *Keeper) Route(module string) (types.IBCModule, bool) {
	_ = "STUB: not implemented"
	return *new(types.IBCModule), false
}
