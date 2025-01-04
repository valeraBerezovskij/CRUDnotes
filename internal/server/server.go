package server

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

//Server initializing
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 10,
	}
	return s.httpServer.ListenAndServe()
}

//Server shutdowning
func (s *Server) Shutdown(ctx context.Context) error{
	return s.httpServer.Shutdown(ctx)
}