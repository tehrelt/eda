package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

type Config struct {
	App struct {
		Env  string `env:"APP_ENV" env-required:"true" env-default:"local"`
		Port int    `env:"APP_PORT" env-required:"true" env-default:"8080"`
	}

	Database struct {
		Host string `env:"DATABASE_HOST" env-required:"true" env-default:"localhost"`
		Port int    `env:"DATABASE_PORT" env-required:"true" env-default:"5432"`
		Name string `env:"DATABASE_NAME" env-required:"true" env-default:"postgres"`
		User string `env:"DATABASE_USER" env-required:"true" env-default:"postgres"`
		Pass string `env:"DATABASE_PASS" env-required:"true" env-default:"postgres"`
	}
}

func New() *Config {
	config := new(Config)

	if err := cleanenv.ReadEnv(config); err != nil {
		panic(fmt.Sprintf("failed to read env: %v", err))
	}

	slog.SetDefault(setupLogger(config.App.Env))

	return config
}

func setupLogger(env string) *slog.Logger {

	switch env {
	case envDev:
		return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
}
