// Copyright (c) 2017 Fadhli Dzil Ikram. All rights reserved.
// This source code is brought to you under MIT license that can be found
// on the LICENSE file.

// Some documentation in this source code are derived from gorilla/mux package.

package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route stores information to match a request and build URLs.
type Route struct {
	*mux.Route
	middleware []Middleware
	handler    http.Handler
}

// BuildVarsFunc adds a custom function to be used to modify build variables
// before a route's URL is built.
func (r *Route) BuildVarsFunc(f mux.BuildVarsFunc) *Route {
	r.BuildVarsFunc(f)
	return r
}

// Middleware adds middleware function to current route.
func (r *Route) Middleware(middleware ...Middleware) *Route {
	r.middleware = append(r.middleware, middleware...)
	// Reapply r.Handler if already defined
	if r.handler != nil {
		r.Handler(r.handler)
	}
	return r
}

// MiddlewareFunc adds middleware function to current route.
func (r *Route) MiddlewareFunc(f func(http.ResponseWriter, *http.Request, http.Handler)) *Route {
	return r.Middleware(MiddlewareFunc(f))
}

// Handler sets a handler for the route.
func (r *Route) Handler(handler http.Handler) *Route {
	if len(r.middleware) > 0 {
		r.Route.Handler(MiddlewareRunner{
			Handler: handler,
			Stack:   r.middleware,
		})
	} else {
		r.Route.Handler(handler)
	}
	r.handler = handler

	return r
}

// HandlerFunc sets a handler function for the route.
func (r *Route) HandlerFunc(f func(http.ResponseWriter, *http.Request)) *Route {
	return r.Handler(http.HandlerFunc(f))
}

// wrapRouter wraps mux.Router object with current middleware configuration.
func (r *Route) wrapRouter(router *mux.Router) *Router {
	return &Router{
		Router:     router,
		middleware: r.middleware,
	}
}

// Subrouter creates a subrouter for the route.
func (r *Route) Subrouter() *Router {
	return r.wrapRouter(r.Route.Subrouter())
}

// Group creates a subrouter for the route in closure function style.
func (r *Route) Group(f func(*Router)) {
	f(r.Subrouter())
}

// Headers adds a matcher for request header values.
func (r *Route) Headers(pairs ...string) *Route {
	r.Route.Headers(pairs...)
	return r
}

// HeadersRegexp accepts a sequence of key/value pairs, where the value has
// regex support.
func (r *Route) HeadersRegexp(pairs ...string) *Route {
	r.Route.HeadersRegexp(pairs...)
	return r
}

// Host adds a matcher for the URL host.
func (r *Route) Host(tpl string) *Route {
	r.Route.Host(tpl)
	return r
}

// MatcherFunc adds a custom function to be used as request matcher.
func (r *Route) MatcherFunc(f mux.MatcherFunc) *Route {
	r.Route.MatcherFunc(f)
	return r
}

// Methods adds a matcher for HTTP methods.
func (r *Route) Methods(methods ...string) *Route {
	r.Route.Methods(methods...)
	return r
}

// Name sets the name for the route, used to build URLs. If the name was
// registered already it will be overwritten.
func (r *Route) Name(name string) *Route {
	r.Route.Name(name)
	return r
}

// Path adds a matcher for the URL path.
func (r *Route) Path(tpl string) *Route {
	r.Route.Path(tpl)
	return r
}

// PathPrefix adds a matcher for the URL path prefix.
func (r *Route) PathPrefix(tpl string) *Route {
	r.Route.PathPrefix(tpl)
	return r
}

// Queries adds a matcher for URL query values.
func (r *Route) Queries(pairs ...string) *Route {
	r.Route.Queries(pairs...)
	return r
}

// Schemes adds a matcher for URL schemes.
func (r *Route) Schemes(schemes ...string) *Route {
	r.Route.Schemes(schemes...)
	return r
}

// Get sets handler with GET method and specified path matcher.
func (r *Route) Get(path string, handler http.Handler) *Route {
	return r.Methods("GET").Path(path).Handler(handler)
}

// GetFunc sets handler function with GET method and specified path matcher.
func (r *Route) GetFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return r.Get(path, http.HandlerFunc(f))
}

// Post sets handler with POST method and specified path matcher.
func (r *Route) Post(path string, handler http.Handler) *Route {
	return r.Methods("POST").Path(path).Handler(handler)
}

// PostFunc sets handler function with POST method and specified path matcher.
func (r *Route) PostFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return r.Post(path, http.HandlerFunc(f))
}

// Put sets handler with PUT method and specified path matcher.
func (r *Route) Put(path string, handler http.Handler) *Route {
	return r.Methods("PUT").Path(path).Handler(handler)
}

// PutFunc sets handler function with PUT method and specified path matcher.
func (r *Route) PutFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return r.Put(path, http.HandlerFunc(f))
}

// Patch sets handler with PATCH method and specified path matcher.
func (r *Route) Patch(path string, handler http.Handler) *Route {
	return r.Methods("PATCH").Path(path).Handler(handler)
}

// PatchFunc sets handler function with PATCH method and specified path matcher.
func (r *Route) PatchFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return r.Patch(path, http.HandlerFunc(f))
}

// Delete sets handler with DELETE method and specified path matcher.
func (r *Route) Delete(path string, handler http.Handler) *Route {
	return r.Methods("DELETE").Path(path).Handler(handler)
}

// DeleteFunc sets handler function with PATCH method and specified
// path matcher.
func (r *Route) DeleteFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return r.Delete(path, http.HandlerFunc(f))
}

// Head sets handler with HEAD method and specified path matcher.
func (r *Route) Head(path string, handler http.Handler) *Route {
	return r.Methods("HEAD").Path(path).Handler(handler)
}

// HeadFunc sets handler function with HEAD method and specified path matcher.
func (r *Route) HeadFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return r.Head(path, http.HandlerFunc(f))
}

// Options sets handler with OPTIONS method and specified path matcher.
func (r *Route) Options(path string, handler http.Handler) *Route {
	return r.Methods("OPTIONS").Path(path).Handler(handler)
}

// OptionsFunc sets handler function with OPTIONS method and specified
// path matcher.
func (r *Route) OptionsFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return r.Options(path, http.HandlerFunc(f))
}
