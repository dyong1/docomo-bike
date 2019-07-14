package main

import (
	"context"
	"docomo-bike/internal/app"
	"docomo-bike/internal/config"
	"docomo-bike/internal/http"
	"fmt"
	nethttp "net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/davecgh/go-spew/spew"
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
	if err := http.Routes(srv, cont); err != nil {
		l.Fatalf("Failed to add routes to the server: %+v", err)
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
