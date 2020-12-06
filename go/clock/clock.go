package clock

import "fmt"

type Clock struct {
	hours   int
	minutes int
}

func New(hours, minutes int) Clock {
	clock := Clock{
		hours:   hours,
		minutes: minutes,
	}

	return clock.Calculate()
}

func (clock Clock) Add(minutes int) Clock {
	clock.minutes += minutes
	return clock.Calculate()
}

func (clock Clock) Subtract(minutes int) Clock {
	clock.minutes -= minutes
	return clock.Calculate()
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hours, clock.minutes)
}

// Probably a better way to do this using modulus but its 0230am and I'm off to bed
func (clock Clock) Calculate() Clock {
	for clock.minutes >= 60 {
		clock.hours += 1
		clock.minutes -= 60
	}
	for clock.minutes < 0 {
		clock.hours -= 1
		clock.minutes += 60
	}

	for clock.hours >= 24 {
		clock.hours -= 24
	}
	for clock.hours < 0 {
		clock.hours += 24
	}

	return clock
}
