package internal

import (
	"fmt"
	"strings"
	"time"

	"github.com/mcbk51/sport_schedule/api"
)

func PrintSchedule(league string, date time.Time, games []api.Game) {
	fmt.Printf("\n📅 Sports Schedule for %s - %s\n", strings.ToUpper(league), date.Format("Monday, January 2, 2006"))
	fmt.Println(strings.Repeat("=", 60))

	// Group games by league
	gamesByLeague := make(map[string][]api.Game)
	for _, game := range games {
		gamesByLeague[game.League] = append(gamesByLeague[game.League], game)
	}

	for leagueName, leagueGames := range gamesByLeague {
		switch leagueName {
		case "NFL":
			fmt.Printf("\n🏈 %s (%d games)\n", leagueName, len(leagueGames))
			fmt.Println(strings.Repeat("-", 50))
		case "NBA":
			fmt.Printf("\n🏀 %s (%d games)\n", leagueName, len(leagueGames))
			fmt.Println(strings.Repeat("-", 50))
		case "MLB":
			fmt.Printf("\n⚾ %s (%d games)\n", leagueName, len(leagueGames))
			fmt.Println(strings.Repeat("-", 50))
		case "NHL":
			fmt.Printf("\n🏒 %s (%d games)\n", leagueName, len(leagueGames))
			fmt.Println(strings.Repeat("-", 50))
		}
		for _, game := range leagueGames {
			// Format time in local timezone
			localTime := game.StartTime.Local()
			timeStr := localTime.Format("3:04 PM")

			// Create the matchup string
			matchup := fmt.Sprintf("%s @ %s", game.AwayTeam, game.HomeTeam)

			// Show scores if game has started/finished
			if game.Status != "Scheduled" && (game.HomeScore > 0 || game.AwayScore > 0) {
				fmt.Printf("  %-6s  %-35s  %s (%d-%d)\n",
					timeStr, matchup, game.Status, game.AwayScore, game.HomeScore)
			} else {
				fmt.Printf("  %-6s  %-35s  %s\n",
					timeStr, matchup, game.Status)
			}
		}
	}

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Printf("Total games: %d\n", len(games))
}

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
