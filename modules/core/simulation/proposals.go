package simulation

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

// Simulation operation weights constants
const (
	DefaultWeight int = 100

	OpWeightMsgUpdateParams       = "op_weight_msg_update_params"                 // #nosec
	OpWeightMsgRecoverClient      = "op_weight_msg_recover_client"                // #nosec
	OpWeightMsgIBCSoftwareUpgrade = "op_weight_msg_schedule_ibc_software_upgrade" // #nosec
)

// ProposalMsgs defines the module weighted proposals' contents
func ProposalMsgs() []simtypes.WeightedProposalMsg { _ = "STUB: not implemented"; return nil }

// SimulateClientMsgUpdateParams returns a MsgUpdateParams for 02-client
func SimulateClientMsgUpdateParams(r *rand.Rand, _ sdk.Context, _ []simtypes.Account) sdk.Msg {
	_ = "STUB: not implemented"
	return *new(sdk.Msg)
}

// SimulateClientMsgRecoverClient returns a MsgRecoverClient for 02-client
func SimulateClientMsgRecoverClient(r *rand.Rand, _ sdk.Context, _ []simtypes.Account) sdk.Msg {
	_ = "STUB: not implemented"
	return *new(sdk.Msg)
}

// SimulateClientMsgScheduleIBCSoftwareUpgrade returns a MsgScheduleIBCSoftwareUpgrade for 02-client
func SimulateClientMsgScheduleIBCSoftwareUpgrade(r *rand.Rand, _ sdk.Context, _ []simtypes.Account) sdk.Msg {
	_ = "STUB: not implemented"
	return *new(sdk.Msg)
}

// SimulateConnectionMsgUpdateParams returns a MsgUpdateParams 03-connection
func SimulateConnectionMsgUpdateParams(r *rand.Rand, _ sdk.Context, _ []simtypes.Account) sdk.Msg {
	_ = "STUB: not implemented"
	return *new(sdk.Msg)
}
