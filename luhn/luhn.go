/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 30-07-2018
 * |
 * | File Name:     luhn.go
 * +===============================================
 */

package luhn

import (
	"strconv"
	"strings"
)

// Valid validates given string with luhn algorithm
func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")
	if len(input) <= 1 {
		return false
	}

	number, err := strconv.Atoi(input)
	if err != nil {
		return false
	}

	count := 0
	sum := 0
	for number != 0 {
		i := number % 10
		if count%2 == 1 {
			i = i * 2
			if i > 9 {
				i -= 9
			}
			sum += i
		} else {
			sum += i
		}
		number /= 10
		count++
	}

	return sum%10 == 0
}
