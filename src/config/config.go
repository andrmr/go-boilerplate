package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

const (
	defaultPrefix = "APP_SETTINGS_"
)

type DbSettings struct {
	Host string `env:"HOST"`
	Port uint32 `env:"PORT"`
}

type Settings struct {
	Name string     `env:"NAME"`
	Age  int        `env:"AGE"`
	Db1  DbSettings `envPrefix:"DB_MANAGEMENT_"`
	Db2  DbSettings `envPrefix:"DB_STATISTICS_"`
}

func parse(prefix string) (Settings, error) {
	s, err := env.ParseAsWithOptions[Settings](env.Options{
		Prefix:          prefix,
		RequiredIfNoDef: true,
	})

	if err != nil {
		return s, fmt.Errorf("config parse: %w", err)
	}

	return s, nil
}

func New() (Settings, error) {
	return parse(defaultPrefix)
}
