package clock

import (
	"fmt"
)

// Define the Clock type here.
type Clock struct {
	hour int
	min  int
}

func New(h, m int) Clock {
	if m < 0 {
		for m < 0 {
			m += 60
			h -= 1
		}

		m = m % 60
	}

	if m >= 60 {
		for m >= 60 {
			m -= 60
			h += 1
		}

		m = m % 60
	}

	if h >= 24 || h < 0 {
		h = h % 24

		for h < 0 {
			h += 24
		}
	}

	return Clock{
		hour: h,
		min:  m,
	}
}

func (c Clock) Add(m int) Clock {
	totalMin := c.min + m
	minutes := totalMin % 60
	totalHour := ((totalMin - minutes) / 60) + c.hour
	hours := totalHour % 24

	c.hour = hours
	c.min = minutes
	return c
}

func (c Clock) Subtract(m int) Clock {
	rawMin := c.min - m
	if rawMin >= 0 {
		c.min = rawMin
		return c
	}

	c.min = rawMin % 60
	hour := (((rawMin * -1) + (rawMin % 60)) / 60) + 1
	c.hour = (c.hour - hour) % 24

	if c.hour < 0 {
		c.hour += 24
	}

	if c.min < 0 {
		c.min += 60
	}

	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.min)
}
