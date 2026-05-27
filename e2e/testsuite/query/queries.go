package query

import (
	"context"

	"github.com/cosmos/interchaintest/v11/ibc"

	"cosmossdk.io/math"

	"github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
	sdk "github.com/cosmos/cosmos-sdk/types"

	transfertypes "github.com/cosmos/ibc-go/v11/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
	ibcexported "github.com/cosmos/ibc-go/v11/modules/core/exported"
)

const queryPathTransferDenoms = "/ibc.applications.transfer.v1.Query/Denoms"

// ModuleAccountAddress returns the address of the given module on the given chain.
// Added because interchaintest's method doesn't work.
func ModuleAccountAddress(ctx context.Context, moduleName string, chain ibc.Chain) (sdk.AccAddress, error) {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress), nil
}

// ClientState queries the client state on the given chain for the provided clientID.
func ClientState(ctx context.Context, chain ibc.Chain, clientID string) (ibcexported.ClientState, error) {
	_ = "STUB: not implemented"
	return *new(ibcexported.ClientState), nil
}

// ClientStatus queries the status of the client by clientID
func ClientStatus(ctx context.Context, chain ibc.Chain, clientID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetValidatorSetByHeight returns the validators of the given chain at the specified height. The returned validators
// are sorted by address.
func GetValidatorSetByHeight(ctx context.Context, chain ibc.Chain, height uint64) ([]*cmtservice.Validator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Balance returns the balance of a specific denomination for a given account by address.
func Balance(ctx context.Context, chain ibc.Chain, address string, denom string) (math.Int, error) {
	_ = "STUB: not implemented"
	return *new(math.Int), nil
}

// Channel queries the channel on a given chain for the provided portID and channelID
func Channel(ctx context.Context, chain ibc.Chain, portID, channelID string) (channeltypes.Channel, error) {
	_ = "STUB: not implemented"
	return *new(channeltypes.Channel), nil
}

// TotalEscrowForDenom queries the total amount of tokens in escrow for a denom
func TotalEscrowForDenom(ctx context.Context, chain ibc.Chain, denom string) (sdk.Coin, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Coin), nil
}

// PacketAcknowledgements queries the packet acknowledgements on the given chain for the provided channel (optional) list of packet commitment sequences.
func PacketAcknowledgements(ctx context.Context, chain ibc.Chain, portID, channelID string, packetCommitmentSequences []uint64) ([]*channeltypes.PacketState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// InterchainAccount queries the interchain account for the given owner and connectionID.
func InterchainAccount(ctx context.Context, chain ibc.Chain, address, connectionID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func TransferDenoms(ctx context.Context, chain ibc.Chain) (*transfertypes.QueryDenomsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
