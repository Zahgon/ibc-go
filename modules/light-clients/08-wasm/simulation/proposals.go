package simulation

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

// Simulation operation weights constants
const (
	DefaultWeightMsgStoreCode int = 100

	OpWeightMsgStoreCode = "op_weight_msg_store_code" // #nosec
)

// ProposalMsgs defines the module weighted proposals' contents
func ProposalMsgs() []simtypes.WeightedProposalMsg { _ = "STUB: not implemented"; return nil }

// SimulateMsgStoreCode returns a random MsgStoreCode for the 08-wasm module
func SimulateMsgStoreCode(r *rand.Rand, _ sdk.Context, _ []simtypes.Account) sdk.Msg {
	_ = "STUB: not implemented"
	return *new(sdk.Msg)
}
