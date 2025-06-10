package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	waybar, err := app.fetchWaybar()
	if err != nil {
		app.serverError(w, r, err)
	}

	lastTrack, err := app.fetchLastTrack()
	if err != nil {
		app.serverError(w, r, err)
	}

	app.render(w, r, http.StatusOK, "home.tmpl", app.generateTemplateData(*waybar, lastTrack))
}

func (app *application) projects(w http.ResponseWriter, r *http.Request) {
	waybar, err := app.fetchWaybar()
	if err != nil {
		app.serverError(w, r, err)
	}

	app.render(w, r, http.StatusOK, "projects.tmpl", app.generateTemplateData(*waybar, nil))
}
