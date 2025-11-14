package config

import (
	"log/slog"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort    string `env:"PORT" envDefault:"8080"`
	PG_DBHost     string `env:"DB_PG_HOST" envDefault:"postgres"`
	PG_DBUser     string `env:"DB_PG_USER" envDefault:"webuser"`
	PG_DBPassword string `env:"DB_PG_PASSWORD" envDefault:"1111"`
	PG_DBName     string `env:"DB_PG_NAME" envDefault:"webdev"`
	PG_DBSSLMode  string `env:"DB_PG_SSLMODE" envDefault:"disable"`
	PG_PORT       string `env:"DB_PG_PORT" envDefault:"5432"`
}

func MustLoad() *Config {

	if err := godotenv.Load(); err != nil {
		slog.Debug("Failed to load .env file", "error", err)
	} else {
		slog.Info("Loaded configuration from .env file")
	}

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		slog.Error("Failed to parse environment variables", "error", err)
		panic("configuration error: " + err.Error())
	}

	return &cfg
}
