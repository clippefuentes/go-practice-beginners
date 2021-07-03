package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius int
}

type square struct {
	sides int
}

func (c circle) area() float64 {
	return math.Pi * (math.Pow(float64(c.radius), 2))
}

func (c square) area() float64 {
	return float64(c.sides) * float64(c.sides)
}

func main () {
	c1 := circle{
		radius: 3,
	}
	s1 := square{
		sides: 3,
	}
	fmt.Printf("Circle C of 3 radius has the area of %v", c1.area())
	fmt.Printf("Squarew S of 3 sides has the area of %v", s1.area())
}
