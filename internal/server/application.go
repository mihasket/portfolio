package server

import (
	"log/slog"
	"net/http"
	"time"

	"portfolio-dotfiles/internal/config"
	"portfolio-dotfiles/internal/github"
	"portfolio-dotfiles/internal/logger"
)

type Application struct {
	Config      config.Config
	Logger      *slog.Logger
	HTTPClient  *http.Client
	GithubCache *github.Cache
}

func NewApplication() (*Application, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: 5 * time.Second}

	return &Application{
		Config:      cfg,
		Logger:      logger.New(cfg.Env),
		HTTPClient:  client,
		GithubCache: github.NewCache(client, 24*time.Hour),
	}, nil
}
