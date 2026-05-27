package keeper

// Option is an extension point to instantiate keeper with non default values
type Option interface {
	apply(*Keeper)
}

type optsFn func(*Keeper)

func (f optsFn) apply(keeper *Keeper) {
	_ = "STUB: not implemented"

	// WithQueryPlugins is an optional constructor parameter to pass custom query plugins for wasmVM requests.
	// Missing fields will be filled with default queriers.
	return
}

func WithQueryPlugins(plugins *QueryPlugins) Option { _ = "STUB: not implemented"; return *new(Option) }
