package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	val, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	meter := float64(val) / 3.281
	fmt.Printf("%d feet is %.2f meter", val, meter)
}