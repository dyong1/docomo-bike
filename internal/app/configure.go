package app

import (
	"docomo-bike/internal/config"
	"docomo-bike/internal/libs/docomo/login"
	"docomo-bike/internal/libs/logging"
	"docomo-bike/internal/services/auth"
	"docomo-bike/internal/services/bikebooking"
	"docomo-bike/internal/services/stationlisting"
	"io/ioutil"
	"os"
	"time"

	"github.com/gojektech/heimdall/httpclient"

	"github.com/dgrijalva/jwt-go"

	"github.com/pkg/errors"
)

func (cont *Container) Configure(cfg config.Config) error {
	cont.AppLogger = logging.New("AppLogger", !cfg.Env.IsProd(), false, os.Stdout, !cfg.Env.IsProd())
	cont.HTTPClientLogger = logging.New("HTTP Client", !cfg.Env.IsProd(), false, os.Stdout, !cfg.Env.IsProd())

	jwtConfig, err := jwtConfig(cfg)
	if err != nil {
		return errors.Wrap(err, "")
	}

	{
		cont.JWTAuthService = authService(jwtConfig, cont.HTTPClientLogger)
		cont.BikeBookingService = bikeBookingService(cont.HTTPClientLogger)
		cont.StationListingService = statingListingService(cont.HTTPClientLogger)
	}

	return nil
}

func jwtConfig(cfg config.Config) (auth.JWTConfig, error) {
	secret, err := ioutil.ReadFile(cfg.JWT.SecretFilePath)
	if err != nil {
		return auth.JWTConfig{}, errors.Wrap(err, "")
	}
	return auth.JWTConfig{
		ExpiresIn:     time.Duration(cfg.JWT.ExpiresInSec * 1000 * 1000),
		Issuer:        cfg.JWT.Issuer,
		Secret:        secret,
		SigningMethod: jwt.GetSigningMethod(cfg.JWT.SigningMethod),
	}, nil
}

func authService(jwtConfig auth.JWTConfig, logger logging.Logger) auth.JWTService {
	loginClient := &login.ScrappingClient{
		HTTPClient: httpclient.NewClient(),
		Logger:     logger,
	}
	return &auth.DocomoJWTService{
		JWT:         jwtConfig,
		LoginClient: loginClient,
	}
}

func bikeBookingService(logger logging.Logger) bikebooking.Service {
	return bikebooking.NewService()
}
func statingListingService(logger logging.Logger) stationlisting.Service {
	return stationlisting.NewService()
}
