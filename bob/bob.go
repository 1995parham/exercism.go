// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	q := false
	y := false

	sremark := strings.Fields(remark)
	if len(sremark) == 0 {
		return "Fine. Be that way!"
	}

	if strings.HasSuffix(sremark[len(sremark)-1], "?") {
		q = true
	}
	if strings.Compare(remark, strings.ToUpper(remark)) == 0 && strings.Compare(remark, strings.ToLower(remark)) != 0 {
		y = true
	}

	if y && q {
		return "Calm down, I know what I'm doing!"
	}
	if y {
		return "Whoa, chill out!"
	}
	if q {
		return "Sure."
	}

	return "Whatever."
}
