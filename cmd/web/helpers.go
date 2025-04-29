package main

import (
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

func sortModules(m []string, waybar WaybarJSON) []Module {
	var modules []Module

	// Tote "clock" spremenljivke lahko verjetno das v neki enum, ker se bo ponavljalo
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
			waybar.Workspace.Buttons = []int{1, 2, 3, 4, 5}
			modules = append(modules, waybar.Workspace)
		}
	}

	return modules
}

func (app *application) generateTemplateData(waybar WaybarJSON) WaybarTmpl {
	var data WaybarTmpl

	data.ModulesLeft = sortModules(waybar.ModulesLeft, waybar)
	data.ModulesCenter = sortModules(waybar.ModulesCenter, waybar)
	data.ModulesRight = sortModules(waybar.ModulesRight, waybar)

	return data
}
