package utils

import (
	"fmt"
	"time"
)

func SecondsUntilThisHour() time.Duration {
	now := time.Now()
	end := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 59, 59, 0, now.Location())
	d, _ := time.ParseDuration(fmt.Sprintf("%ds", end.Unix()-now.Unix()))
	return d
}

func SecondsUntilToday() time.Duration {
	now := time.Now()
	end := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	d, _ := time.ParseDuration(fmt.Sprintf("%ds", end.Unix()-now.Unix()))
	return d
}

func GetThisWeekStartTime() time.Time {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	return weekStart
}

func GetTimeBeforeNow(timeDifference time.Duration) time.Time {
	return time.Now().Add(-timeDifference)
}
