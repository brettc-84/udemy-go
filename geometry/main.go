package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	a float64
	b float64
	c float64
}
type square struct {
	a float64
}

func main() {
	sq := square{
		a: 4,
	}
	printArea(sq)

	tr := triangle{
		a: 3,
		b: 3,
		c: 3,
	}
	printArea(tr)
}

func printArea(s shape) {
	fmt.Println("Area is:", s.getArea())
}

func (t triangle) getArea() float64 {
	// heron's algorithm
	s := (t.a + t.b + t.c) / 2
	return math.Sqrt(s * (s - t.a) * (s - t.b) * (s - t.c))
}

func (s square) getArea() float64 {
	return s.a * s.a
}
