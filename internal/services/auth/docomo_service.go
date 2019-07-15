package auth

import (
	login "docomo-bike/internal/libs/docomo/login"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

type DocomoJWTService struct {
	JWT         JWTConfig
	LoginClient login.Client
}

func (s *DocomoJWTService) Authorize(userID string, plainPassword string) (*JWTAuthResult, error) {
	sessionKey, err := s.LoginClient.Login(userID, plainPassword)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Failed to get session key [userID=%s]", userID))
	}

	token := jwt.NewWithClaims(s.JWT.SigningMethod, &jwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.JWT.ExpiresIn).Unix(),
			Issuer:    s.JWT.Issuer,
		},
		UserID:     userID,
		SessionKey: sessionKey,
	})
	ss, err := token.SignedString(s.JWT.Secret)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Failed to sign claims [userId=%s]", userID))
	}

	return &JWTAuthResult{
		UserID:      userID,
		TokenString: ss,
	}, nil
}

func (s *DocomoJWTService) AuthFromToken(tokenString string) (*Auth, error) {
	var claims jwtClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return s.JWT.Secret, nil
	})
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Failed to parse token [value=%s]", tokenString))
	}
	if !token.Valid {
		return nil, errors.Wrap(err, fmt.Sprintf("Invalid token string [value=%s]", tokenString))
	}
	return &Auth{
		SessionKey: claims.SessionKey,
		UserID:     claims.UserID,
	}, nil
}
