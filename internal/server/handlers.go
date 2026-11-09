package server

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
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

func (app *application) githubContribHandler(w http.ResponseWriter, r *http.Request) {
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

func (app *application) lastFMHandler(w http.ResponseWriter, r *http.Request) {
	url := app.lastFMURL()

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
