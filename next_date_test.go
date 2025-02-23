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
		expected string
	}{
		{time.Date(2025, 2, 23, 0, 0, 0, 0, time.UTC), "20th March"},
		{time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC), "16th January"},
	}

	for _, scenario := range scenarios {
		assert.Equal(t, scenario.expected, getNextMeetupDate(scenario.date))
	}
}
