package server

import (
	"fmt"
	"net/http"
)

func (app *Application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.Env,
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *Application) githubContribHandler(w http.ResponseWriter, r *http.Request) {
	gh, err := app.githubCache.Get("mihasket")
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"github": gh}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *Application) waybarHandler(w http.ResponseWriter, r *http.Request) {
	waybar, err := app.waybarCache.Get(app.config.Waybar.CDN)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"waybar": waybar}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *Application) lastFMHandler(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf(
		"https://ws.audioscrobbler.com/2.0/?method=user.getrecenttracks&user=%s&api_key=%s&format=json&limit=1",
		app.config.LastFM.Username, app.config.LastFM.APIKey,
	)

	lastFM, err := app.lastFMCache.Get(url)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"lastFM": lastFM}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
