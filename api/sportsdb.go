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
}

// ESPNResponse represents the structure of ESPN API response
type ESPNResponse struct {
	Events []struct {
		Name        string `json:"name"`
		ShortName   string `json:"shortName"`
		Date        string `json:"date"`
		Status      struct {
			Type struct {
				Description string `json:"description"`
			} `json:"type"`
		} `json:"status"`
		Competitions []struct {
			Competitors []struct {
				Team struct {
					DisplayName string `json:"displayName"`
					Abbreviation string `json:"abbreviation"`
				} `json:"team"`
				HomeAway string `json:"homeAway"`
				Score    string `json:"score"`
			} `json:"competitors"`
		} `json:"competitions"`
	} `json:"events"`
}

// GetGames fetches games for the specified league and date
func GetGames(league string, date time.Time) ([]Game, error) {
	var games []Game
	
	leagues := []string{"nfl", "nba", "nhl", "mlb"}
	if league != "all" {
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

// fetchGamesForLeague fetches games for a specific league
func fetchGamesForLeague(league string, date time.Time) ([]Game, error) {
	dateStr := date.Format("20060102")
	
	// ESPN API endpoint for different leagues
	var url string
	switch league {
	case "nfl":
		url = fmt.Sprintf("https://site.api.espn.com/apis/site/v2/sports/football/nfl/scoreboard?dates=%s", dateStr)
	case "nba":
		url = fmt.Sprintf("https://site.api.espn.com/apis/site/v2/sports/basketball/nba/scoreboard?dates=%s", dateStr)
	case "nhl":
		url = fmt.Sprintf("https://site.api.espn.com/apis/site/v2/sports/hockey/nhl/scoreboard?dates=%s", dateStr)
	case "mlb":
		url = fmt.Sprintf("https://site.api.espn.com/apis/site/v2/sports/baseball/mlb/scoreboard?dates=%s", dateStr)
	default:
		return nil, fmt.Errorf("unsupported league: %s", league)
	}

	resp, err := http.Get(url)
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

	var espnResp ESPNResponse
	if err := json.Unmarshal(body, &espnResp); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	var games []Game
	for _, event := range espnResp.Events {
		if len(event.Competitions) == 0 || len(event.Competitions[0].Competitors) < 2 {
			continue
		}

		comp := event.Competitions[0]
		var homeTeam, awayTeam string
		var homeScore, awayScore int

		for _, competitor := range comp.Competitors {
			if competitor.HomeAway == "home" {
				homeTeam = competitor.Team.DisplayName
				if competitor.Score != "" {
					fmt.Sscanf(competitor.Score, "%d", &homeScore)
				}
			} else {
				awayTeam = competitor.Team.DisplayName
				if competitor.Score != "" {
					fmt.Sscanf(competitor.Score, "%d", &awayScore)
				}
			}
		}

		startTime, _ := time.Parse(time.RFC3339, event.Date)
		
		game := Game{
			HomeTeam:  homeTeam,
			AwayTeam:  awayTeam,
			StartTime: startTime,
			League:    strings.ToUpper(league),
			Status:    event.Status.Type.Description,
			HomeScore: homeScore,
			AwayScore: awayScore,
		}
		games = append(games, game)
	}

	return games, nil
}
