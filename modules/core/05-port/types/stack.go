package types

type IBCStackBuilder struct {
	middlewares   []Middleware
	baseModule    IBCModule
	channelKeeper ICS4Wrapper
}

func NewIBCStackBuilder(chanKeeper ICS4Wrapper) *IBCStackBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *IBCStackBuilder) Next(middleware Middleware) *IBCStackBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *IBCStackBuilder) Base(baseModule IBCModule) *IBCStackBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *IBCStackBuilder) Build() IBCModule { _ = "STUB: not implemented"; return *new(IBCModule) }

// Build the stack by moving up the middleware list
// and setting the underlying application for each middleware
// and the ICS4wrapper for the underlying module.

// set the top level channel keeper as the ICS4Wrapper
// for the lop level middleware
