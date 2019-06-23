package app

import (
	"docomo-bike/internal/auth"
	"docomo-bike/internal/config"
	"docomo-bike/internal/docomo"
	"io/ioutil"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/google/logger"
	"github.com/pkg/errors"
)

func (a *App) Configure(cfg config.Config) error {
	a.Logger = logger.Init("APP", !cfg.Env.IsProd(), false, os.Stdout)

	jwtConfig, err := jwtConfig(cfg)
	if err != nil {
		return errors.Wrap(err, "")
	}
	docomoClient := &docomo.ScrappingClient{}
	authService := &auth.DocomoJWTAuthService{
		JWT:          jwtConfig,
		DocomoClient: docomoClient,
	}

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
	privateKey, err := ioutil.ReadFile(cfg.JWTPrivateKeyFilePath)
	if err != nil {
		return auth.JWTConfig{}, errors.Wrap(err, "")
	}
	publicKey, err := ioutil.ReadFile(cfg.JWTPublicKeyFilePath)
	if err != nil {
		return auth.JWTConfig{}, errors.Wrap(err, "")
	}
	return auth.JWTConfig{
		ExpiresIn:     time.Duration(cfg.JWTExpiresInSec * 1000 * 1000),
		Issuer:        cfg.JWTIssuer,
		PrivateKey:    privateKey,
		PublicKey:     publicKey,
		SigningMethod: jwt.GetSigningMethod(cfg.JWTSigningMethod),
	}, nil
}
