package httphandler

import (
	"docomo-bike/internal/listing"
	"net/http"
)

func HandleGetReservations(s listing.ReservationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func HandleGetReservation(s listing.ReservationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
