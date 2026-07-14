package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *Application) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(app.logRequests)
	r.Use(commonHeaders)

	r.NotFound(app.notFoundResponse)
	r.MethodNotAllowed(app.methodNotAllowedResponse)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/healthcheck", app.healthcheckHandler)
		r.Get("/github", app.githubContribHandler)
		r.Get("/waybar", app.waybarHandler)
		r.Get("/lastfm", app.lastFMHandler)
	})

	return r
}
