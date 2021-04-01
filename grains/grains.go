// Package grains is a program to calculate the determine the number of grains in chessboard guided by this task
//http://www.javaranch.com/grains.jsp
package grains

import (
	"errors"
)

// Square to get the value of 2 to the power of numbers between 1 and 64
func Square(s int) (uint64, error) {
	if s < 1 || s > 64 {
		return 0, errors.New("error")
	}
	return 1 << (s - 1), nil
}

//Total uses Geometric series formula to get the sum of 2 to the power numbers between 0 and 64
func Total() uint64 {
	return 1<<64 - 1
}
