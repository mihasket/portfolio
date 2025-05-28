package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	waybar, err := app.fetchWaybar()
	if err != nil {
		app.serverError(w, r, err)
	}

	app.render(w, r, http.StatusOK, "home.tmpl", app.generateTemplateData(*waybar))
}

func (app *application) projects(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	waybar, err := app.fetchWaybar()
	if err != nil {
		app.serverError(w, r, err)
	}

	app.render(w, r, http.StatusOK, "projects.tmpl", app.generateTemplateData(*waybar))
}
