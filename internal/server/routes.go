package server

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
)

func spaHandler(dir string) http.Handler {
	fileServer := http.FileServer(http.Dir(dir))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(dir, filepath.Clean("/"+r.URL.Path))

		info, err := os.Stat(path)
		if err != nil || info.IsDir() {
			// Unknown path or directory -> let React Router handle it
			w.Header().Set("Cache-Control", "no-cache")
			http.ServeFile(w, r, filepath.Join(dir, "index.html"))
			return
		}

		// Vite fingerprints everything in /assets, so it can be cached forever
		if strings.HasPrefix(r.URL.Path, "/assets/") {
			w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
		}
		fileServer.ServeHTTP(w, r)
	})
}

func (app *application) registerRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(app.recoverPanic)
	r.Use(app.logRequests)
	r.Use(commonHeaders)

	r.Route("/api/v1", func(r chi.Router) {
		r.NotFound(app.notFoundResponse)
		r.MethodNotAllowed(app.methodNotAllowedResponse)

		r.Get("/healthcheck", app.healthcheckHandler)
		r.Get("/github", app.githubContribHandler)
		r.Get("/lastfm", app.lastFMHandler)
	})

	r.Handle("/*", spaHandler("./frontend/dist"))

	return r
}
