// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Kind is a triangle type
// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// NaT not a triangle
	NaT Kind = iota
	// Equ equilateral
	Equ Kind = iota
	// Iso isosceles
	Iso Kind = iota
	// Sca scalene
	Sca Kind = iota
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}

	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	}

	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return NaT
	}

	if a > b+c || b > a+c || c > a+b {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || b == c || c == a {
		return Iso
	}

	return Sca
}
