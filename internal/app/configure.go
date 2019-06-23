package app

import (
	"docomo-bike/internal/auth"
	"docomo-bike/internal/booking"
	"docomo-bike/internal/config"
	"docomo-bike/internal/httphandler"
	"docomo-bike/internal/libs/docomo/login"
	"docomo-bike/internal/libs/logger"
	"docomo-bike/internal/listing"
	"io/ioutil"
	"os"
	"time"

	"github.com/gojektech/heimdall/httpclient"

	"github.com/dgrijalva/jwt-go"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/pkg/errors"
)

func (a *App) Configure(cfg config.Config) error {
	a.Logger = logger.New("App", !cfg.Env.IsProd(), false, os.Stdout, !cfg.Env.IsProd())

	jwtConfig, err := jwtConfig(cfg)
	if err != nil {
		return errors.Wrap(err, "")
	}

	httpClientLogger := logger.New("HTTP Client", !cfg.Env.IsProd(), false, os.Stdout, !cfg.Env.IsProd())
	authService := authService(jwtConfig, httpClientLogger)
	bookingService := bookingService(httpClientLogger)
	listingReservationService := listingReservationService(httpClientLogger)

	{
		router := chi.NewRouter()
		router.Use(middleware.Logger)
		router.Use(middleware.Recoverer)

		router.Post("/auth", httphandler.HandleAuthorize(authService))
		router.Route("/me", func(me chi.Router) {
			me.Use(httphandler.UseAuth(authService))
			me.Post("/booking", httphandler.HandleBook(bookingService))
			me.Get("/booking", httphandler.HandleGetReservations(listingReservationService))
			me.Get("/booking/{bookingID}", httphandler.HandleGetReservation(listingReservationService))
		})

		a.Router = router
	}

	return nil
}

func jwtConfig(cfg config.Config) (auth.JWTConfig, error) {
	secret, err := ioutil.ReadFile(cfg.JWTSecretFilePath)
	if err != nil {
		return auth.JWTConfig{}, errors.Wrap(err, "")
	}
	return auth.JWTConfig{
		ExpiresIn:     time.Duration(cfg.JWTExpiresInSec * 1000 * 1000),
		Issuer:        cfg.JWTIssuer,
		Secret:        secret,
		SigningMethod: jwt.GetSigningMethod(cfg.JWTSigningMethod),
	}, nil
}

func authService(jwtConfig auth.JWTConfig, logger *logger.Logger) auth.JWTService {
	loginClient := &login.ScrappingClient{
		HTTPClient: httpclient.NewClient(),
		Logger:     logger,
	}
	return &auth.DocomoJWTService{
		JWT:         jwtConfig,
		LoginClient: loginClient,
	}
}

func bookingService(logger *logger.Logger) booking.Service {
	return &booking.DocomoService{}
}
func listingReservationService(logger *logger.Logger) listing.ReservationService {
	return &listing.DocomoReservationService{}
}
