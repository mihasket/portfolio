package main

import (
	"strings"
	"time"

	"github.com/itchyny/timefmt-go"
)

type Module interface {
	GetFormat() string
}

type PulseAudioModule struct {
	Format      string         `json:"format"`
	FormatIcons map[string]any `json:"format-icons"`
	HtmlID      string
}

func (p PulseAudioModule) GetFormat() string {
	//  TODO: Error handling
	//  What to return if there was an error, maybe

	if def, ok := p.FormatIcons["default"].([]any); ok && len(def) > 0 {
		if icon, ok := def[0].(string); ok {
			return icon + " 25"
		}
	}

	return "  25"
}

type NetworkModule struct {
	Format string `json:"format-ethernet"`
	HtmlID string
}

func (n NetworkModule) GetFormat() string {
	return n.Format
}

type ClockModule struct {
	Format string `json:"format"`
	HtmlID string
}

func (c ClockModule) GetFormat() string {
	var (
		clockIcon string
		clockTime string
	)

	start := strings.Index(c.Format, "{")
	end := strings.Index(c.Format, "}")

	if start != -1 && end != -1 && start < end {
		clockIcon = strings.TrimSpace(c.Format[:start])
		clockTime = strings.TrimSpace(c.Format[start+2 : end])
	}

	now := time.Now()
	clockTime = timefmt.Format(now, clockTime)

	return strings.Join([]string{clockIcon, clockTime}, " ")
}

type WindowModule struct {
	Format string `json:"format"`
	HtmlID string
}

func (w WindowModule) GetFormat() string {
	return w.Format
}

type WorkspaceModule struct {
	Format string `json:"format"`
	HtmlID string
}

func (w WorkspaceModule) GetFormat() string {
	return w.Format
}

type WaybarJSON struct {
	Layer         string           `json:"layer"`
	Position      string           `json:"position"`
	Height        int              `json:"height"`
	Width         int              `json:"width"`
	ModulesLeft   []string         `json:"modules-left"`
	ModulesCenter []string         `json:"modules-center"`
	ModulesRight  []string         `json:"modules-right"`
	Spacing       int              `json:"spacing"`
	PulseAudio    PulseAudioModule `json:"pulseaudio"`
	Network       NetworkModule    `json:"network"`
	Clock         ClockModule      `json:"clock"`
	Window        WindowModule     `json:"hyprland/window"`
	Workspace     WorkspaceModule  `json:"hyprland/workspaces"`
}

type templateData struct {
	Layer          string
	Position       string
	Height         int
	Width          int
	ModulesLeft    []Module
	ModulesCenter  []Module
	ModulesRight   []Module
	LastFM         RecentTracksResponse
	GithubCalendar GitHubCalendarResponse
}

type RecentTracksResponse struct {
	RecentTracks struct {
		Track []struct {
			Artist struct {
				Name string `json:"#text"`
			} `json:"artist"`
			Image []struct {
				Size string `json:"size"`
				Text string `json:"#text"`
			} `json:"image"`
			Album struct {
				Text string `json:"#text"`
			} `json:"album"`
			Name string `json:"name"`
			Date struct {
				Uts  string `json:"uts"`
				Text string `json:"#text"`
			} `json:"date"`
		} `json:"track"`
	} `json:"recenttracks"`
}

type GitHubCalendarResponse struct {
	Total         int         `json:"total"`
	Contributions [][]DayData `json:"contributions"`
}

type DayData struct {
	Date      string `json:"date"`
	Intensity string `json:"intensity"`
	Count     int    `json:"count"`
}
