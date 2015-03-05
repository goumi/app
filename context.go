package app

import "github.com/goumi/web"

// Context extends the context to access the module chain
type context struct {
	web.Context

	// Contain the middleware and the index for the current middleware
	chain []web.Handler
	index int
}

// NewContext creates a new module context from the previous context and
// the module middleware
func NewContext(ctx web.Context, m Module) web.Context {
	return &context{
		Context: ctx,
		chain:   m,
		index:   -1,
	}
}

// Next runs the next middleware
func (ctx *context) Next() {

	// Increment
	ctx.index++

	// Check if we have middleware in the current chain
	if ctx.index < len(ctx.chain) {

		// Serve the current handler
		ctx.chain[ctx.index].Serve(ctx)

		// Done
		return
	}

	// Exit current chain, advance to the next one
	ctx.Context.Next()
}
