package cli

import (
	"github.com/cosmos/cosmos-sdk/client"

	"github.com/cosmos/ibc-go/v11/modules/core/04-channel/v2/types"
)

func queryNextSequenceSendABCI(clientCtx client.Context, channelID string) (*types.QueryNextSequenceSendResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check if next sequence send exists

func queryPacketCommitmentABCI(clientCtx client.Context, channelID string, sequence uint64) (*types.QueryPacketCommitmentResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check if packet commitment exists

func queryPacketAcknowledgementABCI(clientCtx client.Context, channelID string, sequence uint64) (*types.QueryPacketAcknowledgementResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check if packet commitment exists

func queryPacketReceiptABCI(clientCtx client.Context, channelID string, sequence uint64) (*types.QueryPacketReceiptResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
