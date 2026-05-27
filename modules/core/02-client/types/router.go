package types

import (
	"github.com/cosmos/ibc-go/v11/modules/core/exported"
)

// Router is a map from a clientType to a LightClientModule instance.
type Router struct {
	routes map[string]exported.LightClientModule
}

// NewRouter returns an instance of the Router.
func NewRouter() *Router { _ = "STUB: not implemented"; return nil }

// AddRoute adds LightClientModule for a given module name. It returns the Router
// so AddRoute calls can be linked. This function will panic if:
// - the Router is sealed,
// - or a module is already registered for the provided client type,
// - or the client type is invalid.
func (rtr *Router) AddRoute(clientType string, module exported.LightClientModule) *Router {
	_ = "STUB: not implemented"
	return nil
}

// HasRoute returns true if the Router has a module registered or false otherwise.
func (rtr *Router) HasRoute(clientType string) bool { _ = "STUB: not implemented"; return false }

// GetRoute returns the LightClientModule registered for the provided client type or false otherwise.
func (rtr *Router) GetRoute(clientType string) (exported.LightClientModule, bool) {
	_ = "STUB: not implemented"
	return *new(exported.LightClientModule), false
}
