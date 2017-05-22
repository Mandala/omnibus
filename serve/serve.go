// Copyright (c) 2017 Fadhli Dzil Ikram. All rights reserved.
// This source code is brought to you under MIT license that can be found
// on the LICENSE file.

package serve

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// ErrServerRunning represents unavailable action on running server
var ErrServerRunning = errors.New("omnibus-server: Already running")

// ErrServerStopped represents unavailable action on stopped server
var ErrServerStopped = errors.New("omnibus-server: Already stopped")

// Store global handle state
var global *Server

/*// Template for default not found handler
var tmplNotFound *template.Template*/

// Options store server related configurations
type Options struct {
	Address         string
	HeaderTimeout   time.Duration
	ConnTimeout     time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration
	MaxHeaderBytes  int
	MaxBytes        int64
}

// Server store server handle state
type Server struct {
	Options
	Handler http.Handler
	mu      sync.Mutex
	server  *http.Server
	stop    chan os.Signal
}

// WithOptions set handle configurations
func (h *Server) WithOptions(o Options) *Server {
	h.Options = o
	return h
}

// Use set handler configuration
func (h *Server) Use(handler http.Handler) *Server {
	h.Handler = handler
	return h
}

// Run server and wait until explicit Stop() function or interrupt signal
func (h *Server) Run() error {
	h.mu.Lock()
	defer h.mu.Unlock()
	// Do not continue if server already running
	if h.server != nil {
		return ErrServerRunning
	}
	// Initialize server handle
	h.server = &http.Server{
		Addr:              h.Address,
		Handler:           h,
		IdleTimeout:       h.IdleTimeout,
		ReadHeaderTimeout: h.HeaderTimeout,
		ReadTimeout:       h.ConnTimeout,
		WriteTimeout:      h.ConnTimeout,
	}
	// Deallocate server handle on function exit
	defer func() {
		h.server = nil
	}()
	// Run http.Server in separate goroutine and initialize error channel
	errChan := make(chan error)
	go func() {
		if err := h.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- err
			return
		}
		errChan <- nil
	}()
	// Initialize stop signal catcher
	h.stop = make(chan os.Signal)
	signal.Notify(h.stop, syscall.SIGTERM, syscall.SIGINT)
	// Wait until stop signal is closed or triggered
	h.mu.Unlock()
	<-h.stop
	h.mu.Lock()
	// Create context for shutdown process
	ctx := context.Background()
	if h.ShutdownTimeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, h.ShutdownTimeout)
		defer cancel()
	}
	// Start shutdown process
	if err := h.server.Shutdown(ctx); err != nil {
		// Close server error channel and return shutdown error instead
		close(errChan)
		return err
	}
	// Return error from listenAndServe
	return <-errChan
}

// Stop server without waiting for interrupt signal
func (h *Server) Stop() error {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.server == nil {
		return ErrServerStopped
	}
	// Stop os.signal notifier before closing the channel
	signal.Stop(h.stop)
	close(h.stop)
	return nil
}

// ServeHTTP implements http.Handler for maximum request body handler
func (h *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Limit body io.Reader with http.MaxBytesReader if MaxBytes option set
	if h.MaxBytes > 0 {
		r.Body = http.MaxBytesReader(w, r.Body, h.MaxBytes)
	}
	// Run the entrypoint handler
	if h.Handler != nil {
		h.Handler.ServeHTTP(w, r)
		return
	}
	// Give 404 to user
	w.WriteHeader(404)
	w.Write([]byte("404 Not Found"))
}

/*// notFoundHandler acts as default not found page for the router
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	tmplNotFound.Execute(w, r.URL.Path)
}*/

// NewServer create new empty server object
func NewServer() *Server {
	return &Server{}
}

func init() {
	// Set global with new server handle
	global = NewServer()
	// Initialize not found template
	/*tmplNotFound = template.Must(template.New("").Parse(`<!DOCTYPE html>
	<html lang="en">
	<head>
	  <title>Not Found</title>
	  <meta charset="UTF-8">
	  <meta name="viewport" content="width=device-width, initial-scale=1">
	  <style>
	  body {
	    font-family: Georgia, serif;
	  }
	  .resource {
	    font-family: "Lucida Console", Monaco, monospace;
	  }
	  </style>
	</head>
	<body>
	  <h1>Page Not Found</h1>
	  <p>
	    Resource <span class="resource">{{.}}</span> was not found on the server.
	    Please double check the entered URL and try again.
	  </p>
	</body>
	</html>
	`))*/
}

// WithOptions set global configurations
func WithOptions(o Options) *Server {
	return global.WithOptions(o)
}

// Use set global handler configuration
func Use(handler http.Handler) *Server {
	return global.Use(handler)
}

// GetOptions get global configurations
func GetOptions() *Options {
	return &global.Options
}

// GetInstance get global instance handle
func GetInstance() *Server {
	return global
}

// Run server and wait until explicit Stop() function or interrupt signal
func Run() error {
	return global.Run()
}

// Stop server without waiting for interrupt signal
func Stop() error {
	return global.Stop()
}
