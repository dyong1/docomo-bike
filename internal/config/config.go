package config

import (
	"strings"

	"github.com/caarlos0/env"
	"github.com/pkg/errors"
)

type Config struct {
	Env Environment `env:"ENV"`

	HTTPServer

	JWT
}
type Environment string

func (e Environment) IsProd() bool {
	return strings.ToLower(string(e)) == "production"
}
func (e Environment) IsDev() bool {
	return strings.ToLower(string(e)) == "development"
}

type HTTPServer struct {
	Host string `env:"HTTP_SERVER_HOST"`
	Port string `env:"HTTP_SERVER_PORT"`
}

type JWT struct {
	ExpiresInSec   int    `env:"JWT_EXPIRES_IN_SEC"`
	Issuer         string `env:"JWT_ISSUER"`
	SecretFilePath string `env:"JWT_SECRET_FILE_PATH"`
	SigningMethod  string `env:"JWT_SIGNING_METHOD"`
}

func (cfg *Config) Load() error {
	if err := env.Parse(cfg); err != nil {
		return errors.Wrap(err, "Failed to parse envs")
	}
	return nil
}
