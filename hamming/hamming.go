package hamming

import (
	"errors"
	"unicode/utf8"
)

// Distance function to calculate the difference between two strings, or errors
func Distance(a, b string) (int, error) {
	var distance int
	type myError interface {
		Error() string
	}

	if utf8.RuneCountInString(a) != utf8.RuneCountInString(b) {
		defer func() {
			if r := recover(); r != nil {
				e := myError(errors.New("panic has occurred"))
				panic(e)
			}
		}()
		return distance, errors.New("panic in here")
	}
	ar, br := []rune(a), []rune(b)
	for i := range ar {
		if ar[i] != br[i] {
			distance++
		}
	}

	return distance, nil
}
