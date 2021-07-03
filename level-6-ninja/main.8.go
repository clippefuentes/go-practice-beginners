package main

import (
	"fmt"
)

func foo () func () int {
	return func () int {
		return 420
	}
}

var y = foo()

func main () {
	t := foo()
	fmt.Println(t())
	fmt.Println("---")
	fmt.Println(y())
	fmt.Println("TEST")
}
