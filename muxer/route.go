// Copyright (c) 2017 Fadhli Dzil Ikram. All rights reserved.
// This source code is brought to you under MIT license that can be found
// on the LICENSE file.

package muxer

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	*mux.Route
	middleware []Middleware
	handler    http.Handler
}

func (r *Route) Middleware(middleware ...Middleware) *Route {
	r.middleware = append(r.middleware, middleware...)
	// Reapply r.Handler if already defined
	if r.handler != nil {
		r.Handler(r.handler)
	}
	return r
}

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

func (r *Route) HandlerFunc(f func(http.ResponseWriter, *http.Request)) *Route {
	return r.Handler(http.HandlerFunc(f))
}

func (r *Route) wrapRouter(router *mux.Router) *Router {
	return &Router{
		Router:     router,
		middleware: r.middleware,
	}
}

func (r *Route) Subrouter() *Router {
	return r.wrapRouter(r.Route.Subrouter())
}

func (r *Route) Group(f func(*Router)) {
	f(r.Subrouter())
}

func (r *Route) Headers(pairs ...string) *Route {
	r.Route.Headers(pairs...)
	return r
}

func (r *Route) HeadersRegexp(pairs ...string) *Route {
	r.Route.HeadersRegexp(pairs...)
	return r
}

func (r *Route) Host(tpl string) *Route {
	r.Route.Host(tpl)
	return r
}

func (r *Route) Methods(methods ...string) *Route {
	r.Route.Methods(methods...)
	return r
}

func (r *Route) Name(name string) *Route {
	r.Route.Name(name)
	return r
}

func (r *Route) Path(tpl string) *Route {
	r.Route.Path(tpl)
	return r
}

func (r *Route) PathPrefix(tpl string) *Route {
	r.Route.PathPrefix(tpl)
	return r
}

func (r *Route) Queries(pairs ...string) *Route {
	r.Route.Queries(pairs...)
	return r
}

func (r *Route) Schemes(schemes ...string) *Route {
	r.Route.Schemes(schemes...)
	return r
}

// Route shorthands

func (r *Route) Get(path string, handler http.Handler) *Route {
	return r.Methods("GET").Path(path).Handler(handler)
}

func (r *Route) GetFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return r.Get(path, http.HandlerFunc(f))
}

func (r *Route) Post(path string, handler http.Handler) *Route {
	return r.Methods("POST").Path(path).Handler(handler)
}

func (r *Route) PostFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return r.Post(path, http.HandlerFunc(f))
}

func (r *Route) Put(path string, handler http.Handler) *Route {
	return r.Methods("PUT").Path(path).Handler(handler)
}

func (r *Route) PutFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return r.Put(path, http.HandlerFunc(f))
}

func (r *Route) Patch(path string, handler http.Handler) *Route {
	return r.Methods("PATCH").Path(path).Handler(handler)
}

func (r *Route) PatchFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return r.Patch(path, http.HandlerFunc(f))
}

func (r *Route) Delete(path string, handler http.Handler) *Route {
	return r.Methods("DELETE").Path(path).Handler(handler)
}

func (r *Route) DeleteFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return r.Delete(path, http.HandlerFunc(f))
}

func (r *Route) Head(path string, handler http.Handler) *Route {
	return r.Methods("HEAD").Path(path).Handler(handler)
}

func (r *Route) HeadFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return r.Head(path, http.HandlerFunc(f))
}

func (r *Route) Options(path string, handler http.Handler) *Route {
	return r.Methods("OPTIONS").Path(path).Handler(handler)
}

func (r *Route) OptionsFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return r.Options(path, http.HandlerFunc(f))
}
