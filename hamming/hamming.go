// Package hamming contains distance function for calculate hamming distance
package hamming

import "fmt"

// Distance calculate hamming distance between two given DNA
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, fmt.Errorf("String length must be equal")
	}
	c := 0
	for i := range a {
		if a[i] != b[i] {
			c++
		}
	}
	return c, nil
}
