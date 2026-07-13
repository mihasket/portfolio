package server

import (
	"net/http"
	"portfolio-dotfiles/internal/lastfm"
	"portfolio-dotfiles/internal/waybar"
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
	gh, err := app.GithubCache.GetContributions("mihasket")
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
	waybar, err := waybar.FetchWaybar()
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
	lastFM, err := lastfm.FetchLastTrack(app.Config.LastFM)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, lastFM, nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}
