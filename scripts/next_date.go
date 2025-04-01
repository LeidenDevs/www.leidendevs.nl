package main

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/dustin/go-humanize"
)

const (
	datetimePattern = `<time datetime="[0-9]{4}-[0-9]{2}-[0-9]{2}">[\s0-9A-Za-z]+</time>`
)

func getNextMeetupDate(now time.Time) time.Time {
	nextMeetup := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)

	// Loop until you find a Thursday
	for nextMeetup.Weekday() != time.Thursday {
		nextMeetup = nextMeetup.AddDate(0, 0, 1)
	}

	return nextMeetup.AddDate(0, 0, 14)
}

func updateHtmlWithNewDate(targetFile string, newDate time.Time) error {
	data, err := os.ReadFile(targetFile)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	newEventStr := fmt.Sprintf("<time datetime=\"%s\">%s</time>",
		fmt.Sprintf("%v-%02d-%v", newDate.Year(), int(newDate.Month()), newDate.Day()),
		fmt.Sprintf("%s %s %d", humanize.Ordinal(newDate.Day()), newDate.Month(), newDate.Year()),
	)

	re := regexp.MustCompile(datetimePattern)
	str := re.ReplaceAllString(string(data), newEventStr)

	err = os.WriteFile(targetFile, []byte(str), 0644)

	return err
}

// Print the date for the next meetup
// which will be the 3rd Thursday of the next month
func main() {
	targetFile := "../index.html"

	nextEventDate := getNextMeetupDate(time.Now())

	fmt.Println("Day of the next meetup is:", nextEventDate.Format(time.RFC822))

	err := updateHtmlWithNewDate(targetFile, nextEventDate)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully updated date in index.html")
}
