package http

import (
	"context"
	"docomo-bike/internal/libs/docomo/login"
	"docomo-bike/internal/services/auth"
	"encoding/json"
	"net/http"
	"strings"
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
			badRequest(w, err.Error())
			return
		}

		authResult, err := authService.Authorize(reqBody.UserID, reqBody.PlainPassword)
		if err != nil {
			internalServerError(w, err.Error())
			return
		}
		jsonres(w, responseBody{
			Token: authResult.TokenString,
		})
	}
}

func UseAuth(authService auth.JWTService, loginClient login.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			hh := strings.Split(r.Header.Get("Authorization"), " ")
			if len(hh) != 2 {
				unauthorized(w, "Invalid authorization")
				return
			}
			auth, err := authService.AuthFromToken(hh[1])
			if err != nil {
				unauthorized(w, err.Error())
				return
			}
			tr, err := loginClient.Test(auth.UserID, auth.SessionKey)
			if err != nil {
				unauthorized(w, err.Error())
				return
			}
			if !tr {
				unauthorized(w, "Session in the token has been expired.")
				return
			}
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), authContextKey{}, auth)))
		})
	}
}

type authContextKey struct{}

func authFromContext(ctx context.Context) *auth.Auth {
	return ctx.Value(authContextKey{}).(*auth.Auth)
}
func userIDFromContext(ctx context.Context) string {
	auth := ctx.Value(authContextKey{}).(*auth.Auth)
	return auth.UserID
}
