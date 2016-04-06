package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	var xx int
	xx = 0
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
		xx = xx + 1
	}
	fmt.Print(xx)
	fmt.Println(len(os.Args))
	fmt.Println(s)
}
