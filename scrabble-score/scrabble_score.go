package scrabble

import (
	"fmt"
	"strings"
)

func Score(str string) int {
	var score int
	var otherMap = map[string] int {}
	letters := [7][10]string{
		{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
		{"D", "G"},
		{"B", "C", "M", "P"},
		{"F", "H", "V", "W", "Y"},
		{"K"},
		{"J", "X"},
		{"Q", "Z", ""},
	}
	for i, v := range letters {
		if i == 0 {
			for _, l := range v {
				if l != ""{
					otherMap[l] = 1
				}
			}
		} else if i == 1 {
			for _, l := range v {
				if l != ""{otherMap[l] = 2}
			}
		}  else if i == 2 {
			for _, l := range v {
				if l != ""{otherMap[l] = 3}
			}
		}  else if i == 3 {
			for _, l := range v {
				if l != ""{
					otherMap[l] = 4
				}
			}
		}  else if i == 4 {
			for _, l := range v {
				if l != "" {
					otherMap[l] = 5
				}
			}
		}  else if i == 5 {
			for _, l := range v {
				if l != ""{
					otherMap[l] = 8
				}
			}
		}  else if i == 6 {
			for _, l := range v {
				fmt.Println(l, "l")
				if l != ""{
					otherMap[l] = 10
				}
			}
		}
	}
	upperCaseString := strings.ToUpper(str)
	//charValue := map[string]int{
	//	"A": 1,
	//	"E": 1,
	//	"I": 1,
	//	"O": 1,
	//	"U": 1,
	//	"L": 1,
	//	"N": 1,
	//	"R": 1,
	//	"S": 1,
	//	"T": 1,
	//	"D": 2,
	//	"G": 2,
	//	"B": 3,
	//	"C": 3,
	//	"M": 3,
	//	"P": 3,
	//	"F": 4,
	//	"H": 4,
	//	"V": 4,
	//	"W": 4,
	//	"Y": 4,
	//	"K": 5,
	//	"J": 8,
	//	"X": 8,
	//	"Q": 10,
	//	"Z": 10,
	//}
	for _, v := range upperCaseString{
		value, ok := charValue[string(v)]
		if ok == true {
			score += value
		}
	}
	return score
}