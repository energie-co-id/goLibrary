package qryFilter

import (
	"strconv"
	"time"
)

func DateSplitFilter(date string, dateType string, from string, to string) (finalStartTime time.Time, finalStopTime time.Time, err error) {
	customTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		customTime = time.Now()
	}
	var startTime time.Time
	var stopTime time.Time

	switch dateType {
	case "daily":
		startTime = time.Date(customTime.Year(), customTime.Month(), customTime.Day(), -7, 0, 0, 0, customTime.Location())
		stopTime = time.Date(customTime.Year(), customTime.Month(), customTime.Day(), 16, 59, 59, 99, customTime.Location())
	case "weekly":
		startTime = time.Date(customTime.Year(), customTime.Month(), customTime.Day()-7, 0, 0, 0, 0, customTime.Location())
		stopTime = time.Date(customTime.Year(), customTime.Month(), customTime.Day(), 16, 59, 59, 99, customTime.Location())
	case "monthly":
		startTime = time.Date(customTime.Year(), customTime.Month(), 1, 0, 0, 0, 0, customTime.Location())
		stopTime = time.Date(customTime.Year(), startTime.Month()+1, startTime.Day()-1, 23, 59, 59, 99, customTime.Location())
	case "yearly":
		year, err := strconv.ParseInt(date, 10, 64)
		if err != nil {
			return startTime, stopTime, err
		}
		startTime = time.Date(int(year), 0, 1, 0, 0, 0, 0, customTime.Location())
		stopTime = time.Date(int(year), 13, startTime.Day()-1, 23, 59, 59, 99, customTime.Location())
	default:
		from, err := time.Parse("2006-01-02", from)
		if err != nil {
			return startTime, stopTime, err
		}
		to, error := time.Parse("2006-01-02", to)
		if error != nil {
			return startTime, stopTime, err
		}
		startTime = time.Date(from.Year(), from.Month(), from.Day(), -7, 0, 0, 0, from.Location())
		stopTime = time.Date(to.Year(), to.Month(), to.Day(), 16, 59, 59, 99, to.Location())
	}

	return startTime, stopTime, nil
}
