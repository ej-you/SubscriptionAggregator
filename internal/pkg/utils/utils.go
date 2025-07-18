package utils

import (
	"fmt"
	"time"
)

const _dateFmt = "01-2006" // format to parse date into string to time.Time

// ParseDate parses date from given string by the const template.
func ParseDate(dateStr string) (time.Time, error) {
	date, err := time.Parse(_dateFmt, dateStr)
	if err != nil {
		return date, fmt.Errorf("parse date: %w", err)
	}
	return date, nil
}
