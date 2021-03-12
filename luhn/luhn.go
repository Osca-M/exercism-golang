package luhn

import (
	"strconv"
	"strings"
)

// Valid used to validate a string based on the Luhn formula
func Valid(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	if len(s) <= 1 {
		return false
	}
	var result int
	step := len(s)%2 != 0

	for _, r := range s {
		v, err := strconv.Atoi(string(r))
		if err != nil {
			return false
		}
		step = !step
		if step {
			v *= 2
			if v > 9 {
				v -= 9
			}
		}
		result += v
	}
	return result%10 == 0
}
