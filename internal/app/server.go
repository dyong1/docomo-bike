package app

import (
	"net/http"
)

type Server struct {
	app *App
}

func (s *Server) ServeHTTP(addr string) error {
	return http.ListenAndServe(addr, s.app.Router)
}

func NewServer() (*Server, error) {
	a := &App{}
	if err := a.Configure(); err != nil {
		return nil, err
	}
	return &Server{
		app: a,
	}, nil
}
