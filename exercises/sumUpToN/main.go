package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Sum up to N
//
//  1. Get two numbers from the command-line: min and max
//  2. Convert them to integers (using Atoi)
//  3. By using a loop, sum the numbers between min and max
//
// RESTRICTIONS
//  1. Be sure to handle the errors. So, if a user doesn't
//     pass enough arguments or she passes non-numerics,
//     then warn the user and exit from the program.
//
//  2. Also, check that, min < max.
//
// HINT
//  For converting the numbers, you can use `strconv.Atoi`.
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 = 55
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
	for i := minInt; i <= maxInt; i++ {
		fmt.Printf("%d", i)
		if i != maxInt {
			fmt.Printf(" + ")
		}
		sum += i
	}
	fmt.Printf(" = %d", sum)
}