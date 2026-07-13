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

type Identifiable struct {
	ID string `json:"id"`
}

type SimpleModule struct {
	Identifiable
	Format string `json:"format"`
}

func (s SimpleModule) GetFormat() string {
	return s.Format
}

type Module interface {
	GetFormat() string
}

type NetworkModule struct {
	Identifiable
	Format string `json:"format-ethernet"`
}

func (n NetworkModule) GetFormat() string {
	return n.Format
}

type ClockModule struct {
	Identifiable
	Format string `json:"format"`
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

type PulseAudioModule struct {
	Identifiable
	Format      string         `json:"format"`
	FormatIcons map[string]any `json:"format-icons"`
}

func (p PulseAudioModule) GetFormat() string {
	return ""
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
	DateIcon      SimpleModule     `json:"custom/dateicon"`
	LogoIcon      SimpleModule     `json:"custom/logo"`
	Window        SimpleModule     `json:"hyprland/window"`
	Workspace     SimpleModule     `json:"hyprland/workspaces"`
}

type WaybarData struct {
	Layer         string   `json:"layer"`
	Position      string   `json:"position"`
	Height        int      `json:"height"`
	Width         int      `json:"width"`
	ModulesLeft   []Module `json:"modulesLeft"`
	ModulesCenter []Module `json:"modulesCenter"`
	ModulesRight  []Module `json:"modulesRight"`
}

func buildModuleMap(w *WaybarJSON) map[string]Module {
	w.Clock.ID = "clock"
	w.PulseAudio.ID = "pulseaudio"
	w.Network.ID = "network"
	w.Window.ID = "window"
	w.Workspace.ID = "workspaces"
	w.DateIcon.ID = "custom-dateicon"
	w.LogoIcon.ID = "custom-logo"

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

func FetchWaybar(client *http.Client, cdn string) (*WaybarData, error) {
	resp, err := client.Get(cdn)
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
