/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 27-08-2018
 * |
 * | File Name:     twelve_days.go
 * +===============================================
 */

package twelve

import "fmt"

var days = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

var presents = map[int]string{
	1:  "a Partridge in a Pear Tree",
	2:  "two Turtle Doves",
	3:  "three French Hens",
	4:  "four Calling Birds",
	5:  "five Gold Rings",
	6:  "six Geese-a-Laying",
	7:  "seven Swans-a-Swimming",
	8:  "eight Maids-a-Milking",
	9:  "nine Ladies Dancing",
	10: "ten Lords-a-Leaping",
	11: "eleven Pipers Piping",
	12: "twelve Drummers Drumming",
}

// Song generates the lyrics to 'The Twelve Days of Christmas'.
func Song() string {
	var s string

	for i := 1; i < 12; i++ {
		s += fmt.Sprintf("%s\n", Verse(i))
	}

	s += Verse(12)

	return s
}

// Verse generate one day of twelve days lyric!
func Verse(day int) string {
	present := ""

	for i := day; i >= 1; i-- {
		switch i {
		case 1:
			present += presents[i]
		case 2:
			present += fmt.Sprintf("%s, and ", presents[i])
		default:
			present += fmt.Sprintf("%s, ", presents[i])
		}
	}

	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s.", days[day], present)
}
