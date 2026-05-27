package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

// CreateClient generates a new client identifier and invokes the associated light client module in order to
// initialize a new client. An isolated prefixed store will be reserved for this client using the generated
// client identifier. The light client module is responsible for setting any client-specific data in the store
// via the Initialize method. This includes the client state, initial consensus state and any associated
// metadata. The generated client identifier will be returned if a client was successfully initialized.
func (k *Keeper) CreateClient(ctx sdk.Context, clientType string, clientState, consensusState []byte) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// UpdateClient updates the consensus state and the state root from a provided header.
func (k *Keeper) UpdateClient(ctx sdk.Context, clientID string, clientMsg exported.ClientMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// UpgradeClient upgrades the client to a new client state if this new client was committed to
// by the old client at the specified upgrade height
func (k *Keeper) UpgradeClient(
	ctx sdk.Context,
	clientID string,
	upgradedClient, upgradedConsState, upgradeClientProof, upgradeConsensusStateProof []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

// RecoverClient will invoke the light client module associated with the subject clientID requesting it to
// recover the subject client given a substitute client identifier. The light client implementation
// is responsible for validating the parameters of the substitute (ensuring they match the subject's parameters)
// as well as copying the necessary consensus states from the substitute to the subject client store.
// The substitute must be Active and the subject must not be Active.
func (k *Keeper) RecoverClient(ctx sdk.Context, subjectClientID, substituteClientID string) error {
	_ = "STUB: not implemented"
	return nil
}
