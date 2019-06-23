package httphandler

import (
	"docomo-bike/internal/booking"
	"net/http"
)

func HandleBook(s booking.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
