package main

import "fmt"

func main() {
	yearBd := 1997
	for {
		fmt.Println(yearBd)
		if (yearBd >= 2022) {
			break;
		}
		yearBd++
	}
}