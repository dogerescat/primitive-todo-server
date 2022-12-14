package http

import (
	"context"
	"fmt"
	"net/http"
	"primitive-todo-server/internal/db"
)

type Server struct {
	server *http.Server
}

func NewServer(port int, d db.DB) *Server {
	mux := http.NewServeMux()
	mux.Handle("/create", &createHandler{db: d})
	mux.Handle("/list", &listHandler{db: d})
	return &Server{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: mux,
		},
	}
}

func (s *Server) Start() error {
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to serve: %w", err)
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown %w", err)
	}

	return nil
}
