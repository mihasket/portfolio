package lastfm

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type recentTracksResponse struct {
	RecentTracks struct {
		Track []rawTrack `json:"track"`
	} `json:"recenttracks"`
}

type rawTrack struct {
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
		Uts json.Number `json:"uts"`
	} `json:"date"`
}

type Track struct {
	Name           string `json:"name"`
	Artist         string `json:"artist"`
	Album          string `json:"album"`
	ImageUrlLarge  string `json:"imageUrlLarge"`
	ImageUrlMedium string `json:"imageUrlMedium"`
	Timestamp      int64  `json:"timestamp"`
}

func FetchLastTrack(client *http.Client, url string) (*Track, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var raw recentTracksResponse
	err = json.Unmarshal(body, &raw)
	if err != nil {
		return nil, err
	}

	if len(raw.RecentTracks.Track) == 0 {
		return nil, fmt.Errorf("lastfm: no recent tracks returned")
	}

	return mapTrack(raw.RecentTracks.Track[0]), nil
}

func mapTrack(t rawTrack) *Track {
	track := &Track{
		Name:           t.Name,
		Artist:         t.Artist.Name,
		Album:          t.Album.Text,
		ImageUrlLarge:  extraLargeImage(t.Image),
		ImageUrlMedium: mediumImage(t.Image),
	}

	if ts, err := t.Date.Uts.Int64(); err == nil {
		track.Timestamp = ts
	}

	return track
}

func extraLargeImage(images []struct {
	Size string `json:"size"`
	Text string `json:"#text"`
}) string {
	for _, img := range images {
		if img.Size == "extralarge" {
			return img.Text
		}
	}
	return ""
}

func mediumImage(images []struct {
	Size string `json:"size"`
	Text string `json:"#text"`
}) string {
	for _, img := range images {
		if img.Size == "medium" {
			return img.Text
		}
	}
	return ""
}
