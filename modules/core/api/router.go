package api

// Router contains all the module-defined callbacks required by IBC Protocol V2.
type Router struct {
	// routes is a map from portID to IBCModule
	routes map[string]IBCModule
	// prefixRoutes is a map from portID prefix to IBCModule
	prefixRoutes map[string]IBCModule
}

// NewRouter creates a new Router instance.
func NewRouter() *Router { _ = "STUB: not implemented"; return nil }

// AddRoute registers a route for a given portID to a given IBCModule.
//
// Panics:
//   - if a route with the same portID has already been registered
//   - if the portID is not alphanumeric
func (rtr *Router) AddRoute(portID string, cbs IBCModule) *Router {
	_ = "STUB: not implemented"
	return nil
}

// Prevent existing prefix routes from colliding with the new direct route to avoid confusing behavior.

// AddPrefixRoute registers a route for a given portID prefix to a given IBCModule.
// A prefix route matches any portID that starts with the given prefix.
//
// Panics:
//   - if `portIDPrefix` is not alphanumeric.
//   - if a direct route `portIDPrefix` has already been registered.
//   - if a prefix of `portIDPrefix` is already registered as a prefix.
//   - if `portIDPrefix` is a prefix of am already registered prefix.
func (rtr *Router) AddPrefixRoute(portIDPrefix string, cbs IBCModule) *Router {
	_ = "STUB: not implemented"
	return nil
}

// If the prefix is a prefix of an already registered route, we panic to avoid confusing behavior.

// Prevent two scenarios:
//  * Adding a string that prefix is already registered e.g.
//    add prefix "portPrefix" and try to add "portPrefixSomeSuffix".
//  * Adding a string that is a prefix of already registered prefix route e.g.
//    add prefix "portPrefix" and try to add "port".

// Route returns the IBCModule for a given portID.
func (rtr *Router) Route(portID string) IBCModule {
	_ = "STUB: not implemented"
	return *new(IBCModule)
}

// HasRoute returns true if the Router has a module registered (whether it's a direct or a prefix route)
// for the portID or false if no module is registered for it.
func (rtr *Router) HasRoute(portID string) bool { _ = "STUB: not implemented"; return false }

// getRoute is a helper function that retrieves the IBCModule for a given portID.
func (rtr *Router) getRoute(portID string) (IBCModule, bool) {
	_ = "STUB: not implemented"
	// Direct routes take precedence over prefix routes
	return *new(IBCModule), false
}

// If the portID is not found as a direct route, check for prefix routes

// Note that this iteration is deterministic because there can only ever be one prefix route
// that matches a given portID. This is because of the checks in AddPrefixRoute preventing
// any colliding prefixes to be added.

// At this point neither a direct route nor a prefix route was found
