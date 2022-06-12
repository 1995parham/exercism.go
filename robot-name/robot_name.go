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
	"fmt"
	"math/rand"
	"time"
)

// not safe for concurrent use.
var used map[string]bool

func init() {
	rand.Seed(time.Now().UnixNano())

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
		for used[r.name] {
			digits := rand.Int63n(1000)    // [0, 1000)
			alpha1 := rand.Int63n(26) + 65 // [A, Z]
			alpha2 := rand.Int63n(26) + 65 // [A, Z]

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
