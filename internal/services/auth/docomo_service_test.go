package auth

import (
	login "docomo-bike/internal/libs/docomo/login"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestAuthFromToken(t *testing.T) {
	jwt := JWTConfig{
		ExpiresIn:     time.Hour,
		Issuer:        "auth",
		Secret:        []byte("secret"),
		SigningMethod: jwt.SigningMethodHS512,
	}
	loginClientMock := &login.ClientMock{
		LoginFunc: func(id string, password string) (string, error) {
			return "sessionKey", nil
		},
	}
	serv := NewService(jwt, loginClientMock)

	ar, err := serv.Authorize("userId", "password")
	assert.NoError(t, err)

	auth, err := serv.AuthFromToken(ar.TokenString)
	assert.NoError(t, err)
	assert.Equal(t, "userId", auth.UserID)
	assert.Equal(t, "sessionKey", auth.SessionKey)
}
