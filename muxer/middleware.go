// Copyright (c) 2017 Fadhli Dzil Ikram. All rights reserved.
// This source code is brought to you under MIT license that can be found
// on the LICENSE file.

package muxer

import (
	"net/http"
)

type Middleware interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request, next http.Handler)
}

type MiddlewareFunc func(http.ResponseWriter, *http.Request, http.Handler)

func (m MiddlewareFunc) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.Handler) {
	m(w, r, next)
}

type MiddlewareRunner struct {
	Stack   []Middleware
	Handler http.Handler
}

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
