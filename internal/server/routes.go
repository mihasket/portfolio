package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *Application) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(app.logRequests)
	r.Use(commonHeaders)

	r.Get("/api/v1/healthcheck", app.healthcheckHandler)
	r.Get("/api/v1/github", app.githubContribHandler)
	r.Get("/api/v1/waybar", app.waybarHandler)
	r.Get("/api/v1/lastfm", app.lastFMHandler)

	return r
}
