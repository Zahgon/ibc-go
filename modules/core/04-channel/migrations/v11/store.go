package v11

import (
	corestore "cosmossdk.io/core/store"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/core/keeper"
)

const (
	KeyNextSeqSendPrefix = "nextSequenceSend"
	KeyChannelEndPrefix  = "channelEnds"
	KeyChannelPrefix     = "channels"
	KeyPortPrefix        = "ports"
)

// NextSequenceSendV1Key returns the store key for the send sequence of a particular
// channel binded to a specific port.
func NextSequenceSendV1Key(portID, channelID string) []byte { _ = "STUB: not implemented"; return nil }

// MigrateStore migrates the channel store to add support for IBC v2
// for all OPEN UNORDERED channels by:
// - Adding client counterparty information keyed to the channel ID
// - Migrating the NextSequenceSend path to use the v2 format
// - Store an alias key mapping the v1 channel ID to the underlying client ID
func MigrateStore(ctx sdk.Context, storeService corestore.KVStoreService, cdc codec.BinaryCodec,
	ibcKeeper *keeper.Keeper,
) error {
	_ = "STUB: not implemented"
	return nil
}

// only add counterparty for channels that are OPEN and UNORDERED
// set a base client mapping from the channelId to the underlying base client

// migrate the NextSequenceSend key to the v2 format for every channel

// set the NextSequenceSend in the v2 keeper

// remove the old NextSequenceSend key
