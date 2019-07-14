package http

import (
	"docomo-bike/internal/app"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes(srv *Server, cont *app.Container) (chi.Routes, error) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/auth", HandleAuthorize(cont.JWTAuthService))

	ar := r.With(UseAuth(cont.JWTAuthService, cont.DocomoClients.Login))
	ar.Get("/stations/{stationId}", HandleGetStation(cont.StationListingService))

	ar.Post("/bookings", HandleNewBikeBooking(cont.BikeBookingService))
	ar.Get("/bookings/current", HandleGetCurrentBikeBooking(cont.BikeBookingService))
	ar.Post("/bookings/current/cancel", HandleCancelBikeBooking(cont.BikeBookingService))
	ar.Post("/bookings/{bookingId}/complete", HandleCompleteBookingBike(cont.BikeBookingService))

	srv.httpServer.Handler = r

	return r, nil
}
