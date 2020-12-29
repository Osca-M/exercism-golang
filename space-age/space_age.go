package space

import (
	"math"
)

// Planet name of any of out 8 planets
type Planet string

// Age a function that takes in seconds and a planet name then from that calculates the age in years of the seconds
// relative to time on earth
func Age(s float64, p Planet) float64 {
	x := s / 31557600
	switch p {
	case "Earth":
		return math.Round(x*100) / 100
	case "Mercury":
		return math.Round(x/0.2408467*100) / 100
	case "Venus":
		return math.Round(x/0.61519726*100) / 100
	case "Mars":
		return math.Round(x/1.8808158*100) / 100
	case "Jupiter":
		return math.Round(x/11.862615*100) / 100
	case "Saturn":
		return math.Round(x/29.447498*100) / 100
	case "Uranus":
		return math.Round(x/84.016846*100) / 100
	case "Neptune":
		return math.Round(x/164.79132*100) / 100
	}
	return 0
}
