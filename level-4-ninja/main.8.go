package main

import "fmt"

func main() {
	x := map[string][]string{
		"1": {"clyne", "paulo"},
		"2": {"laurent", "tsai"},
	}
	fmt.Println(x)

	for  k, v := range x {
		fmt.Println(k, v)
	}

	x["3"] = []string{"the", "neighboorhod"}

	for  k, v := range x {
		fmt.Println(k, v)
	}
}