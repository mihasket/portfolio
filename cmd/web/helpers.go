package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data templateData) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, r, err)
		return
	}

	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(status)

	buf.WriteTo(w)
}

func (app *application) fetchWaybar() (*WaybarJSON, error) {
	res, err := http.Get("https://cdn.jsdelivr.net/gh/mihasket/dotfiles@master/.config/waybar/config.jsonc")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var waybar WaybarJSON
	err = json.Unmarshal(body, &waybar)

	if err != nil {
		return nil, err
	}

	return &waybar, nil
}

func (app *application) fetchLastTrack() (*RecentTracksResponse, error) {
	url := fmt.Sprintf(
		"https://ws.audioscrobbler.com/2.0/?method=user.getrecenttracks&user=%s&api_key=%s&format=json&limit=1",
		app.config.lastFm.LAST_FM_USERNAME, app.config.lastFm.LAST_FM_API_KEY,
	)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var lastTrack RecentTracksResponse
	err = json.Unmarshal(body, &lastTrack)
	if err != nil {
		return nil, err
	}

	return &lastTrack, nil
}

func sortModules(m []string, waybar WaybarJSON) []Module {
	var modules []Module

	for _, v := range m {
		switch v {
		case "clock":
			waybar.Clock.HtmlID = v
			modules = append(modules, waybar.Clock)
		case "pulseaudio":
			waybar.PulseAudio.HtmlID = v
			modules = append(modules, waybar.PulseAudio)
		case "network":
			waybar.Network.HtmlID = v
			modules = append(modules, waybar.Network)
		case "hyprland/window":
			waybar.Window.HtmlID = "window"
			modules = append(modules, waybar.Window)
		case "hyprland/workspaces":
			waybar.Workspace.HtmlID = "workspaces"
			modules = append(modules, waybar.Workspace)
		}
	}

	return modules
}

func (app *application) generateTemplateData(waybar WaybarJSON, lastTrack *RecentTracksResponse) templateData {
	var data templateData

	data.ModulesLeft = sortModules(waybar.ModulesLeft, waybar)
	data.ModulesCenter = sortModules(waybar.ModulesCenter, waybar)
	data.ModulesRight = sortModules(waybar.ModulesRight, waybar)

	if lastTrack == nil {
		return data
	}

	data.LastFM = *lastTrack

	if len(data.LastFM.RecentTracks.Track[0].Album.Text) >= 24 {
		data.LastFM.RecentTracks.Track[0].Album.Text =
			data.LastFM.RecentTracks.Track[0].Album.Text[0:24] + "..."
	}

	return data
}
