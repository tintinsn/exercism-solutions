// Package gigasecond provides a function to calculate the moment
// when someone has lived for 1 gigasecond.
package gigasecond

import (
	"time"
)

// AddGigasecond takes a time and returns the time
func AddGigasecond(t time.Time) time.Time {
	duration := 1000000000 * time.Second
	t = t.Add(duration)
	return t
}
