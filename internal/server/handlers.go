package server

import (
	"fmt"
	"net/http"
)

func (app *Application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": app.Config.Env,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.Logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

func (app *Application) githubContribHandler(w http.ResponseWriter, r *http.Request) {
	gh, err := app.GithubCache.Get("mihasket")
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, gh, nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *Application) waybarHandler(w http.ResponseWriter, r *http.Request) {
	waybar, err := app.WaybarCache.Get(app.Config.Waybar.CDN)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, waybar, nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *Application) lastFMHandler(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf(
		"https://ws.audioscrobbler.com/2.0/?method=user.getrecenttracks&user=%s&api_key=%s&format=json&limit=1",
		app.Config.LastFM.Username, app.Config.LastFM.APIKey,
	)

	lastFM, err := app.LastFMCache.Get(url)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, lastFM, nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}
