package main

import "fmt"

func main() {
	x := "clyne"

	switch x {
	case "clyne":
		fmt.Println("CLYNE")
	case "pauli":
		fmt.Println("Pauli? r u serious")
	default:
		fmt.Println("Hiiiiii")
	}
}