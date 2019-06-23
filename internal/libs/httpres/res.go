package httpres

import (
	"net/http"

	"github.com/alioygur/gores"
)

// JSON ignores errors internally created.
func JSON(w http.ResponseWriter, data interface{}) {
	if err := gores.JSON(w, http.StatusOK, data); err != nil {
		gores.Error(w, http.StatusInternalServerError, err.Error())
	}
}

func Error(w http.ResponseWriter, statusCode int, err error) {
	gores.Error(w, statusCode, err.Error())
}
