package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTP `env:"http"`
		//Log    `yaml:"logger"`
		DB `env:"postgres"`
	}

	HTTP struct {
		Port string `env-required:"true" env:"HTTP_PORT"`
	}
	//Log struct {
	//	Level string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
	//}
	DB struct {
		DbName   string `env-required:"true" env:"POSTGRES_DB"`
		Host     string `env-required:"true" env:"POSTGRES_HOST"`
		Port     string `env-required:"true" env:"POSTGRES_PORT"`
		SslMode  string `env-required:"true" env:"POSTGRES_SSL"`
		User     string `env-required:"true" env:"POSTGRES_USER"`
		Password string `env-required:"true" env:"POSTGRES_PASSWORD"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(".env", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
