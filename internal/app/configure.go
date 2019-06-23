package app

import (
	"docomo-bike/internal/auth"
	"docomo-bike/internal/config"
	"docomo-bike/internal/docomo/login"
	"docomo-bike/internal/libs/logger"
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
	appLogger := logger.New("App", !cfg.Env.IsProd(), false, os.Stdout, !cfg.Env.IsProd())
	a.Logger = appLogger

	jwtConfig, err := jwtConfig(cfg)
	if err != nil {
		return errors.Wrap(err, "")
	}

	authService := authService(jwtConfig, appLogger)

	{
		router := chi.NewRouter()
		router.Use(middleware.Logger)
		router.Use(middleware.Recoverer)

		router.Post("/auth", auth.HandleAuthorize(authService))

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

func authService(jwtConfig auth.JWTConfig, logger *logger.Logger) *auth.DocomoJWTAuthService {
	loginClient := &login.ScrappingClient{
		HTTPClient: httpclient.NewClient(),
		Logger:     logger,
	}
	return &auth.DocomoJWTAuthService{
		JWT:         jwtConfig,
		LoginClient: loginClient,
	}
}
