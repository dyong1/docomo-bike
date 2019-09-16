package app

import (
	"context"
	"docomo-bike/internal/config"
	"docomo-bike/internal/libs/docomo/getstation"
	"docomo-bike/internal/libs/docomo/login"
	"docomo-bike/internal/libs/env"
	"docomo-bike/internal/libs/logging"
	"docomo-bike/internal/services/auth"
	"docomo-bike/internal/services/booking"
	"docomo-bike/internal/services/listing"
	"io/ioutil"
	"os"
	"time"

	"github.com/gojektech/heimdall/httpclient"

	"github.com/dgrijalva/jwt-go"

	"github.com/pkg/errors"
)

func (cont *Container) Configure(cfg config.Config) error {
	cont.AppLogger = logging.New("AppLogger", !env.IsProd(cfg.Env), false, os.Stdout, !env.IsProd(cfg.Env))
	cont.HTTPClientLogger = logging.New("HTTP Client", !env.IsProd(cfg.Env), false, os.Stdout, !env.IsProd(cfg.Env))

	cont.DocomoClients.Login = &login.ScrappingClient{
		HTTPClient: httpclient.NewClient(),
		Logger:     cont.HTTPClientLogger,
	}
	cont.DocomoClients.GetStation = &getstation.ScrappingClient{
		HTTPClient: httpclient.NewClient(),
		Logger:     cont.HTTPClientLogger,
	}

	jwtConfig, err := jwtConfig(cfg)
	if err != nil {
		return errors.Wrap(err, "")
	}
	{
		cont.JWTAuthService = auth.NewService(jwtConfig, cont.DocomoClients.Login)
		cont.BikeBookingService = booking.NewService()
		cont.StationListingService = listing.NewService(cont.DocomoClients.GetStation)
	}

	return nil
}

func jwtConfig(cfg config.Config) (auth.JWTConfig, error) {
	secret, err := ioutil.ReadFile(cfg.JWT.SecretFilePath)
	if err != nil {
		return auth.JWTConfig{}, errors.Wrap(err, "")
	}
	return auth.JWTConfig{
		ExpiresIn:     time.Duration(cfg.JWT.ExpiresInSec) * time.Second,
		Issuer:        cfg.JWT.Issuer,
		Secret:        secret,
		SigningMethod: jwt.GetSigningMethod(cfg.JWT.SigningMethod),
	}, nil
}

func (c *Container) Shutdown(ctx context.Context) error {
	return nil
}
