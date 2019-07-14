package http

import (
	"context"
	"docomo-bike/internal/app"
	"net/http"

	"github.com/pkg/errors"
)

func NewServer() (*Server, error) {
	return &Server{}, nil
}

type Server struct {
	container  *app.Container
	httpServer *http.Server
}

func (s *Server) ServeHTTP(addr string) error {
	s.httpServer = &http.Server{
		Addr: addr,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.container.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "Failed to shutdown the app conatiner")
	}
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "Failed to shutdown the http server")
	}
	return nil
}
