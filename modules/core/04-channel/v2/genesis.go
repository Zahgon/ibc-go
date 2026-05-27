package channelv2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/core/04-channel/v2/keeper"
	"github.com/cosmos/ibc-go/v11/modules/core/04-channel/v2/types"
)

func InitGenesis(ctx sdk.Context, k *keeper.Keeper, gs types.GenesisState) {
	_ = "STUB: not implemented"
	// set acks
	return
}

// set commits

// set receipts

// set async packets

// set send sequences

func ExportGenesis(ctx sdk.Context, k *keeper.Keeper) types.GenesisState {
	_ = "STUB: not implemented"
	return *new(types.GenesisState)
}
