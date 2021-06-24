package main

import "fmt"

type person struct {
	first string
	last string
	favGames []string
}

func main() {
	p1 := person{
		first: "clyne",
		last: "fuentes",
		favGames: []string{"cyberpunk"},
	}
	p2 := person{
		first: "clippe",
		last: "cruz",
		favGames: []string{"rdr2"},
	}
	fmt.Println(p1)
	fmt.Println(p2)
}