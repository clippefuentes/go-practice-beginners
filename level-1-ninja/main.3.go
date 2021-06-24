package main

import "fmt"

type clyne int

var clip clyne = 2
var y int

func main() {
	fmt.Println(clip)
	y = int(clip)
	fmt.Println(y)
}