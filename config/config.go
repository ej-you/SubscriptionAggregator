// Package config provides loading config data from
// external sources like env variables.
package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Server
		DB
	}

	Server struct {
		Name            string        `env:"SERVER_NAME" env-default:"Subscription Aggregator API"`
		Port            string        `env:"SERVER_PORT" env-default:"8000"`
		ShutdownTimeout time.Duration `env:"SERVER_SHUTDOWN_TIMEOUT" env-default:"5s"`
	}

	DB struct {
		MigrationsURL string `env:"MIGRATIONS_URL" env-default:"file://migrations"`
		User          string `env-required:"true" env:"POSTGRES_USER"`
		Password      string `env-required:"true" env:"POSTGRES_PASSWORD"`
		Host          string `env-required:"true" env:"POSTGRES_HOST"`
		Port          string `env-required:"true" env:"POSTGRES_PORT"`
		Name          string `env-required:"true" env:"POSTGRES_DB"`
		ConnString    string
		ConnURL       string
	}
)

// New returns app config loaded from ENV-vars.
func New() (*Config, error) {
	cfg := &Config{}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, fmt.Errorf("load env variables: %w", err)
	}
	cfg.DB.ConnString = fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable connect_timeout=10",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	)
	cfg.DB.ConnURL = fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable&connect_timeout=10",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	)
	return cfg, nil
}
