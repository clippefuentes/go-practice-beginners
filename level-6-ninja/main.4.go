package main

import (
	"fmt"
)

type person struct  {
	first string
	last string
	age int
}

func (p person) speak() {
	fmt.Printf("My name is %v, %v and i am %v years old", p.first, p.last, p.age)
}

func main () {
	p1 := person{
		first: "Clyne",
		last: "Fuentes",
		age: 24,
	}
	p1.speak()
}
