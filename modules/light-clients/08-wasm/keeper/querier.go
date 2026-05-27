package keeper

import (
	"encoding/json"

	wasmvmtypes "github.com/CosmWasm/wasmvm/v3/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/modules/light-clients/08-wasm/v11/types"
)

/*
`queryHandler` is a contextual querier which references the global `ibcwasm.QueryPluginsI`
to handle queries. The global `ibcwasm.QueryPluginsI` points to a `types.QueryPlugins` which
contains two sub-queriers: `types.CustomQuerier` and `types.StargateQuerier`. These sub-queriers
can be replaced by the user through the options api in the keeper.

In addition, the `types.StargateQuerier` references a global `types.QueryRouter` which points
to `baseapp.GRPCQueryRouter`.

This design is based on wasmd's (v0.50.0) querier plugin design.
*/

var _ wasmvmtypes.Querier = (*queryHandler)(nil)

// defaultAcceptList defines a set of default allowed queries made available to the Querier.
var defaultAcceptList = []string{
	"/ibc.core.client.v1.Query/VerifyMembership",
}

// queryHandler is a wrapper around the sdk.Context and the CallerID that calls
// into the query plugins.
type queryHandler struct {
	Ctx      sdk.Context
	Plugins  QueryPlugins
	CallerID string
}

// newQueryHandler returns a default querier that can be used in the contract.
func newQueryHandler(ctx sdk.Context, plugins QueryPlugins, callerID string) *queryHandler {
	_ = "STUB: not implemented"
	return nil
}

// GasConsumed implements the wasmvmtypes.Querier interface.
func (q *queryHandler) GasConsumed() uint64 { _ = "STUB: not implemented"; return 0 }

// Query implements the wasmvmtypes.Querier interface.
func (q *queryHandler) Query(request wasmvmtypes.QueryRequest, gasLimit uint64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// discard all changes/events in subCtx by not committing the cached context

// make sure we charge the higher level context even on panic

type (
	CustomQuerier   func(ctx sdk.Context, request json.RawMessage) ([]byte, error)
	StargateQuerier func(ctx sdk.Context, request *wasmvmtypes.StargateQuery) ([]byte, error)

	// QueryPlugins is a list of queriers that can be used to extend the default querier.
	QueryPlugins struct {
		Custom   CustomQuerier
		Stargate StargateQuerier
	}
)

// Merge merges the query plugin with a provided one.
func (e QueryPlugins) Merge(x *QueryPlugins) QueryPlugins {
	_ = "STUB: not implemented"
	// only update if this is non-nil and then only set values
	return *new(QueryPlugins)
}

// HandleQuery implements the ibcwasm.QueryPluginsI interface.
func (e QueryPlugins) HandleQuery(ctx sdk.Context, caller string, request wasmvmtypes.QueryRequest) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewDefaultQueryPlugins returns the default set of query plugins
func NewDefaultQueryPlugins(queryRouter types.QueryRouter) QueryPlugins {
	_ = "STUB: not implemented"
	return *new(QueryPlugins)
}

// AcceptListStargateQuerier allows all queries that are in the provided accept list.
// This function returns protobuf encoded responses in bytes.
func AcceptListStargateQuerier(acceptedQueries []string, queryRouter types.QueryRouter) func(sdk.Context, *wasmvmtypes.StargateQuery) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil
}

// append user defined accepted queries to default list defined above.

// RejectCustomQuerier rejects all custom queries
func RejectCustomQuerier() func(sdk.Context, json.RawMessage) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil
}

// Wasmd Issue [#759](https://github.com/CosmWasm/wasmd/issues/759)
// Don't return error string for worries of non-determinism
func redactError(err error) error {
	_ = "STUB: not implemented"
	// Do not redact system errors
	// SystemErrors must be created in 08-wasm and we can ensure determinism
	return nil
}
