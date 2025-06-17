package main

import (
	"html/template"
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

	calendar, err := app.fetchGithubContribution()
	if err != nil {
		app.serverError(w, r, err)
	}

	app.render(w, r, http.StatusOK, "home.tmpl", app.generateTemplateData(*waybar, lastTrack, calendar))
}

func (app *application) projects(w http.ResponseWriter, r *http.Request) {
	waybar, err := app.fetchWaybar()
	if err != nil {
		app.serverError(w, r, err)
	}

	app.render(w, r, http.StatusOK, "projects.tmpl", app.generateTemplateData(*waybar, nil, nil))
}

func (app *application) terminal(w http.ResponseWriter, r *http.Request) {
	waybar, err := app.fetchWaybar()
	if err != nil {
		app.serverError(w, r, err)
	}

	app.render(w, r, http.StatusOK, "terminal.tmpl", app.generateTemplateData(*waybar, nil, nil))
}

func (app *application) zshPromptTemplate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	ts, err := template.ParseFiles("./ui/html/components/zsh.tmpl")
	if err != nil {
		app.serverError(w, r, err)
	}

	err = ts.ExecuteTemplate(w, "zsh", nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}
