package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Game represents a sports game
type Game struct {
	HomeTeam    string    `json:"home_team"`
	AwayTeam    string    `json:"away_team"`
	StartTime   time.Time `json:"start_time"`
	League      string    `json:"league"`
	Status      string    `json:"status"`
	HomeScore   int       `json:"home_score"`
	AwayScore   int       `json:"away_score"`
	GameID      string    `json:"game_id,omitempty"`
}

// MySportsFeedsResponse represents the structure of MySportsFeeds API response
type MySportsFeedsResponse struct {
	Games []struct {
		Schedule struct {
			ID              int    `json:"id"`
			StartTime       string `json:"startTime"`
			EndedTime       string `json:"endedTime"`
			AwayTeam        Team   `json:"awayTeam"`
			HomeTeam        Team   `json:"homeTeam"`
			PlayedStatus    string `json:"playedStatus"`
			DelayedOrPostponedReason string `json:"delayedOrPostponedReason"`
		} `json:"schedule"`
		Score struct {
			AwayScoreTotal int `json:"awayScoreTotal"`
			HomeScoreTotal int `json:"homeScoreTotal"`
		} `json:"score"`
	} `json:"games"`
}

type Team struct {
	ID           int    `json:"id"`
	Abbreviation string `json:"abbreviation"`
	City         string `json:"city"`
	Name         string `json:"name"`
}

const (
	// MySportsFeeds API base URL
	baseURL = "https://api.mysportsfeeds.com/v2.1/pull"
	
	// You need to get API key from MySportsFeeds
	// Format: "username:password" (base64 encoded for basic auth)
	apiKey = "YOUR_API_KEY_HERE"
)

// GetGames fetches games for the specified league and date
func GetGames(league string, date time.Time) ([]Game, error) {
	var games []Game
	
	// MySportsFeeds supported leagues
	leagues := []string{"nfl", "nba", "nhl", "mlb"}
	if league != "all" {
		// Validate league
		found := false
		for _, l := range leagues {
			if strings.ToLower(league) == l {
				found = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("unsupported league: %s. Supported leagues: %v", league, leagues)
		}
		leagues = []string{strings.ToLower(league)}
	}

	for _, l := range leagues {
		leagueGames, err := fetchGamesForLeague(l, date)
		if err != nil {
			fmt.Printf("Warning: Could not fetch games for %s: %v\n", l, err)
			continue
		}
		games = append(games, leagueGames...)
	}

	return games, nil
}

// fetchGamesForLeague fetches games for a specific league from MySportsFeeds
func fetchGamesForLeague(league string, date time.Time) ([]Game, error) {
	// Get the season year - MySportsFeeds uses different season formats
	season := getSeason(league, date)
	dateStr := date.Format("20060102")
	
	// MySportsFeeds API endpoint
	url := fmt.Sprintf("%s/%s/%s/games.json?fordate=%s", baseURL, league, season, dateStr)

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	// Create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add basic auth header
	if apiKey != "YOUR_API_KEY_HERE" {
		req.Header.Set("Authorization", "Basic "+apiKey)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "SportsApp/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 401 {
		return nil, fmt.Errorf("authentication failed - check your API key")
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var msfResp MySportsFeedsResponse
	if err := json.Unmarshal(body, &msfResp); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return parseGamesResponse(msfResp, league)
}

// parseGamesResponse parses the MySportsFeeds response into our Game struct
func parseGamesResponse(resp MySportsFeedsResponse, league string) ([]Game, error) {
	var games []Game
	
	for _, game := range resp.Games {
		schedule := game.Schedule
		score := game.Score
		
		// Parse start time
		startTime, err := time.Parse(time.RFC3339, schedule.StartTime)
		if err != nil {
			fmt.Printf("Warning: Could not parse start time for game %d: %v\n", schedule.ID, err)
			startTime = time.Now()
		}

		// Determine game status
		status := "Scheduled"
		if schedule.PlayedStatus == "COMPLETED" {
			status = "Final"
		} else if schedule.PlayedStatus == "LIVE" {
			status = "In Progress"
		} else if schedule.PlayedStatus == "POSTPONED" {
			status = "Postponed"
		}

		// Build team names
		homeTeam := fmt.Sprintf("%s %s", schedule.HomeTeam.City, schedule.HomeTeam.Name)
		awayTeam := fmt.Sprintf("%s %s", schedule.AwayTeam.City, schedule.AwayTeam.Name)

		gameObj := Game{
			HomeTeam:  homeTeam,
			AwayTeam:  awayTeam,
			StartTime: startTime,
			League:    strings.ToUpper(league),
			Status:    status,
			HomeScore: score.HomeScoreTotal,
			AwayScore: score.AwayScoreTotal,
			GameID:    fmt.Sprintf("%d", schedule.ID),
		}
		games = append(games, gameObj)
	}

	return games, nil
}

// getSeason determines the season string for MySportsFeeds API
func getSeason(league string, date time.Time) string {
	year := date.Year()
	
	switch league {
	case "nfl":
		// NFL season runs from September to February
		if date.Month() >= 9 {
			return fmt.Sprintf("%d-regular", year)
		}
		return fmt.Sprintf("%d-regular", year-1)
	case "nba":
		// NBA season runs from October to April
		if date.Month() >= 10 {
			return fmt.Sprintf("%d-%d-regular", year, year+1)
		}
		return fmt.Sprintf("%d-%d-regular", year-1, year)
	case "nhl":
		// NHL season runs from October to April
		if date.Month() >= 10 {
			return fmt.Sprintf("%d-%d-regular", year, year+1)
		}
		return fmt.Sprintf("%d-%d-regular", year-1, year)
	case "mlb":
		// MLB season runs from April to September
		return fmt.Sprintf("%d-regular", year)
	default:
		return fmt.Sprintf("%d-regular", year)
	}
}

// GetPlayerStats fetches player statistics for a specific game
func GetPlayerStats(league string, gameID string) (interface{}, error) {
	season := getSeason(league, time.Now())
	url := fmt.Sprintf("%s/%s/%s/games/%s/playerstats.json", baseURL, league, season, gameID)

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	if apiKey != "YOUR_API_KEY_HERE" {
		req.Header.Set("Authorization", "Basic "+apiKey)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var stats interface{}
	if err := json.Unmarshal(body, &stats); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return stats, nil
}

// GetStandings fetches league standings
func GetStandings(league string) (interface{}, error) {
	season := getSeason(league, time.Now())
	url := fmt.Sprintf("%s/%s/%s/standings.json", baseURL, league, season)

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	if apiKey != "YOUR_API_KEY_HERE" {
		req.Header.Set("Authorization", "Basic "+apiKey)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var standings interface{}
	if err := json.Unmarshal(body, &standings); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return standings, nil
}
