package http

import (
	"context"
	"docomo-bike/internal/services/auth"
	"encoding/json"
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
		dec := json.NewDecoder(r.Body)
		if err := dec.Decode(&reqBody); err != nil {
			badRequest(w, err)
			return
		}

		authResult, err := authService.Authorize(reqBody.UserID, reqBody.PlainPassword)
		if err != nil {
			internalServerError(w, err)
			return
		}
		jsonres(w, responseBody{
			Token: authResult.TokenString,
		})
	}
}

func UseAuth(authService auth.JWTService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			hh := r.Header["Authorization"]
			if len(hh) != 2 {
				unauthorized(w, fmt.Errorf("Invalid authorization"))
				return
			}
			auth, err := authService.AuthFromToken(hh[1])
			if err != nil {
				unauthorized(w, err)
				return
			}
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), AuthContextKey{}, auth)))
		})
	}
}

type AuthContextKey struct{}
