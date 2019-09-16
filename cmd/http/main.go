package main

import (
	"context"
	"docomo-bike/internal/app"
	"docomo-bike/internal/config"
	"docomo-bike/internal/http"
	"docomo-bike/internal/libs/env"
	"fmt"
	nethttp "net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi"
	"github.com/google/logger"
)

func main() {
	l := logger.Init("Main", false, false, os.Stdout)

	cfg := config.Config{}
	if err := cfg.Load(); err != nil {
		l.Fatalf("Failed to parse env variables: %+v", err)
	}
	l.Infof("Config loaded: %s", spew.Sdump(cfg))

	cont := app.NewContainer()
	if err := cont.Configure(cfg); err != nil {
		l.Fatalf("Failed to configure the app container: %+v", err)
	}
	srv, err := http.NewServer()
	if err != nil {
		l.Fatalf("Failed to create a server: %+v", err)
	}
	routes, err := http.Routes(srv, cont)
	if err != nil {
		l.Fatalf("Failed to add routes to the server: %+v", err)
	}
	if env.IsDev(cfg.Env) {
		err = chi.Walk(routes, func(method string, route string, handler nethttp.Handler, middlewares ...func(nethttp.Handler) nethttp.Handler) error {
			l.Infof(`Route added: [%s] %s`, method, route)
			return nil
		})
		if err != nil {
			l.Fatalf("Failed to walk routes: %+v", err)
		}
	}

	addr := fmt.Sprintf("%s:%s", cfg.HTTPServer.Host, cfg.HTTPServer.Port)
	l.Infof("Listening on %s", addr)

	err = srv.ServeHTTP(addr)
	if err != nethttp.ErrServerClosed {
		l.Fatalf("Server stopped unexpectedly: %+v", err)
	}

	idleConnsClosed := make(chan struct{})
	go watchSignal(idleConnsClosed, srv)
	<-idleConnsClosed

	l.Error("Server shutdown")
}
func watchSignal(idleConnsClosed chan struct{}, srv *http.Server) {
	sigint := make(chan os.Signal, 1)

	// interrupt signal sent from terminal
	signal.Notify(sigint, os.Interrupt)
	// sigterm signal sent from k8s
	signal.Notify(sigint, syscall.SIGTERM)

	<-sigint

	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Fatalf("Failed shutting down server: %+v", err)
	}
	close(idleConnsClosed)
}
