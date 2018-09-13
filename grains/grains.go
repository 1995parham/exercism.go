/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 02-08-2018
 * |
 * | File Name:     grains.go
 * +===============================================
 */

// Package grains calculate the number of grains of wheat on a chessboard given that the number on each square doubles.
package grains

import "fmt"

// Square returns number of grains on each square
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, fmt.Errorf("n must be in [1, 64]")
	}
	var v uint64 = 1
	for i := 0; i < n-1; i++ {
		v *= 2
	}

	return v, nil
}

// Total returns total number of grains on the chessboard
func Total() uint64 {
	var v uint64 = 1
	var s uint64

	for i := 0; i < 64; i++ {
		s += v
		v *= 2
	}

	return s
}
