package isogram

import "unicode"

// IsIsogram detects given word is isogram or not
func IsIsogram(w string) bool {
	detection := make(map[rune]bool)

	for _, r := range w {
		if !unicode.IsNumber(r) && !unicode.IsLetter(r) {
			continue
		}
		r = unicode.ToUpper(r)
		if detection[r] {
			return false
		}
		detection[r] = true
	}

	return true
}
