package server

import (
	"context"
	"fmt"
	"geteway/internal/config"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(cfg config.Config, handler http.Handler) error {
	op := "server.Run"
	s.httpServer = &http.Server{
		Addr:           ":" + cfg.Port,
		Handler:        handler,
		MaxHeaderBytes: 0,
		ReadTimeout:    cfg.Timeout,
		WriteTimeout:   cfg.Timeout,
	}
	return fmt.Errorf("%w:%s", s.httpServer.ListenAndServe(), op)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
