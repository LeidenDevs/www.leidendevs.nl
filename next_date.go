package main

import (
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
)

func getNextMeetupDate(now time.Time) string {
	next := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, time.UTC)

	// Loop until you find a Thursday
	for next.Weekday() != time.Thursday {
		next = next.AddDate(0, 0, 1)
	}

	next = next.AddDate(0, 0, 14)
	return fmt.Sprintf("%s %s", humanize.Ordinal(next.Day()), next.Month())
}

// Print the date for the next meetup
// which will be the 3rd Thursday of the next month
func main() {
	fmt.Println("Day of the next meetup is:", getNextMeetupDate(time.Now()))
}
