// Package gigasecond provides second based measure time system functions
package gigasecond

import "time"

// AddGigasecond calculates when the given date would be a gigasecond old
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1_000_000_000)
}
