package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Calendar struct {
	Total         int         `json:"total"`
	Contributions [][]DayData `json:"contributions"`
}

type DayData struct {
	Date      string `json:"date"`
	Intensity string `json:"intensity"`
	Count     int    `json:"count"`
}

func FetchGithubContribution(client *http.Client, username string) (*Calendar, error) {
	resp, err := client.Get(fmt.Sprintf("https://gh-calendar.rschristian.dev/user/%s", username))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result Calendar
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
