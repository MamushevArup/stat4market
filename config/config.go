package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ClickHouse ClickHouse
	HTTP       HTTP `yaml:"http"`
}

type ClickHouse struct {
	Host     string `env:"CLICKHOUSE_HOST" env-default:"localhost"`
	Port     int    `env:"CLICKHOUSE_PORT" env-default:"8123"`
	Database string `env:"CLICKHOUSE_DB"`
	Password string `env:"CLICKHOUSE_PASSWORD"`
	User     string `env:"CLICKHOUSE_USER" env-default:"user"`
}

type HTTP struct {
	Port          int    `yaml:"port"`
	Host          string `yaml:"host"`
	Timeout       string `yaml:"timeout"`
	IdleTimeout   string `yaml:"idle_timeout"`
	HeaderTimeout string `yaml:"header_timeout"`
}

func New(configPath string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(configPath, cfg)
	if err != nil {
		return nil, fmt.Errorf("error reading config file path:%v: %w", configPath, err)
	}

	return cfg, nil
}
