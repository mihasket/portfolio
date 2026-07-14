package server

import (
	"log/slog"
	"net/http"
	"time"

	"portfolio-dotfiles/internal/cache"
	"portfolio-dotfiles/internal/config"
	"portfolio-dotfiles/internal/github"
	"portfolio-dotfiles/internal/lastfm"
	"portfolio-dotfiles/internal/logger"
	"portfolio-dotfiles/internal/waybar"
)

type application struct {
	config      config.Config
	logger      *slog.Logger
	httpClient  *http.Client
	githubCache *cache.Cache[*github.Calendar]
	waybarCache *cache.Cache[*waybar.WaybarData]
	lastFMCache *cache.Cache[*lastfm.Track]
}

func newApplication() (*application, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: 5 * time.Second}

	appLogger := logger.New(cfg.Env)
	slog.SetDefault(appLogger)

	app := &application{
		config:     cfg,
		logger:     appLogger,
		httpClient: client,
	}

	app.githubCache = cache.New(24*time.Hour, func(username string) (*github.Calendar, error) {
		return github.FetchGithubContribution(client, username)
	})

	app.waybarCache = cache.New(2*time.Hour, func(cdn string) (*waybar.WaybarData, error) {
		return waybar.FetchWaybar(client, app.config.Waybar.CDN)
	})

	app.lastFMCache = cache.New(30*time.Second, func(url string) (*lastfm.Track, error) {
		return lastfm.FetchLastTrack(client, url)
	})

	return app, nil
}
