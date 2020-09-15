package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return 2 * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	s := square{
		side: 42.7,
	}

	c := circle{
		radius: 89.31,
	}

	info(s)
	info(c)
}
