// Copyright (c) 2017 Fadhli Dzil Ikram. All rights reserved.
// This source code is brought to you under MIT license that can be found
// on the LICENSE file.

package muxer

import (
	"errors"
	"net/url"

	"net/http"

	"github.com/gorilla/mux"
)

var ErrRouteNotFound = errors.New("muxer: Named route was not found")

type Router struct {
	*mux.Router
	middleware []Middleware
}

func NewRouter() *Router {
	router := &Router{
		Router: mux.NewRouter(),
	}
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	return router
}

func (r *Router) wrapRoute(route *mux.Route) *Route {
	return &Route{
		Route:      route,
		middleware: r.middleware,
	}
}

func (r *Router) GetURL(name string, pairs ...string) (*url.URL, error) {
	if route := r.Router.Get(name); route != nil {
		return route.URL(pairs...)
	}
	return nil, ErrRouteNotFound
}

func (r *Router) NewRoute() *Route {
	return r.wrapRoute(r.Router.NewRoute())
}

func (r *Router) Middleware(middleware ...Middleware) *Route {
	return r.NewRoute().Middleware(middleware...)
}

func (r *Router) Handle(path string, handler http.Handler) *Route {
	return r.NewRoute().Path(path).Handler(handler)
}

func (r *Router) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return r.Handle(path, http.HandlerFunc(f))
}

func (r *Router) Headers(pairs ...string) *Route {
	return r.NewRoute().Headers(pairs...)
}

func (r *Router) Host(tpl string) *Route {
	return r.NewRoute().Host(tpl)
}

func (r *Router) Methods(methods ...string) *Route {
	return r.NewRoute().Methods(methods...)
}

func (r *Router) Path(tpl string) *Route {
	return r.NewRoute().Path(tpl)
}

func (r *Router) PathPrefix(tpl string) *Route {
	return r.NewRoute().PathPrefix(tpl)
}

func (r *Router) Queries(pairs ...string) *Route {
	return r.NewRoute().Queries(pairs...)
}

func (r *Router) Schemes(schemes ...string) *Route {
	return r.NewRoute().Schemes(schemes...)
}

func (r *Router) SkipClean(value bool) *Router {
	r.Router.SkipClean(value)
	return r
}

func (r *Router) StrictSlash(value bool) *Router {
	r.Router.StrictSlash(value)
	return r
}

func (r *Router) UseEncodedPath() *Router {
	r.Router.UseEncodedPath()
	return r
}

// Router shorthands

func (r *Router) Get(path string, handler http.Handler) *Route {
	return r.NewRoute().Get(path, handler)
}

func (r *Router) GetFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return r.NewRoute().GetFunc(path, f)
}

func (r *Router) Post(path string, handler http.Handler) *Route {
	return r.NewRoute().Post(path, handler)
}

func (r *Router) PostFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return r.NewRoute().PostFunc(path, f)
}

func (r *Router) Put(path string, handler http.Handler) *Route {
	return r.NewRoute().Put(path, handler)
}

func (r *Router) PutFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return r.NewRoute().PutFunc(path, f)
}

func (r *Router) Patch(path string, handler http.Handler) *Route {
	return r.NewRoute().Patch(path, handler)
}

func (r *Router) PatchFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return r.NewRoute().PatchFunc(path, f)
}

func (r *Router) Delete(path string, handler http.Handler) *Route {
	return r.NewRoute().Delete(path, handler)
}

func (r *Router) DeleteFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return r.NewRoute().DeleteFunc(path, f)
}

func (r *Router) Head(path string, handler http.Handler) *Route {
	return r.NewRoute().Head(path, handler)
}

func (r *Router) HeadFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return r.NewRoute().HeadFunc(path, f)
}

func (r *Router) Options(path string, handler http.Handler) *Route {
	return r.NewRoute().Options(path, handler)
}

func (r *Router) OptionsFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return r.NewRoute().OptionsFunc(path, f)
}
