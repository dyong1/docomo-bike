package http

import (
	"fmt"
	"net/http"

	"github.com/alioygur/gores"
)

func badRequest(w http.ResponseWriter, msg string) {
	gores.Error(w, http.StatusBadRequest, fmt.Sprintf("[BadRequest] %s", msg))
}
func internalServerError(w http.ResponseWriter, msg string) {
	gores.Error(w, http.StatusInternalServerError, fmt.Sprintf("[InternalServerError] %s", msg))
}
func unauthorized(w http.ResponseWriter, msg string) {
	gores.Error(w, http.StatusUnauthorized, fmt.Sprintf("[Unauthorized] %s", msg))
}
func notFound(w http.ResponseWriter, msg string) {
	gores.Error(w, http.StatusNotFound, fmt.Sprintf("[NotFound] %s", msg))
}
