/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 26-07-2018
 * |
 * | File Name:     accumulate.go
 * +===============================================
 */

package accumulate

// Accumulate performs given operation on each element of given list and return new list
func Accumulate(list []string, operation func(string) string) []string {
	results := make([]string, len(list))

	for i, v := range list {
		results[i] = operation(v)
	}

	return results
}
