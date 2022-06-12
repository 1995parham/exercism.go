package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	acronym := ""

	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && (c != '\'')
	}

	for _, word := range strings.FieldsFunc(s, f) {
		acronym += strings.ToUpper(word[0:1])
	}

	return acronym
}
