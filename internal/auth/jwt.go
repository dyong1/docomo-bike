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

type DocomoAuthService struct {
	jwt          JWTConfig
	docomoClient docomo.Client
}

func (s *DocomoAuthService) JWTAuthorize(userID string, plainPassword string) (*JWTAuthResult, error) {
	sessionKey, err := s.docomoClient.Login(userID, plainPassword)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Failed to get session key [userID=%s]", userID))
	}

	token := jwt.NewWithClaims(s.jwt.SigningMethod, &jwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.jwt.ExpiresIn).Unix(),
			Issuer:    s.jwt.Issuer,
		},
		SessionKey: sessionKey,
	})
	ss, err := token.SignedString(s.jwt.PrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Failed to sign claims [userId=%s]", userID))
	}

	return &JWTAuthResult{
		TokenString: ss,
	}, nil
}

func (s *DocomoAuthService) VerifyJWTToken(tokenString string) error {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return s.jwt.PublicKey, nil
	})
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Invalid token string [value=%s]", tokenString))
	}
	return nil
}
