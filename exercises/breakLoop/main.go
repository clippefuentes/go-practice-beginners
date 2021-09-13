package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Break Up
//
//  1. Extend the "Only Evens" exercise
//  2. This time, use an infinite loop.
//  3. Break the loop when it reaches to the `max`.
//
// RESTRICTIONS
//  You should use the `break` statement once.
//
// HINT
//  Do not forget incrementing the `i` before the `continue`
//  statement and at the end of the loop.
//
// EXPECTED OUTPUT
//    go run main.go 1 10
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Enter min and max")
		return
	}
	min, max := os.Args[1], os.Args[2]
	minInt, err1 := strconv.Atoi(min)
	maxInt, err2 := strconv.Atoi(max)

	if err1 != nil || err2 != nil {
		fmt.Println("Both input must be a number")
		return
	}

	if minInt > maxInt {
		fmt.Println("enter smaller number first than the second one")
		return
	}
	sum := 0
	i := minInt
	for {
		if i % 2 != 0 {
			i++
			continue
		}
		
		if i == maxInt {
			fmt.Printf("%d ", i)
			sum += i
			i++
			break;
		}
		fmt.Printf("%d + ", i)
		sum += i
		i++
	}
	fmt.Printf("= %d", sum)
}