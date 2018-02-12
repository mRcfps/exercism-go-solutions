// Package gigasecond do calculations about gigaseconds.
package gigasecond

import "time"

// AddGigasecond adds a gigasecond to the given time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}
