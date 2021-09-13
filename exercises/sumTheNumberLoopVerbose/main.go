package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Sum the Numbers
//
//  1. By using a loop, sum the numbers between 1 and 10.
//  2. Print the sum.
//
// EXPECTED OUTPUT
//  Sum: 55
// ---------------------------------------------------------

func main() {
	sum := 0
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d ", i)
		if i != 10 {
			fmt.Printf("+")
		}
		sum += i
	}
	fmt.Printf("= %d", sum)
}