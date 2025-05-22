package main

import (
	"encoding/json"
	"html/template"
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

	// fmt.Printf("Decoded: %+v\n", waybar)
	templateData := app.generateTemplateData(waybar)

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
		"./ui/html/pages/waybar.tmpl",
		"./ui/html/pages/zsh.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", templateData)
	if err != nil {
		app.serverError(w, r, err)
	}
}
