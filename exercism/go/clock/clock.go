package clock

import "fmt"

const testVersion = 4

type Clock int

func New(hour, minute int) Clock {
	minutesPerDay := 60 * 24
	minutes := hour*60 + minute
	minuteInDay := minutes % minutesPerDay
	c := Clock(minuteInDay)
	if c < 0 {
		c = New(0, int(c)+minutesPerDay)
	}
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%0.2d:%0.2d", c/60, c%60)
}

func (c Clock) Add(minutes int) Clock {
	return New(0, int(c)+minutes)
}
