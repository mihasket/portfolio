package server

import (
	"log/slog"
	"net/http"
	"time"

	"portfolio-dotfiles/internal/cache"
	"portfolio-dotfiles/internal/config"
	"portfolio-dotfiles/internal/github"
	"portfolio-dotfiles/internal/logger"
	"portfolio-dotfiles/internal/waybar"
)

type Application struct {
	Config      config.Config
	Logger      *slog.Logger
	HTTPClient  *http.Client
	GithubCache *cache.Cache[*github.Calendar]
	WaybarCache *cache.Cache[*waybar.WaybarData]
}

func NewApplication() (*Application, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: 5 * time.Second}

	app := &Application{
		Config:     cfg,
		Logger:     logger.New(cfg.Env),
		HTTPClient: client,
	}

	app.GithubCache = cache.New(24*time.Hour, func(username string) (*github.Calendar, error) {
		return github.FetchGithubContribution(client, username)
	})

	app.WaybarCache = cache.New(2*time.Hour, func(cdn string) (*waybar.WaybarData, error) {
		return waybar.FetchWaybar(client, app.Config.Waybar.CDN)
	})

	return app, nil
}
