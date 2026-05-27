package simulation

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	controllerkeeper "github.com/cosmos/ibc-go/v11/modules/apps/27-interchain-accounts/controller/keeper"
	hostkeeper "github.com/cosmos/ibc-go/v11/modules/apps/27-interchain-accounts/host/keeper"
)

// Simulation operation weights constants
const (
	DefaultWeightMsgUpdateParams int = 100

	OpWeightMsgUpdateParams = "op_weight_msg_update_params" // #nosec
)

// ProposalMsgs defines the module weighted proposals' contents
func ProposalMsgs(controllerKeeper *controllerkeeper.Keeper, hostKeeper *hostkeeper.Keeper) []simtypes.WeightedProposalMsg {
	_ = "STUB: not implemented"
	return nil
}

// SimulateHostMsgUpdateParams returns a MsgUpdateParams for the host module
func SimulateHostMsgUpdateParams(_ *rand.Rand, _ sdk.Context, _ []simtypes.Account) sdk.Msg {
	_ = "STUB: not implemented"
	return *new(sdk.Msg)
}

// SimulateControllerMsgUpdateParams returns a MsgUpdateParams for the controller module
func SimulateControllerMsgUpdateParams(_ *rand.Rand, _ sdk.Context, _ []simtypes.Account) sdk.Msg {
	_ = "STUB: not implemented"
	return *new(sdk.Msg)
}
