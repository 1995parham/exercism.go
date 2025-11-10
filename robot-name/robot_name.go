/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 21-08-2018
 * |
 * | File Name:     robot_name.go
 * +===============================================
 */

package robotname

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

// not safe for concurrent use.
var used map[string]bool

func init() {
	used = make(map[string]bool)
	used[""] = true
}

// Robot represents robot.
type Robot struct {
	set  bool
	name string
}

// Name creates new name for your robot
func (r *Robot) Name() (string, error) {
	if !r.set {
		// Check if namespace is exhausted (26*26*10*10*10 = 676,000 possible names)
		// Subtract 1 to account for the empty string in the used map
		maxNames := 26*26*10*10*10 + 1
		if len(used) >= maxNames {
			return "", errors.New("namespace exhausted")
		}

		for used[r.name] {
			digits := rand.IntN(1000)   // [0, 1000)
			alpha1 := rand.IntN(26) + 65 // [A, Z]
			alpha2 := rand.IntN(26) + 65 // [A, Z]

			r.name = fmt.Sprintf("%c%c%03d", alpha1, alpha2, digits)
		}
		used[r.name] = true
		r.set = true
	}

	return r.name, nil
}

// Reset resets robot into its factory settings
func (r *Robot) Reset() {
	r.set = false
}
