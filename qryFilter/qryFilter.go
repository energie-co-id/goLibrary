package qryFilter

import (
	"time"
)

func DateSplitFilter(date string, dateType string, from string, to string) (finalStartTime string, finalStopTime string) {
	customTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		customTime = time.Now()
	}
	var startTime time.Time
	var stopTime time.Time

	switch dateType {
	case "daily":
		startTime = time.Date(customTime.Year(), customTime.Month(), customTime.Day(), 0, 0, 0, 0, customTime.Location())
		stopTime = time.Date(customTime.Year(), customTime.Month(), customTime.Day(), 23, 59, 59, 99, customTime.Location())
	case "weekly":
		startTime = time.Date(customTime.Year(), customTime.Month(), customTime.Day()-6, 0, 0, 0, 0, customTime.Location())
		stopTime = time.Date(customTime.Year(), customTime.Month(), customTime.Day(), 23, 59, 59, 99, customTime.Location())
	case "monthly":
		startTime = time.Date(customTime.Year(), customTime.Month(), 1, 0, 0, 0, 0, customTime.Location())
		stopTime = time.Date(customTime.Year(), startTime.Month()+1, startTime.Day()-1, 23, 59, 59, 99, customTime.Location())
	default:
		from, err := time.Parse("2006-01-02", from)
		if err != nil {
			return "error", "error"
		}
		to, error := time.Parse("2006-01-02", to)
		if error != nil {
			return "error", "error"
		}
		startTime = time.Date(from.Year(), from.Month(), from.Day(), 0, 0, 0, 0, from.Location())
		stopTime = time.Date(to.Year(), to.Month(), to.Day(), 23, 59, 59, 99, to.Location())
	}

	return startTime.Format(time.RFC3339), stopTime.Format(time.RFC3339)
}
