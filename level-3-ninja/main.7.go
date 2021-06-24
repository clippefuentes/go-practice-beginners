package main

import "fmt"

func main() {
	x := "clyne"

	if x == "clyne" {
		fmt.Println("clyne is awesome")
	} else if x == "paulo" {
		fmt.Println("stop calling me paulo")
	} else {
		fmt.Println("Who u")
	}
}