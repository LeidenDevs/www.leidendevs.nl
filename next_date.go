package main

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/dustin/go-humanize"
)

func getNextMeetupDate(now time.Time) time.Time {
	next := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)

	// Loop until you find a Thursday
	for next.Weekday() != time.Thursday {
		next = next.AddDate(0, 0, 1)
	}

	return next.AddDate(0, 0, 14)
}

// Print the date for the next meetup
// which will be the 3rd Thursday of the next month
func main() {
	fmt.Println("Day of the next meetup is:", getNextMeetupDate(time.Now()))

	// Update index.html with the date
	data, err := os.ReadFile("index.html")
	if err != nil {
		panic(err)
	}

	nextEventDate := getNextMeetupDate(time.Now())

	newEventStr := fmt.Sprintf("<time datetime=\"%s\">%s</time>",
		fmt.Sprintf("%v-%02d-%v", nextEventDate.Year(), int(nextEventDate.Month()), nextEventDate.Day()),
		fmt.Sprintf("%s %s %s", humanize.Ordinal(nextEventDate.Day()), nextEventDate.Month(), nextEventDate.Year()),
	)

	// fmt.Sprintf("%s %s", humanize.Ordinal(next.Day()), next.Month())
	re := regexp.MustCompile(`<time datetime="[0-9]{4}-[0-9]{2}-[0-9]{2}">[\s0-9A-Za-z]+</time>`)
	str := re.ReplaceAllString(string(data), newEventStr)

	fmt.Println("Writing to index.html")
	err = os.WriteFile("index.html", []byte(str), 0644)
	if err != nil {
		panic(err)
	}
}
