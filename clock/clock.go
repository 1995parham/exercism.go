/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 04-08-2018
 * |
 * | File Name:     clock.go
 * +===============================================
 */

package clock

import "fmt"

// Clock represents clock with hour and minute
type Clock struct {
	minute int
	hour   int
}

// New creates new clock
func New(hour, minute int) Clock {
	m := hour*60 + minute
	m %= (24 * 60)
	if m < 0 {
		m += 24 * 60
	}

	return Clock{
		minute: m % 60,
		hour:   m / 60,
	}
}

// String returns string representation of clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add moves clock m minutes forward
func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}

// Subtract moves clock m minutes backward
func (c Clock) Subtract(minutes int) Clock {
	return New(c.hour, c.minute-minutes)
}
