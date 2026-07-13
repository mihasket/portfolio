package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
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

func fetchGithubContribution(client *http.Client, username string) (*GitHubCalendarResponse, error) {
	resp, err := client.Get(fmt.Sprintf("https://gh-calendar.rschristian.dev/user/%s", username))
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

type cacheEntry struct {
	data      *GitHubCalendarResponse
	expiresAt time.Time
}

type Cache struct {
	mu     sync.RWMutex
	ttl    time.Duration
	client *http.Client
	data   map[string]cacheEntry
}

func NewCache(client *http.Client, ttl time.Duration) *Cache {
	return &Cache{
		ttl:    ttl,
		client: client,
		data:   make(map[string]cacheEntry),
	}
}

func (c *Cache) GetContributions(username string) (*GitHubCalendarResponse, error) {
	c.mu.RLock()
	entry, ok := c.data[username]
	c.mu.RUnlock()

	if ok && time.Now().Before(entry.expiresAt) {
		return entry.data, nil
	}

	result, err := fetchGithubContribution(c.client, "mihasket")
	if err != nil {
		return nil, err
	}

	c.mu.Lock()
	c.data[username] = cacheEntry{
		data:      result,
		expiresAt: time.Now().Add(c.ttl),
	}
	c.mu.Unlock()

	return result, nil
}
