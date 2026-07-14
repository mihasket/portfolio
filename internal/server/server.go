package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

func NewServer() (*http.Server, *Application, error) {
	app, err := NewApplication()
	if err != nil {
		return nil, nil, err
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.Addr),
		Handler:      app.registerRoutes(),
		ErrorLog:     slog.NewLogLogger(app.Logger.Handler(), slog.LevelError),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return srv, app, nil
}
