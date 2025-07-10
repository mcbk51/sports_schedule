package config

import (
	"time"
)

type Config struct {
	League     string
	Date       time.Time
	TimeZone   string
	OutputFile string
	Verbose    bool
}

func DefaultConfig() *Config {
	return &Config{
		League:     "all",
		Date:       time.Now(),
		TimeZone:   "Local",
		OutputFile: "",
		Verbose:    false,
	}
}

func SupportedLeagues() []string {
	return []string{"nfl", "nba", "nhl", "mlb", "all"}
}

func IsValidLeague(league string) bool {
	supported := SupportedLeagues()
	for _, l := range supported {
		if l == league {
			return true
		}
	}
	return false
}
