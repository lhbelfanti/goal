package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

type shape interface {
	area() float64
}

func main () {
	ci := circle{
		radius: 3.80,
	}

	sq := square{
		length: 12,
	}

	info(ci)
	info(sq)
}

func info(sh shape) {
	fmt.Println(sh.area())
}
