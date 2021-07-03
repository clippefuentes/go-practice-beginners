package main

import "fmt"

func main() {
	fmt.Println("Test")

	ii := []int{1,2,3,4,5}
	fmt.Println(foo(ii...))
	fmt.Println(bar(ii))
}

func foo (xi ...int) int {
	total := 0
	for _, n := range xi {
		total += n
	}
	return total
}

func bar (xi []int) int {
	total := 0
	for _, n := range xi {
		total += n
	}
	return total
}