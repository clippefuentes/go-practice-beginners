package main

import (
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Case Insensitive Search
//
//  Allow for case-insensitive searching
//
// EXAMPLE
//  Let's say that the user runs the program like this:
//    go run main.go LAZY
//
//  Or like this: go run main.go lAzY
//  Or like this: go run main.go lazy
//
//  For all cases above, the program should find
//  the "lazy" keyword.
// ---------------------------------------------------------
const corpus = "lazy cat jumps again and again and again"

func main() {
	query := os.Args[1:]
	words := strings.Fields(corpus)

queries:
	for _, q := range query {
		qWord := strings.ToLower(q)
		search:
		for i, w := range words {
			switch q {
			case "and", "or", "the":
				break searchw
			}
			if w == qWord {
				fmt.Printf("#%d: %s \n", i+1, w)
				continue queries
			}
		}
	}
}