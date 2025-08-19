package main

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	waybar, err := app.fetchWaybar()
	if err != nil {
		app.serverError(w, r, err)
	}

	// TODO: Change the date, because lastFM returns GMT+0
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

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		app.serverError(w, r, err)
	}

	var zsh ZshData
	err = json.Unmarshal(body, &zsh)
	if err != nil {
		app.serverError(w, r, err)
	}

	ts, err := template.ParseFiles("./ui/html/components/terminalZsh.tmpl")
	if err != nil {
		app.serverError(w, r, err)
	}

	err = ts.ExecuteTemplate(w, "terminalZsh", zsh.ID)
	if err != nil {
		app.serverError(w, r, err)
	}
}
