package http

import (
	"docomo-bike/internal/services/bikebooking"
	"net/http"
)

func HandleBook(s bikebooking.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
