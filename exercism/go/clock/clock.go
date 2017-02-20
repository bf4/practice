package clock

import "fmt"

const testVersion = 4

type Clock struct {
	Hour, Minute int
}

func New(hour, minute int) Clock {
	adjustMinute := func(m int) int {
		for {
			if m < 0 {
				hour -= 1
				m = m + 60
			} else if m < 60 {
				return m
			} else {
				hour += 1
				m = m - 60
			}
		}
	}
	minute = adjustMinute(minute)
	adjustHour := func(h int) int {
		for {
			if h < 0 {
				h = h + 24
			} else if h < 24 {
				return h
			} else {
				h = h - 24
			}
		}
	}
	hour = adjustHour(hour)
	return Clock{hour, minute}
}

func (c Clock) String() string {
	return fmt.Sprintf("%0.2d:%0.2d", c.Hour, c.Minute)
}

func (c Clock) Add(minutes int) Clock {
	return New(c.Hour, c.Minute+minutes)
}
