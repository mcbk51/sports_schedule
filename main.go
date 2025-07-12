package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mcbk51/sport_schedule/api"
	"github.com/mcbk51/sport_schedule/config"
)

func main() {
	// Setting up the main flags of the function and errors
	league := flag.String("league", "all", "League to show games for (nfl, nba, nhl, mlb, or all)")
	date := flag.String("date", "today", "Date (today, tomorrow, or MM-DD-YYYY)")
	team := flag.String("team", "", "Team name to filter by (e.g., 'Lakers', 'Giants', 'Yankees')")
	flag.Parse()

	parsedDate, err := config.ParseDate(*date)
	if err != nil {
		fmt.Println("Invalid date:", err)
		os.Exit(1)
	}

	games, err := api.GetGames(*league, parsedDate)
	if err != nil {
		fmt.Printf("Error fetching games: %v\n", err)
		os.Exit(1)
	}

	if *team != "" {
		games = config.FilterByTeam(games, *team)
	}

	if len(games) == 0 {
		if *team != "" {
			fmt.Printf("No games found for team '%s' on %s.\n", team, parsedDate.Format("January 2, 2006"))
		} else {
			fmt.Printf("No games found for %s on %s.\n", *league, parsedDate.Format("January 2, 2006"))
		}
		return
	}

	config.PrintSchedule(*league, parsedDate, games)
}
