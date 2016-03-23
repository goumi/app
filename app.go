/*
Package app implements a minimal and flexible middleware stack that builds on
go http package in a similar way that KoaJS leverages NodeJS.

It provides adapters for the native http.Handler's but it creates a more modular
way to build applications based on the decorator pattern. The purpose is to provide
a good way to use middlware to decorate the context (response - request wrapper).

A usage example:

	m := app.New()

	// Loaded from elsewhere
	var handler web.Handler
	var httphandler http.Handler

	// Handler
	m.Use(handler)

	// http.Handler
	m.Use(web.HTTPHandler(httphandler))

	// Handler function
	m.Use(web.HandlerFunc(func(ctx Context) {

		// Do your stuff

		// Call next middlware
		ctx.Next()

		// You can do something else after the stack has run
	}))

*/
package app

import "net/http"
import "github.com/goumi/web"

// App is a chain application handler
type App interface {
	web.Handler
	http.Handler

	// Adds the use of middleware on top of the Handler
	Use(web.Handler)
}

// Module contains a chain of handlers
type Module []web.Handler

// New module
func New() App {
	return &Module{}
}

// ServeHTTP() serves as an entry point into the module, and translates the
// response writer and request into a single structure.
func (m *Module) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Create a context from the response and request
	ctx := web.NewContext(w, r)

	// Serve the app using the new context
	m.Serve(ctx)
}

// Serve extends the module with a new context, that holds the module
// middleware.
func (m *Module) Serve(ctx web.Context) {

	// Sandbox the context middleware
	ctx = NewContext(ctx, *m)

	// Run the middleware
	ctx.Next()
}

// Use adds a Handler to the the module chain
func (m *Module) Use(h web.Handler) {
	*m = append(*m, h)
}
