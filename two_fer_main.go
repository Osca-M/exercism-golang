package main

import (
	"./two-fer"
	"fmt"
)

func main() {
	var name string
	fmt.Println("Enter your Name: ")
	fmt.Scanln(&name)
	fmt.Println(twofer.ShareWith(name))
}