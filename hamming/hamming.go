package hamming

import (
	"fmt"
)

// Distance function to calculate the difference between two strings, or errors
func Distance(a, b string) (int, error) {
	ar, br := []rune(a), []rune(b)
	if len(ar) != len(br) {
		return 0, fmt.Errorf("the length of %s is not equal to the length of %s", a, b)
	}
	var distance int
	for i := range ar {
		if ar[i] != br[i] {
			distance++
		}
	}
	return distance, nil
}
