/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 25-07-2018
 * |
 * | File Name:     raindrop.go
 * +===============================================
 */

// Package raindrops provides function for speak the raindrop easily :joy:
package raindrops

import (
	"strconv"
	"strings"
)

type word struct {
	key   int
	value string
}

var lang = []word{
	word{3, "Pling"},
	word{5, "Plang"},
	word{7, "Plong"},
}

// Convert creates string of given number in raindrop-speak
func Convert(n int) string {
	b := &strings.Builder{}

	for _, w := range lang {
		if n%w.key == 0 {
			b.WriteString(w.value)
		}
	}

	if b.Len() == 0 {
		return strconv.Itoa(n)
	}

	return b.String()
}
