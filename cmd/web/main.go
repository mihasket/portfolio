package main

import (
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type lastFmEnv struct {
	LAST_FM_API_KEY  string
	LAST_FM_USERNAME string
}

type config struct {
	addr   string
	env    string
	lastFm lastFmEnv
}

type application struct {
	config        config
	logger        *slog.Logger
	templateCache map[string]*template.Template
}

func main() {
	var cfg config
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	if cfg.env != "production" {
		err := godotenv.Load(".env")

		if err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
	}

	cfg.lastFm.LAST_FM_API_KEY = os.Getenv("LAST_FM_API_KEY")
	cfg.lastFm.LAST_FM_USERNAME = os.Getenv("LAST_FM_USERNAME")

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		config:        cfg,
		logger:        logger,
		templateCache: templateCache,
	}

	srv := &http.Server{
		Addr:         cfg.addr,
		Handler:      app.routes(),
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Info("starting server", "addr", srv.Addr)

	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
