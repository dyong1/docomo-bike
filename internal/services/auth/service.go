package auth

import (
	login "docomo-bike/internal/libs/docomo/login"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	Authorize(userID string, plainPassword string) (*JWTAuthResult, error)
	AuthFromToken(tokenString string) (*Auth, error)
}

func NewService(jwt JWTConfig, loginClient login.Client) JWTService {
	return &DocomoJWTService{
		JWT:         jwt,
		LoginClient: loginClient,
	}
}

type JWTAuthResult struct {
	UserID      string
	TokenString string
}

type JWTConfig struct {
	ExpiresIn     time.Duration
	Issuer        string
	Secret        []byte
	SigningMethod jwt.SigningMethod
}
type Auth struct {
	UserID     string
	SessionKey string
}

type jwtClaims struct {
	jwt.StandardClaims

	UserID     string `json:"userId"`
	SessionKey string `json:"sessionKey"`
}
