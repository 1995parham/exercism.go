/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 26-07-2018
 * |
 * | File Name:     listops.go
 * +===============================================
 */

// Package listops provides functional language like functions for go slics
package listops

// IntList is a list of integers
type IntList []int

// Length == len(l)
func (l IntList) Length() int {
	return len(l)
}

// Reverse returns new list based on l in backward order
func (l IntList) Reverse() IntList {
	rl := make(IntList, len(l))

	for i := range l {
		rl[len(l)-1-i] = l[i]
	}

	return rl
}
