package romannumerals

import "fmt"

// ToRomanNumeral takes in a int and returns a roman number equivalent or an error
func ToRomanNumeral(year int) (string, error) {
	if year > 3000 {
		return "", fmt.Errorf("{year} is greater than 3000")
	}
	if year < 1 {
		return "", fmt.Errorf("{year} is less than 1")
	}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	roman := ""
	i := 0
	for year > 0 {
		k := year / values[i]
		for j := 0; j < k; j++ {
			roman += symbols[i]
			year -= values[i]
		}
		i++
	}
	return roman, nil
}
