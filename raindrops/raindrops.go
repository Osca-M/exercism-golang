package raindrops

import "strconv"

// Convert function takes in a int and checks its modulus against 3,5, and 7 and returns a string
func Convert(n int) string {
	var result string
	if n%3 == 0 {
		result += "Pling"
	}
	if n%5 == 0 {
		result += "Plang"
	}
	if n%7 == 0 {
		result += "Plong"
	}
	if result == "" {
		return strconv.Itoa(n)
	}
	return result
}
