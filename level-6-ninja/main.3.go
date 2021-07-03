package main

import (
	"fmt"
)

func main () {
	defer foo()
	fmt.Println("TEST")
}

func foo() {
	defer func(){
		fmt.Println("??????")
	}()
	fmt.Println("THIS IS FOO")
}