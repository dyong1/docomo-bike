package app

import (
	"docomo-bike/internal/auth"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (a *App) Configure() error {
	authService := &auth.BasicService{}

	{
		router := chi.NewRouter()
		router.Use(middleware.Logger)
		router.Use(middleware.Recoverer)

		router.Post("/auth", auth.HandleAuthorize(authService))

		a.Router = router
	}

	return nil
}
