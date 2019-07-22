package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func main() {
	sq := square{
		sideLength: 4,
	}
	printArea(sq)

	tr := triangle{
		height: 5,
		base:   5.6,
	}
	printArea(tr)
}

func printArea(s shape) {
	fmt.Println("Area is:", s.getArea())
}

func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
