package server

import (
	"log/slog"

	"portfolio-dotfiles/internal/config"
	"portfolio-dotfiles/internal/logger"
)

type Application struct {
	Config config.Config
	Logger *slog.Logger
}

func NewApplication() (*Application, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	return &Application{
		Config: cfg,
		Logger: logger.New(cfg.Env),
	}, nil
}
