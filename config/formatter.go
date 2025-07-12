package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mcbk51/sport_schedule/api"
)

// ANSI color codes
const (
	ColorReset   = "\033[0m"
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorMagenta = "\033[35m"
	ColorCyan    = "\033[36m"
	ColorWhite   = "\033[37m"

	ColorBrightRed     = "\033[91m"
	ColorBrightGreen   = "\033[92m"
	ColorBrightYellow  = "\033[93m"
	ColorBrightBlue    = "\033[94m"
	ColorBrightMagenta = "\033[95m"
	ColorBrightCyan    = "\033[96m"
	ColorBrightWhite   = "\033[97m"

	ColorBold      = "\033[1m"
	ColorUnderline = "\033[4m"
	ColorDim       = "\033[2m"
)

// Checks if the terminal supports color output
func supportsColor() bool {
	term := os.Getenv("TERM")
	noColor := os.Getenv("NO_COLOR")

	if noColor != "" {
		return false
	}

	if term == "dumb" || term == "" {
		return false
	}
	return true
}

func colorize(text, colorCode string) string {
	if supportsColor() {
		return colorCode + text + ColorReset
	}
	return text
}

func getStatusColor(status string) string {
	switch {
	case strings.Contains(status, "In Progress") || strings.Contains(status, "Live"):
		return ColorBold + ColorBrightGreen
	case strings.Contains(status, "Delayed") || strings.Contains(status, "Postponed"):
		return ColorBold + ColorBrightYellow
	case strings.Contains(status, "Final") || strings.Contains(status, "Completed"):
		return ColorBrightMagenta
	case strings.Contains(status, "Scheduled"):
		return ColorCyan
	default:
		return ColorWhite
	}
}

// Main print function for the schedule
func PrintSchedule(league string, date time.Time, games []api.Game) {

	header := fmt.Sprintf("📅 Sports Schedule for %s - %s", strings.ToUpper(league), date.Format("Monday, January 2, 2006"))
	fmt.Printf("\n%s\n", colorize(header, ColorBold+ColorBrightWhite))
	fmt.Println(colorize(strings.Repeat("=", 60), ColorBlue))

	// Group games by league
	gamesByLeague := make(map[string][]api.Game)
	for _, game := range games {
		gamesByLeague[game.League] = append(gamesByLeague[game.League], game)
	}

	for leagueName, leagueGames := range gamesByLeague {
		// Adding the specific logo to each sport with colors
		var leagueHeader string
		switch leagueName {
		case "NFL":
			leagueHeader = fmt.Sprintf("🏈 %s (%d games)", leagueName, len(leagueGames))
		case "NBA":
			leagueHeader = fmt.Sprintf("🏀 %s (%d games)", leagueName, len(leagueGames))
		case "MLB":
			leagueHeader = fmt.Sprintf("⚾ %s (%d games)", leagueName, len(leagueGames))
		case "NHL":
			leagueHeader = fmt.Sprintf("🏒 %s (%d games)", leagueName, len(leagueGames))
		default:
			leagueHeader = fmt.Sprintf("%s (%d games)", leagueName, len(leagueGames))
		}

		fmt.Printf("\n%s\n", colorize(leagueHeader, ColorBold+ColorBrightYellow))
		fmt.Println(colorize(strings.Repeat("-", 50), ColorBlue))

		for _, game := range leagueGames {
			// Format time in local timezone
			localTime := game.StartTime.Local()
			timeStr := localTime.Format("3:04 PM")

			matchup := fmt.Sprintf("%s @ %s", game.AwayTeam, game.HomeTeam)

			coloredTime := colorize(timeStr, ColorBold+ColorMagenta)

			coloredMatchup := colorize(matchup, ColorBrightWhite)

			statusColor := getStatusColor(game.Status)

			// Show scores if game has started/finished
			if game.Status != "Scheduled" && (game.HomeScore > 0 || game.AwayScore > 0) {
				statusWithScore := fmt.Sprintf("%s (%d-%d)", game.Status, game.AwayScore, game.HomeScore)
				coloredStatus := colorize(statusWithScore, statusColor)
				fmt.Printf("  %-15s  %-45s  %s\n", coloredTime, coloredMatchup, coloredStatus)
			} else {
				coloredStatus := colorize(game.Status, statusColor)
				fmt.Printf("  %-15s  %-45s  %s\n", coloredTime, coloredMatchup, coloredStatus)
			}
		}
	}

	fmt.Println(colorize("\n"+strings.Repeat("=", 60), ColorBlue))
}

// Adding a filter by team flag
func FilterByTeam(games []api.Game, teamName string) []api.Game {
	var filteredGames []api.Game
	teamLower := strings.ToLower(teamName)

	for _, game := range games {
		homeTeamLower := strings.ToLower(game.HomeTeam)
		awayTeamLower := strings.ToLower(game.AwayTeam)

		if strings.Contains(homeTeamLower, teamLower) || strings.Contains(awayTeamLower, teamLower) {
			filteredGames = append(filteredGames, game)
		}
	}

	return filteredGames
}

// Setting the date flag to work with multiple formats
func ParseDate(dateStr string) (time.Time, error) {
	now := time.Now()

	switch strings.ToLower(dateStr) {
	case "today":
		return now, nil
	case "tomorrow":
		return now.AddDate(0, 0, 1), nil
	default:
		if parsedTime, err := time.Parse("01-02-2006", dateStr); err == nil {
			return parsedTime, nil
		}
		if parsedTime, err := time.Parse("2006-01-02", dateStr); err == nil {
			return parsedTime, nil
		}
		return time.Time{}, fmt.Errorf("invalid date format. Use 'today', 'tomorrow', 'MM-DD-YYYY', or 'YYYY-MM-DD'")
	}
}
