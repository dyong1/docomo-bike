package http

import (
	"docomo-bike/internal/app"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes(srv *Server, cont *app.Container) error {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Post("/auth", HandleAuthorize(cont.JWTAuthService))
	router.Route("/me", func(me chi.Router) {
		me.Use(UseAuth(cont.JWTAuthService))
		me.Post("/booking", HandleBook(cont.BikeBookingService))
	})

	srv.httpServer.Handler = router

	return nil
}
