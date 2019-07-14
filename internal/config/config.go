package config

import (
	"github.com/jinzhu/configor"
	"github.com/pkg/errors"
)

type Config struct {
	Env string `env:"ENV"`

	HTTPServer

	JWT
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
	if err := configor.Load(cfg); err != nil {
		return errors.Wrap(err, "Failed to parse envs")
	}
	return nil
}
