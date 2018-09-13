/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 24-07-2018
 * |
 * | File Name:     sapce_age.go
 * +===============================================
 */

package space

// Planet represent space planet :joy:
type Planet string

var orbital = map[Planet]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age gets an age in seconds, calculate how old someone would be on other planet
func Age(age float64, planet Planet) float64 {
	earth := age / 31557600

	return earth / orbital[planet]
}
