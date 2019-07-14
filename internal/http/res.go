package http

import (
	"net/http"

	"github.com/alioygur/gores"
)

func jsonres(w http.ResponseWriter, body interface{}) {
	if err := gores.JSON(w, http.StatusOK, body); err != nil {
		internalServerError(w, err)
	}
}
