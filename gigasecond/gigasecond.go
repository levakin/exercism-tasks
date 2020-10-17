// Package gigasecond provides method to operate with time
package gigasecond

import "time"

// Gigasecond is 10^9 (1,000,000,000) seconds.
const Gigasecond = time.Second * 1000000000

// AddGigasecond given a moment, determines the moment that would be after a gigasecond has passed and returns its value.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(Gigasecond)
}
