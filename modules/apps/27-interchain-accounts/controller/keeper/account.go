package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
)

// RegisterInterchainAccount is the entry point to registering an interchain account:
// - It generates a new port identifier using the provided owner string.
// - Callers are expected to provide the appropriate application version string.
// - For example, this could be an ICS27 encoded metadata type or an ICS29 encoded metadata type with a nested application version.
// - A new MsgChannelOpenInit is routed through the MsgServiceRouter, executing the OnOpenChanInit callback stack as configured.
// - An error is returned if the port identifier is already in use. Gaining access to interchain accounts whose channels
// have closed cannot be done with this function. A regular MsgChannelOpenInit must be used.
//
// Deprecated: this is a legacy API that is only intended to function correctly in workflows where an underlying authentication application has been set.
// Calling this API will result in all packet callbacks being routed to the underlying application.

// Please use MsgRegisterInterchainAccount for use cases which do not need to route to an underlying application.

// Prior to v6.x.x of ibc-go, the controller module was only functional as middleware, with authentication performed
// by the underlying application. For a full summary of the changes in v6.x.x, please see ADR009.
// This API will be removed in later releases.
func (k *Keeper) RegisterInterchainAccount(ctx sdk.Context, connectionID, owner, version string,
	ordering channeltypes.Order,
) error {
	_ = "STUB: not implemented"
	return nil
}

// use ORDER_UNORDERED as default in case ordering is NONE

// registerInterchainAccount registers an interchain account, returning the channel id of the MsgChannelOpenInitResponse
// and an error if one occurred.
func (k *Keeper) registerInterchainAccount(ctx sdk.Context, connectionID, portID, version string,
	ordering channeltypes.Order,
) (string, error) {
	_ = "STUB: not implemented"
	// if there is an active channel for this portID / connectionID return an error
	return "", nil
}

// NOTE: The sdk msg handler creates a new EventManager, so events must be correctly propagated back to the current context
