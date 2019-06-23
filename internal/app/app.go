package app

import (
	"context"
	"docomo-bike/internal/config"

	"github.com/go-chi/chi"
	"github.com/google/logger"
	"github.com/pkg/errors"
)

func NewApp(cfg config.Config) (*App, error) {
	a := &App{}
	if err := a.Configure(cfg); err != nil {
		return nil, errors.Wrap(err, "")
	}
	return a, nil
}

type App struct {
	Router chi.Router
	Logger *logger.Logger
}

func (a *App) Shutdown(ctx context.Context) error {
	return nil
}
