package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type LastFM struct {
	APIKey   string
	Username string
}

type Config struct {
	Addr   int
	Env    string
	LastFM LastFM
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}

	return fallback
}

func Load() (Config, error) {
	var cfg Config

	cfg.Env = getEnv("ENV", "development")

	if cfg.Env != "production" {
		if err := godotenv.Load(".env"); err != nil {
			return Config{}, err
		}
	}

	addr, err := strconv.Atoi(getEnv("ADDR", "4000"))
	if err != nil {
		return Config{}, err
	}
	cfg.Addr = addr

	cfg.LastFM.APIKey = os.Getenv("LAST_FM_API_KEY")
	cfg.LastFM.Username = os.Getenv("LAST_FM_USERNAME")

	return cfg, nil
}
