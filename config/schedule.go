package config

import (
	"fmt"
	"time"
)

// ScheduleConfig holds schedule-specific configuration
type ScheduleConfig struct {
	ShowScores    bool
	ShowStatus    bool
	CompactMode   bool
	SortByTime    bool
	FilterByTeam  string
	MaxGames      int
}

// DefaultScheduleConfig returns default schedule configuration
func DefaultScheduleConfig() *ScheduleConfig {
	return &ScheduleConfig{
		ShowScores:    true,
		ShowStatus:    true,
		CompactMode:   false,
		SortByTime:    true,
		FilterByTeam:  "",
		MaxGames:      0, // 0 means no limit
	}
}

// TimeRange represents a time range for filtering games
type TimeRange struct {
	Start time.Time
	End   time.Time
}

// NewTimeRange creates a new time range
func NewTimeRange(start, end time.Time) *TimeRange {
	return &TimeRange{
		Start: start,
		End:   end,
	}
}

// Contains checks if a time falls within the range
func (tr *TimeRange) Contains(t time.Time) bool {
	return (t.After(tr.Start) || t.Equal(tr.Start)) && (t.Before(tr.End) || t.Equal(tr.End))
}

// String returns a string representation of the time range
func (tr *TimeRange) String() string {
	return fmt.Sprintf("%s to %s", 
		tr.Start.Format("3:04 PM"), 
		tr.End.Format("3:04 PM"))
}
