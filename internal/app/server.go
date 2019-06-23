package app

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
)

func NewServer(app *App) (*Server, error) {
	return &Server{
		app: app,
	}, nil
}

type Server struct {
	app        *App
	httpServer *http.Server
}

func (s *Server) ServeHTTP(addr string) error {
	s.httpServer = &http.Server{
		Addr:    addr,
		Handler: s.app.Router,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.app.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "Failed to shutdown app")
	}
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "Failed to shutdown app")
	}
	return nil
}
