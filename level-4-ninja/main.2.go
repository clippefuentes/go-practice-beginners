package main

import "fmt"

func main() {
	x := []int{18,11,3,5,1,556,123,56,33}
	fmt.Println(x)
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", x)
}