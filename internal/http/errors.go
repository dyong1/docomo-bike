package http

import (
	"net/http"

	"github.com/alioygur/gores"
)

func badRequest(w http.ResponseWriter, err error) {
	gores.Error(w, http.StatusBadRequest, err.Error())
}
func internalServerError(w http.ResponseWriter, err error) {
	gores.Error(w, http.StatusInternalServerError, err.Error())
}
func unauthorized(w http.ResponseWriter, err error) {
	gores.Error(w, http.StatusUnauthorized, err.Error())
}
