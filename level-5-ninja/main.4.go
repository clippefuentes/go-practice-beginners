package main

import "fmt"

func main() {
	x :=  struct {
		fName string
		cars map[string]string
	}{
		fName: "clyne",
		cars: map[string]string{
			"black": "subaro",
		},
	}
	fmt.Println(x)
}