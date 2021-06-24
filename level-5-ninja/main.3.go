package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck {
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
		fourWheel: true,
	}
	x:= sedan {
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		luxury: true,
	}
	fmt.Println(t)
	fmt.Println(x)
}