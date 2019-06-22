package handlers

import (
	"docomo-bike/internal/auth"
	"docomo-bike/internal/libs/httpreq"
	"docomo-bike/internal/libs/httpres"
	"net/http"
)

func HandleAuthorize(authService auth.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody struct {
			UserID        string `json:"userId"`
			PlainPassword string `json:"plainPassword"`
		}
		if err := httpreq.JSONBody(r, &reqBody); err != nil {
			httpres.Error(w, http.StatusInternalServerError, err)
			return
		}

		authResult, err := authService.JWTAuthorize(reqBody.UserID, reqBody.PlainPassword)
		if err != nil {
			httpres.Error(w, http.StatusInternalServerError, err)
			return
		}
		httpres.JSON(w, authResult)
	}
}

func RequireJWT(authService auth.Service) http.HandlerFunc {
	return func(next http.HandlerFunc) {
		next(w, r)
	}
}
