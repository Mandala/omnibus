// Copyright (c) 2017 Fadhli Dzil Ikram. All rights reserved.
// This source code is brought to you under MIT license that can be found
// on the LICENSE file.

package route

import (
	"net/http"
)

// Middleware define the middleware interface that resembles to http.Handler.
type Middleware interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request, next http.Handler)
}

// MiddlewareFunc acts as simple middleware function to interface converter.
type MiddlewareFunc func(http.ResponseWriter, *http.Request, http.Handler)

// ServeHTTP implements route.Middleware interface.
func (m MiddlewareFunc) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.Handler) {
	m(w, r, next)
}

// MiddlewareRunner will orderly run the middleware stack slice and handler.
type MiddlewareRunner struct {
	Stack   []Middleware
	Handler http.Handler
}

// ServeHTTP implements http.Handler interface.
func (h MiddlewareRunner) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if len(h.Stack) > 0 {
		h.Stack[0].ServeHTTP(w, r, MiddlewareRunner{
			Handler: h.Handler,
			Stack:   h.Stack[1:],
		})
		return
	}
	h.Handler.ServeHTTP(w, r)
}
