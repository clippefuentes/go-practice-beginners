package main

import (
	"fmt"
)

var x = func () {
	fmt.Println("THIS IS X")
}

func main () {
	f := func () {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}
	x()
	f()
	fmt.Println("TEST")
}
