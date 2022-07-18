package utils

import (
	"log"
	"time"
)

func ParseTime(layout string, date string) time.Time {
	dateTime, err := time.Parse(layout, date)
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}
	return dateTime
}
