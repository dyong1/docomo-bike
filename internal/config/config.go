package config

import (
	"strings"

	"github.com/caarlos0/env"
	"github.com/pkg/errors"
)

type Config struct {
	Env            Environment `env:"ENV,required"`
	HTTPServerHost string      `env:"HTTP_SERVER_HOST,required"`
	HTTPServerPort string      `env:"HTTP_SERVER_PORT,required"`

	JWTExpiresInSec   int    `env:"JWT_EXPIRES_IN_SEC,required"`
	JWTIssuer         string `env:"JWT_ISSUER,required"`
	JWTSecretFilePath string `env:"JWT_SECRET_FILE_PATH,required"`
	JWTSigningMethod  string `env:"JWT_SIGNING_METHOD,required"`
}

func (cfg *Config) Load() error {
	if err := env.Parse(cfg); err != nil {
		return errors.Wrap(err, "Failed to parse envs")
	}
	return nil
}

type Environment string

func (e Environment) IsProd() bool {
	return strings.ToLower(string(e)) == "production"
}
func (e Environment) IsDev() bool {
	return strings.ToLower(string(e)) == "development"
}
