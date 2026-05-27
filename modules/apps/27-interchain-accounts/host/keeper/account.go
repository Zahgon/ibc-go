package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// createInterchainAccount creates a new interchain account. An address is generated using the host connectionID, the controller portID,
// and block dependent information. An error is returned if an account already exists for the generated account.
// An interchain account type is set in the account keeper and the interchain account address mapping is updated.
func (k *Keeper) createInterchainAccount(ctx sdk.Context, connectionID, controllerPortID string) (sdk.AccAddress,
	error,
) {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress), nil
}
