package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Leap Year
//
//  Find out whether a given year is a leap year or not.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a year number
//
//  go run main.go eighties
//    "eighties" is not a valid year.
//
//  go run main.go 2018
//    2018 is not a leap year.
//
//  go run main.go 2100
//    2100 is not a leap year.
//
//  go run main.go 2019
//    2019 is not a leap year.
//
//  go run main.go 2020
//    2020 is a leap year.
//
//  go run main.go 2024
//    2024 is a leap year.
// ---------------------------------------------------------

func main() {
	// If the year is evenly divisible by 4, go to step 2. ...
	// If the year is evenly divisible by 100, go to step 3. ...
	// If the year is evenly divisible by 400, go to step 4. .
	if len(os.Args) != 2 {
		fmt.Println("Give me a year number")
		return
	}

	v, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q is not a valid year", os.Args[1])
		return
	}

	if v%400 == 0 {
		fmt.Printf("%d is a leap year", v)
	} else if v%100 == 0 {
		fmt.Printf("%d is not a leap year", v)
	} else if  v%4  == 0 {
		fmt.Printf("%d is a leap year", v)
	} else  {
		fmt.Printf("%d is not a leap year", v)
	}
}