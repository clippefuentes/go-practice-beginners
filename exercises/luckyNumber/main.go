package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) != 2 {
		fmt.Println("Please Input Number")
		return
	}

	v, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Enter a number")
		return
	}

	if v <= 0 {
		fmt.Println("Enter positive number")
		return
	}


	for turn := 1; turn <=5; turn++ {
		guess := rand.Intn(v+1)

		if guess == v {
			if turn == 1 {
				fmt.Println("You win at first time")
			} else {
				fmt.Println("You win")
			}
			return
		}
	}
	fmt.Println("try again")
}