package hamming

import (
	"errors"
	"fmt"
)

// Distance returns difference between strings or an error in the case we have one
func Distance(a, b string) (int, error) {
	distance := 0
	type myError interface {
		Error() string
	}
	fmt.Println(a, "a")
	fmt.Println(b, "b")

	if len(a) != len(b) {
		defer func() {
			if r := recover(); r != nil {
				e := myError(errors.New("panic has occurred"))
				panic(e)
			}
		}()
		return distance, errors.New("panic in here")
	}
	for i := range a {
		if string([]rune(a)[i]) != string([]rune(b)[i]) {
			distance++
		}
	}

	return distance, nil
}
