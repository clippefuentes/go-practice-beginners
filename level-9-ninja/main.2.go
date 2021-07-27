package main

import (
	"fmt"
)

type person struct {
	first string
}

type human interface {
	speak()
}

func (p *person) speak () {
	fmt.Println("Hello Im", p.first)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person {
		first: "Clyne",
	}
	saySomething(&p1)
}