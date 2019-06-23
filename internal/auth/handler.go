package auth

import (
	"docomo-bike/internal/libs/httpreq"
	"docomo-bike/internal/libs/httpres"
	"fmt"
	"net/http"
)

func HandleAuthorize(jwtService JWTService) http.HandlerFunc {
	type requestBody struct {
		UserID        string `json:"userId"`
		PlainPassword string `json:"plainPassword"`
	}
	type responseBody struct {
		Token string `json:"token"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody requestBody
		if err := httpreq.JSONBody(r, &reqBody); err != nil {
			httpres.Error(w, http.StatusInternalServerError, err)
			return
		}

		authResult, err := jwtService.JWTAuthorize(reqBody.UserID, reqBody.PlainPassword)
		if err != nil {
			httpres.Error(w, http.StatusInternalServerError, err)
			return
		}
		httpres.JSON(w, responseBody{
			Token: authResult.TokenString,
		})
	}
}

func RequireAuth(jwtService JWTService) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			hh := r.Header["Authorization"]
			if len(hh) != 2 {
				httpres.Error(w, http.StatusUnauthorized, fmt.Errorf("Invalid authorization"))
				return
			}
			if err := jwtService.VerifyJWTToken(hh[1]); err != nil {
				httpres.Error(w, http.StatusUnauthorized, err)
				return
			}
			next(w, r)
		}
	}
}
