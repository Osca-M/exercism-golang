package summultiples

// SumMultiples calculates sum of numbers within between 0 and limit that are divisbile by a number in divisors
func SumMultiples(limit int, divisors ...int) int {
	var sm int = 0
	for n := 1; n < limit; n++ {
		for _, s := range divisors {
			if s == 0 {
				continue
			}
			r := n % s
			if r == 0 {
				sm = sm + n
				break
			}
		}
	}
	return sm
}
