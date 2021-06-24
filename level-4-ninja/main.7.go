package main

import "fmt"

func main() {
	x := [][]string {
		{"Clyne" , "Paulo"},
		{"Clyne2" , "Paulo2"},
	}
	
	for i, xs := range x {
		fmt.Println("index", i)
		for j, val := range xs {
			fmt.Printf("index: %v \t value: %v \t \n", j, val)
		}
	}
}