package migrations

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// PruneExpiredConsensusStates prunes all expired tendermint consensus states. This function
// may optionally be called during in-place store migrations. The ibc store key must be provided.
func PruneExpiredConsensusStates(ctx sdk.Context, cdc codec.BinaryCodec, clientKeeper ClientKeeper) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// keep track of the total consensus states pruned so chains can
// understand how much space is saved when the migration is run
