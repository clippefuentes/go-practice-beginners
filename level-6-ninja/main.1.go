package main

import "fmt"

func main() {
	fmt.Println("TEST")
	fmt.Println(test1())
	fmt.Println(test2())
}

func test1() int {
	return 420
}

func test2() (int, string) {
	return 1, "TETS 2"
}