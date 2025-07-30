package services

import (
	"strconv"
	"time"
)

func ParseStringToTime(date, layout string) (time.Time, error) {
	// Parse the date string using the provided layout
	parsedTime, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

func ParseStringToInt64(value string) (int64, error) {
	result, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func ParseStringToInt(value string) (int, error) {
	result, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return result, nil
}
