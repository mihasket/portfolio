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

type Module interface{}

type SimpleModule struct {
	Format string `json:"format"`
}

type NetworkModule struct {
	Format string `json:"format-ethernet"`
}

func (n NetworkModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Format string `json:"format"`
	}{
		Format: n.Format,
	})
}

type ClockModule struct {
	SimpleModule
}

func (c *ClockModule) ResolveFormat() {
	var (
		icon   string
		layout string
	)

	start := strings.Index(c.Format, "{")
	end := strings.Index(c.Format, "}")

	if start != -1 && end != -1 && start < end {
		icon = strings.TrimSpace(c.Format[:start])
		layout = strings.TrimSpace(c.Format[start+2 : end])
	}

	now := time.Now().UTC().Add(2 * time.Hour)
	c.Format = strings.Join([]string{
		icon,
		timefmt.Format(now, layout),
	}, " ")
}

type PulseAudioModule struct {
	SimpleModule
	FormatIcons map[string]any `json:"format-icons"`
}

func (p PulseAudioModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Format string `json:"format"`
	}{
		Format: getIcon(p.FormatIcons, "headphone"),
	})
}

type WorkspaceModule struct {
	SimpleModule
	FormatIcons map[string]any `json:"format-icons"`
}

func (w WorkspaceModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Format       string `json:"format"`
		FormatActive string `json:"format-active"`
	}{
		Format:       getIcon(w.FormatIcons, "1"),
		FormatActive: getIcon(w.FormatIcons, "active"),
	})
}

func getIcon(icons map[string]any, key string) string {
	if value, ok := icons[key].(string); ok {
		return value
	}

	return ""
}

type WaybarJSON struct {
	Layer         string           `json:"layer"`
	Height        int              `json:"height"`
	ModulesLeft   []string         `json:"modules-left"`
	ModulesCenter []string         `json:"modules-center"`
	ModulesRight  []string         `json:"modules-right"`
	Spacing       int              `json:"spacing"`
	PulseAudio    PulseAudioModule `json:"pulseaudio"`
	Network       NetworkModule    `json:"network"`
	Clock         ClockModule      `json:"clock"`
	DateIcon      SimpleModule     `json:"custom/dateicon"`
	LogoIcon      SimpleModule     `json:"custom/logo"`
	Workspace     WorkspaceModule  `json:"hyprland/workspaces"`
}

type WaybarData struct {
	Layer  string `json:"layer"`
	Height int    `json:"height"`

	Modules map[string]Module `json:"modules"`
	Layout  Layout            `json:"layout"`
}

type Layout struct {
	Left   []string `json:"left"`
	Center []string `json:"center"`
	Right  []string `json:"right"`
}

func buildModules(w *WaybarJSON) map[string]Module {
	w.Clock.ResolveFormat()

	return map[string]Module{
		"clock":               w.Clock,
		"network":             w.Network,
		"pulseaudio":          w.PulseAudio,
		"hyprland/workspaces": w.Workspace,
		"custom/dateicon":     w.DateIcon,
		"custom/logo":         w.LogoIcon,
	}
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

	modules := buildModules(&waybarJson)

	return &WaybarData{
		Layer:  waybarJson.Layer,
		Height: waybarJson.Height,

		Modules: modules,

		Layout: Layout{
			Left:   waybarJson.ModulesLeft,
			Center: waybarJson.ModulesCenter,
			Right:  waybarJson.ModulesRight,
		},
	}, nil
}
