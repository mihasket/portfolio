package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /projects", app.projects)
	mux.HandleFunc("GET /terminal", app.terminal)

	mux.HandleFunc("GET /templates/zsh", app.zshPromptTemplate)

	return app.recoverPanic(app.logRequests(commonHeaders(mux)))
}
