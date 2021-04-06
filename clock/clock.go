package clock

import "fmt"

const (
	minutesPerHour = 60
	hoursPerDay    = 24
)

// Clock a struct holding hours and minute fields
type Clock struct {
	hour   int
	minute int
}

// New a function to create a clock type
func New(hour, minute int) Clock {
	return Clock{
		hour:   hour,
		minute: minute,
	}.format()
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hour, clock.minute)
}

// Add a method of type Clock to add minutes
func (clock Clock) Add(minutes int) Clock {
	clock.minute += minutes
	return clock.format()
}

// Substract  a method of type Clock to subtract minutes
func (clock Clock) Subtract(minutes int) Clock {
	clock.minute -= minutes
	return clock.format()
}

func (clock Clock) format() Clock {
	clock.hour += clock.minute / minutesPerHour
	if clock.minute %= minutesPerHour; clock.minute < 0 {
		clock.hour--
		clock.minute += minutesPerHour
	}
	if clock.hour %= hoursPerDay; clock.hour < 0 {
		clock.hour += hoursPerDay
	}
	return clock
}
