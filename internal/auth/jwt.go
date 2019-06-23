package auth

import (
	"docomo-bike/internal/docomo"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	JWTAuthorize(userID string, plainPassword string) (*JWTAuthResult, error)
	VerifyJWTToken(tokenString string) error
}

type JWTAuthResult struct {
	TokenString string
}

type JWTConfig struct {
	ExpiresIn     time.Duration
	Issuer        string
	PrivateKey    []byte
	PublicKey     []byte
	SigningMethod jwt.SigningMethod
}
type jwtClaims struct {
	jwt.StandardClaims

	SessionKey string
}

type DocomoJWTAuthService struct {
	JWT          JWTConfig
	DocomoClient docomo.Client
}

func (s *DocomoJWTAuthService) JWTAuthorize(userID string, plainPassword string) (*JWTAuthResult, error) {
	sessionKey, err := s.DocomoClient.Login(userID, plainPassword)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Failed to get session key [userID=%s]", userID))
	}

	token := jwt.NewWithClaims(s.JWT.SigningMethod, &jwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.JWT.ExpiresIn).Unix(),
			Issuer:    s.JWT.Issuer,
		},
		SessionKey: sessionKey,
	})
	ss, err := token.SignedString(s.JWT.PrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Failed to sign claims [userId=%s]", userID))
	}

	return &JWTAuthResult{
		TokenString: ss,
	}, nil
}

func (s *DocomoJWTAuthService) VerifyJWTToken(tokenString string) error {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return s.JWT.PublicKey, nil
	})
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Invalid token string [value=%s]", tokenString))
	}
	return nil
}
