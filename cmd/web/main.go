package main

import (
	"os"

	"portfolio-dotfiles/internal/server"
)

func main() {
	srv, app, err := server.NewServer()
	if err != nil {
		os.Exit(1)
	}

	app.Logger.Info("starting server", "addr", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		app.Logger.Error("server failed", "err", err)
		os.Exit(1)
	}
}
