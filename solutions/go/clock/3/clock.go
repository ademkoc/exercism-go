package clock

import (
	"fmt"
)

// Define the Clock type here.
type Clock int

const dayInMinutes = 1440

func normalize(minutes int) Clock {
	if minutes < 0 {
		return Clock((minutes % dayInMinutes) + dayInMinutes)
	} else if minutes >= dayInMinutes {
		return Clock(minutes % dayInMinutes)
	}
	return Clock(minutes)
}

func New(h, m int) Clock {
	return normalize(h*60 + m)
}

func (c Clock) Add(m int) Clock {
	return normalize(int(c) + m)
}

func (c Clock) Subtract(m int) Clock {
	return normalize(int(c) - m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}
