package auth

import (
	"context"
	"docomo-bike/internal/libs/httpreq"
	"docomo-bike/internal/libs/httpres"
	"fmt"
	"net/http"
)

func HandleAuthorize(authService JWTAuthService) http.HandlerFunc {
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

func UseAuth(authService JWTAuthService) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
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
			next(w, r.WithContext(context.WithValue(r.Context(), AuthContextKey{}, auth)))
		}
	}
}

type AuthContextKey struct{}
