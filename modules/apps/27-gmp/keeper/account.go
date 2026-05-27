package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/apps/27-gmp/types"
)

// getOrCreateICS27Account retrieves an existing ICS27 account or creates a new one if it doesn't exist.
func (k *Keeper) getOrCreateICS27Account(ctx context.Context, accountID *types.AccountIdentifier) (*types.ICS27Account, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create a new account

// GetOrComputeICS27Adderss retrieves an existing ICS27 account address or computes it if it doesn't exist. This doesn't modify the store.
func (k *Keeper) GetOrComputeICS27Address(ctx context.Context, accountID *types.AccountIdentifier) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Compute a new address

// GetAccount retrieves an existing ICS27 account based on the account address.
func (k *Keeper) GetAccount(ctx context.Context, address sdk.AccAddress) (*types.ICS27Account, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
