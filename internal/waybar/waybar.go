package waybar

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	return ""
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

	now := time.Now().UTC().Add(2 * time.Hour)

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

type CustomModule struct {
	Format string `json:"format"`
	HtmlID string
}

func (c CustomModule) GetFormat() string {
	return c.Format
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
	DateIcon      CustomModule     `json:"custom/dateicon"`
	LogoIcon      CustomModule     `json:"custom/logo"`
	Window        WindowModule     `json:"hyprland/window"`
	Workspace     WorkspaceModule  `json:"hyprland/workspaces"`
}

type WaybarData struct {
	Layer         string
	Position      string
	Height        int
	Width         int
	ModulesLeft   []Module
	ModulesCenter []Module
	ModulesRight  []Module
}

func buildModuleMap(w *WaybarJSON) map[string]Module {
	w.Clock.HtmlID = "clock"
	w.PulseAudio.HtmlID = "pulseaudio"
	w.Network.HtmlID = "network"
	w.Window.HtmlID = "window"
	w.Workspace.HtmlID = "workspaces"
	w.DateIcon.HtmlID = "custom-dateicon"
	w.LogoIcon.HtmlID = "custom-logo"

	return map[string]Module{
		"clock":               w.Clock,
		"pulseaudio":          w.PulseAudio,
		"network":             w.Network,
		"hyprland/window":     w.Window,
		"hyprland/workspaces": w.Workspace,
		"custom/dateicon":     w.DateIcon,
		"custom/logo":         w.LogoIcon,
	}
}

func sortModules(names []string, moduleMap map[string]Module) []Module {
	modules := make([]Module, 0, len(names))

	for _, name := range names {
		if m, ok := moduleMap[name]; ok {
			modules = append(modules, m)
		}
	}

	return modules
}

func FetchWaybar() (*WaybarData, error) {
	resp, err := http.Get("https://cdn.jsdelivr.net/gh/mihasket/dotfiles@master/.config/waybar/config.jsonc")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status %d fetching waybar config", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var waybarJson WaybarJSON
	err = json.Unmarshal(body, &waybarJson)

	if err != nil {
		return nil, err
	}

	moduleMap := buildModuleMap(&waybarJson)

	waybar := WaybarData{
		Layer:         waybarJson.Layer,
		Position:      waybarJson.Position,
		Height:        waybarJson.Height,
		Width:         waybarJson.Width,
		ModulesLeft:   sortModules(waybarJson.ModulesLeft, moduleMap),
		ModulesCenter: sortModules(waybarJson.ModulesCenter, moduleMap),
		ModulesRight:  sortModules(waybarJson.ModulesRight, moduleMap),
	}

	return &waybar, nil
}
