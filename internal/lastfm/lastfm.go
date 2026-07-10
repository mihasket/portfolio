package lastfm

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"portfolio-dotfiles/internal/config"
)

// Needed to check for empty date
type LastFmDate struct {
	Uts  json.Number `json:"uts"`
	Text string      `json:"#text"`
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
			Name string     `json:"name"`
			Date LastFmDate `json:"date"`
		} `json:"track"`
	} `json:"recenttracks"`
}

func FetchLastTrack(lastFM config.LastFM) (*RecentTracksResponse, error) {
	url := fmt.Sprintf(
		"https://ws.audioscrobbler.com/2.0/?method=user.getrecenttracks&user=%s&api_key=%s&format=json&limit=1",
		lastFM.Username, lastFM.APIKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var lastTrack RecentTracksResponse
	err = json.Unmarshal(body, &lastTrack)
	if err != nil {
		return nil, err
	}

	if (LastFmDate{}) != lastTrack.RecentTracks.Track[0].Date {
		timestamp, err := lastTrack.RecentTracks.Track[0].Date.Uts.Int64()
		if err != nil {
			return nil, err
		}

		t := time.Unix(timestamp, 0).UTC()
		loc := time.FixedZone("GMT+2", 2*60*60)
		localTime := t.In(loc)

		lastTrack.RecentTracks.Track[0].Date.Text = localTime.Format("02 Jan 2006, 15:04")
	}

	return &lastTrack, nil
}
