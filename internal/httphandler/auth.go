package httphandler

import (
	"context"
	"docomo-bike/internal/auth"
	"docomo-bike/internal/libs/httpreq"
	"docomo-bike/internal/libs/httpres"
	"fmt"
	"net/http"
)

func HandleAuthorize(authService auth.JWTService) http.HandlerFunc {
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
			httpres.Error(w, http.StatusBadRequest, err)
			return
		}

		authResult, err := authService.Authorize(reqBody.UserID, reqBody.PlainPassword)
		if err != nil {
			httpres.Error(w, http.StatusInternalServerError, err)
			return
		}
		httpres.JSON(w, responseBody{
			Token: authResult.TokenString,
		})
	}
}

func UseAuth(authService auth.JWTService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			hh := r.Header["Authorization"]
			if len(hh) != 2 {
				httpres.Error(w, http.StatusUnauthorized, fmt.Errorf("Invalid authorization"))
				return
			}
			auth, err := authService.AuthFromToken(hh[1])
			if err != nil {
				httpres.Error(w, http.StatusUnauthorized, err)
				return
			}
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), AuthContextKey{}, auth)))
		})
	}
}

type AuthContextKey struct{}
