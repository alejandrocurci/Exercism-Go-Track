package clock

import "fmt"

type Clock struct {
	minutes int
	hours   int
}

// New is a constructor for Clock type
// negative hours and minutes roll over continuously
func New(hour, minute int) Clock {
	for minute < 0 {
		minute += 60
		hour -= 1
	}
	for hour < 0 {
		hour += 24
	}
	m := minute % 60
	h := (hour + (minute-m)/60) % 24
	return Clock{
		minutes: m,
		hours:   h,
	}
}

// String returns Clock type as "00:00" format
func (c Clock) String() string {
	minStr := fmt.Sprintf("%d", c.minutes)
	hourStr := fmt.Sprintf("%d", c.hours)
	if c.minutes < 10 {
		minStr = fmt.Sprintf("0%d", c.minutes)
	}
	if c.hours < 10 {
		hourStr = fmt.Sprintf("0%d", c.hours)
	}
	return hourStr + ":" + minStr
}

// Add adds minutes to a clock
func (c Clock) Add(minutes int) Clock {
	return New(c.hours, c.minutes+minutes)
}

// Subtract subtracts minutes to a clock
func (c Clock) Subtract(minutes int) Clock {
	return New(c.hours, c.minutes-minutes)
}
