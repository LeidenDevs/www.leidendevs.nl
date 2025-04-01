// Write a test for the getNextMeetupDate function
package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetNextMeetupDate(t *testing.T) {
	scenarios := []struct {
		date     time.Time
		expected time.Time
	}{
		{time.Date(2025, 2, 23, 0, 0, 0, 0, time.UTC), time.Date(2025, time.February, 20, 0, 0, 0, 0, time.UTC)},
		{time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, time.December, 19, 0, 0, 0, 0, time.UTC)},
	}

	for _, scenario := range scenarios {
		assert.Equal(t, scenario.expected, getNextMeetupDate(scenario.date))
	}
}
