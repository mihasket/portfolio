package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	url := "https://cdn.jsdelivr.net/gh/mihasket/dotfiles@master/.config/waybar/config.jsonc"
	res, err := http.Get(url)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	var waybar WaybarJSON
	err = json.Unmarshal(body, &waybar)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.render(w, r, http.StatusOK, "home.tmpl", app.generateTemplateData(waybar))
}
