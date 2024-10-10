package app

import "github.com/tehrelt/eda/orders/internal/config"

type App struct {
	cfg *config.Config
}

func New(cfg *config.Config) *App {
	return &App{
		cfg: cfg,
	}
}
