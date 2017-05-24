// Copyright (c) 2017 Fadhli Dzil Ikram. All rights reserved.
// This source code is brought to you under MIT license that can be found
// on the LICENSE file.

// Some documentation in this source code are derived from gorilla/mux package.

package route

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

// ErrRouteNotFound define named route not found on GetURL method
var ErrRouteNotFound = errors.New("muxer: Named route was not found")

// Router registers routes to be matched and dispatches a handler.
type Router struct {
	*mux.Router
	middleware []Middleware
}

// NewRouter returns a new router instance.
func NewRouter() *Router {
	router := &Router{
		Router: mux.NewRouter(),
	}
	router.NotFoundHandler = NotFoundHandler
	return router
}

// GetURL gets URL object from a named router in route list.
func (r *Router) GetURL(name string, pairs ...string) (*url.URL, error) {
	if route := r.Router.Get(name); route != nil {
		return route.URL(pairs...)
	}
	return nil, ErrRouteNotFound
}

// wrapRoute wraps mux.Route object with current middleware configuration.
func (r *Router) wrapRoute(route *mux.Route) *Route {
	return &Route{
		Route:      route,
		middleware: r.middleware,
	}
}

// NewRoute registers an empty route.
func (r *Router) NewRoute() *Route {
	return r.wrapRoute(r.Router.NewRoute())
}

// Middleware adds middleware function to current route.
func (r *Router) Middleware(middleware ...Middleware) *Route {
	return r.NewRoute().Middleware(middleware...)
}

// BuildVarsFunc adds a custom function to be used to modify build variables
// before a route's URL is built.
func (r *Router) BuildVarsFunc(f mux.BuildVarsFunc) *Route {
	return r.NewRoute().BuildVarsFunc(f)
}

// MiddlewareFunc adds middleware function to current route.
func (r *Router) MiddlewareFunc(f func(http.ResponseWriter, *http.Request, http.Handler)) *Route {
	return r.NewRoute().MiddlewareFunc(f)
}

// Handle registers a new route with a matcher for the URL path. See
// Route.Path() and Route.Handler().
func (r *Router) Handle(path string, handler http.Handler) *Route {
	return r.NewRoute().Path(path).Handler(handler)
}

// HandleFunc registers a new route with a matcher for the URL path. See
// Route.Path() and Route.HandlerFunc().
func (r *Router) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return r.Handle(path, http.HandlerFunc(f))
}

// Headers registers a new route with a matcher for request header values.
func (r *Router) Headers(pairs ...string) *Route {
	return r.NewRoute().Headers(pairs...)
}

// Host registers a new route with a matcher for the URL host.
func (r *Router) Host(tpl string) *Route {
	return r.NewRoute().Host(tpl)
}

// MatcherFunc registers a new route with a custom matcher function.
func (r *Router) MatcherFunc(f mux.MatcherFunc) *Route {
	return r.NewRoute().MatcherFunc(f)
}

// Methods registers a new route with a matcher for HTTP methods.
func (r *Router) Methods(methods ...string) *Route {
	return r.NewRoute().Methods(methods...)
}

// Path registers a new route with a matcher for the URL path.
func (r *Router) Path(tpl string) *Route {
	return r.NewRoute().Path(tpl)
}

// PathPrefix registers a new route with a matcher for the URL path prefix.
func (r *Router) PathPrefix(tpl string) *Route {
	return r.NewRoute().PathPrefix(tpl)
}

// Queries registers a new route with a matcher for URL query values.
func (r *Router) Queries(pairs ...string) *Route {
	return r.NewRoute().Queries(pairs...)
}

// Schemes registers a new route with a matcher for URL schemes.
func (r *Router) Schemes(schemes ...string) *Route {
	return r.NewRoute().Schemes(schemes...)
}

// SkipClean defines the path cleaning behaviour for new routes.
func (r *Router) SkipClean(value bool) *Router {
	r.Router.SkipClean(value)
	return r
}

// StrictSlash defines the trailing slash behavior for new routes.
func (r *Router) StrictSlash(value bool) *Router {
	r.Router.StrictSlash(value)
	return r
}

// UseEncodedPath tells the router to match the encoded original path to
// the routes.
func (r *Router) UseEncodedPath() *Router {
	r.Router.UseEncodedPath()
	return r
}

// Get sets handler with GET method and specified path matcher.
func (r *Router) Get(path string, handler http.Handler) *Route {
	return r.NewRoute().Get(path, handler)
}

// GetFunc sets handler function with GET method and specified path matcher.
func (r *Router) GetFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return r.NewRoute().GetFunc(path, f)
}

// Post sets handler with POST method and specified path matcher.
func (r *Router) Post(path string, handler http.Handler) *Route {
	return r.NewRoute().Post(path, handler)
}

// PostFunc sets handler function with POST method and specified path matcher.
func (r *Router) PostFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return r.NewRoute().PostFunc(path, f)
}

// Put sets handler with PUT method and specified path matcher.
func (r *Router) Put(path string, handler http.Handler) *Route {
	return r.NewRoute().Put(path, handler)
}

// PutFunc sets handler function with PUT method and specified path matcher.
func (r *Router) PutFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return r.NewRoute().PutFunc(path, f)
}

// Patch sets handler with PATCH method and specified path matcher.
func (r *Router) Patch(path string, handler http.Handler) *Route {
	return r.NewRoute().Patch(path, handler)
}

// PatchFunc sets handler function with PATCH method and specified path matcher.
func (r *Router) PatchFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return r.NewRoute().PatchFunc(path, f)
}

// Delete sets handler with DELETE method and specified path matcher.
func (r *Router) Delete(path string, handler http.Handler) *Route {
	return r.NewRoute().Delete(path, handler)
}

// DeleteFunc sets handler function with PATCH method and specified
// path matcher.
func (r *Router) DeleteFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return r.NewRoute().DeleteFunc(path, f)
}

// Head sets handler with HEAD method and specified path matcher.
func (r *Router) Head(path string, handler http.Handler) *Route {
	return r.NewRoute().Head(path, handler)
}

// HeadFunc sets handler function with HEAD method and specified path matcher.
func (r *Router) HeadFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return r.NewRoute().HeadFunc(path, f)
}

// Options sets handler with OPTIONS method and specified path matcher.
func (r *Router) Options(path string, handler http.Handler) *Route {
	return r.NewRoute().Options(path, handler)
}

// OptionsFunc sets handler function with OPTIONS method and specified
// path matcher.
func (r *Router) OptionsFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return r.NewRoute().OptionsFunc(path, f)
}
