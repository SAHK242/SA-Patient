package config

import (
	"go.uber.org/fx"
	cfg "patient/config"
)

func NewConfig() cfg.Config {
	return cfg.NewViper(
		cfg.WithRequiredConfig([]string{
			"DB_HOST",
			"DB_PORT",
			"DB_USER",
			"DB_PASS",
		}),
	)
}

var Module = fx.Provide(NewConfig)
