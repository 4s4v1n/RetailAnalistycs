package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	APP      App
	HTTP     Http
	Postgres Postgres
	Jwt      Jwt
}

type App struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type Http struct {
	Port string `yaml:"port"`
}

type AdminCredentials struct {
	Login    string `env:"ADMIN_LOGIN"`
	Password string `env:"ADMIN_PASSWORD"`
}

type VisitorCredentials struct {
	Login    string `env:"VISITOR_LOGIN"`
	Password string `env:"VISITOR_PASSWORD"`
}

type DSN struct {
	Admin   string `env:"ADMIN_DSN"`
	Visitor string `env:"VISITOR_DSN"`
}

type Roles struct {
	AdminCredentials
	VisitorCredentials
}

type Postgres struct {
	DSN
	Roles
	MaxOpenConns    int           `yaml:"max_open_conns"`
	MaxIdleConns    int           `yaml:"max_idle_conns"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
}

type Jwt struct {
	TokenTTL   time.Duration `yaml:"token_ttl"`
	RefreshTTL time.Duration `yaml:"refresh_ttl"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("parse .env: %w", err)
	}

	if err := cleanenv.ReadConfig("config/config.yml", cfg); err != nil {
		return nil, fmt.Errorf(".yml: %w", err)
	}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, fmt.Errorf(".env: %w", err)
	}
	return cfg, nil
}
