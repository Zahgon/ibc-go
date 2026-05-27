package simulation

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

// Simulation operation weights constants
const (
	DefaultWeightMsgUpdateParams int = 100

	OpWeightMsgUpdateParams = "op_weight_msg_update_params" // #nosec
)

// ProposalMsgs defines the module weighted proposals' contents
func ProposalMsgs() []simtypes.WeightedProposalMsg { _ = "STUB: not implemented"; return nil }

// SimulateMsgUpdateParams returns a MsgUpdateParams
func SimulateMsgUpdateParams(_ *rand.Rand, _ sdk.Context, _ []simtypes.Account) sdk.Msg {
	_ = "STUB: not implemented"
	return *new(sdk.Msg)
}
