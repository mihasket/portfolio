package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GitHubCalendarResponse struct {
	Total         int         `json:"total"`
	Contributions [][]DayData `json:"contributions"`
}

type DayData struct {
	Date      string `json:"date"`
	Intensity string `json:"intensity"`
	Count     int    `json:"count"`
}

func FetchGithubContribution(username string) (*GitHubCalendarResponse, error) {
	resp, err := http.Get(fmt.Sprintf("https://gh-calendar.rschristian.dev/user/%s", username))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result GitHubCalendarResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
