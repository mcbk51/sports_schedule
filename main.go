package main

import (
	"fmt"
	"os"
	"flag"
	"time"
	"strings"
	"github.com/mcbk51/sport_schedule/api"
	"github.com/mcbk51/sport_schedule/internal"
)

func main() {
	league := flag.String("league", "all", "League to show games for (nfl, nba, nhl, mlb, or all)")
	date := flag.String("date", "today", "Date (today, tomorrow, or MM-DD-YYYY)")
	flag.Parse()

	parsedDate, err := parseDate(*date)
	if err != nil {
		fmt.Println("Invalid date:", err)
		os.Exit(1)
	}

	games, err := api.GetGames(*league, parsedDate)
	if err != nil {
		fmt.Printf("Error fetching games: %v\n", err)
		os.Exit(1)
	}

	if len(games) == 0 {
		fmt.Printf("No games found for %s on %s.\n", *league, parsedDate.Format("January 2, 2006"))
		return
	}

	internal.PrintSchedule(*league, parsedDate, games)
}

// parseDate parses the date string and returns a time.Time
func parseDate(dateStr string) (time.Time, error) {
	now := time.Now()
	
	switch strings.ToLower(dateStr) {
	case "today":
		return now, nil
	case "tomorrow":
		return now.AddDate(0, 0, 1), nil
	default:
		// Try to parse MM-DD-YYYY format
		if parsedTime, err := time.Parse("01-02-2006", dateStr); err == nil {
			return parsedTime, nil
		}
		// Try to parse YYYY-MM-DD format
		if parsedTime, err := time.Parse("2006-01-02", dateStr); err == nil {
			return parsedTime, nil
		}
		return time.Time{}, fmt.Errorf("invalid date format. Use 'today', 'tomorrow', 'MM-DD-YYYY', or 'YYYY-MM-DD'")
	}
}
