package config

import (
	"fmt"
	"time"
)

type ScheduleConfig struct {
	ShowScores    bool
	ShowStatus    bool
	CompactMode   bool
	SortByTime    bool
	FilterByTeam  string
	MaxGames      int
}

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

type TimeRange struct {
	Start time.Time
	End   time.Time
}

func NewTimeRange(start, end time.Time) *TimeRange {
	return &TimeRange{
		Start: start,
		End:   end,
	}
}

func (tr *TimeRange) Contains(t time.Time) bool {
	return (t.After(tr.Start) || t.Equal(tr.Start)) && (t.Before(tr.End) || t.Equal(tr.End))
}

func (tr *TimeRange) String() string {
	return fmt.Sprintf("%s to %s", 
		tr.Start.Format("3:04 PM"), 
		tr.End.Format("3:04 PM"))
}
