package isogram

import "unicode"

//IsIsogram takes in a string, determines whether its an isogram
func IsIsogram(word string) bool {
	var seen = map[rune]bool{}
	for _, r := range word {
		r := unicode.ToUpper(r)
		if r == ' ' || r == '-' {
			continue
		}
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true

}
