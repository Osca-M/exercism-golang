package main

import (
	"./hamming"
	"fmt"
)

func main()  {
	fmt.Println(hamming.Distance("GGACGGATTCTG", "AGGACGGATTCT"))
}
//GGACGGATTCTG
//AGGACGGATTCT