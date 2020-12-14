package hamming

import (
	"errors"
	"unicode/utf8"
)
func Distance(a, b string) (int, error) {
	distance := 0
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
