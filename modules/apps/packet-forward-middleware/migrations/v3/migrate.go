package v3

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/apps/packet-forward-middleware/types"
	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
)

// Migrate migrates the x/packetforward module state from the consensus version
// 2 to version 3
func Migrate(ctx sdk.Context, bankKeeper types.BankKeeper, channelKeeper types.ChannelKeeper, transferKeeper types.TransferKeeper) error {
	_ = "STUB: not implemented"
	return nil
}

// 4. Set the total escrowed for each denom

func TotalEscrow(ctx sdk.Context, bankKeeper types.BankKeeper, channelKeeper types.ChannelKeeper, portID string) (sdk.Coins, []channeltypes.IdentifiedChannel) {
	_ = "STUB: not implemented"
	return *new(sdk.Coins), nil
}

// 1. Iterate over all IBC transfer channels

// 2. For each channel, get the escrow address and corresponding bank balance

// 3. Aggregate the bank balances to calculate the expected total escrowed
