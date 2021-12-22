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

type shape interface {
	area() float64
}

func info(s shape) {
	switch s.(type) {
	case square:
		fmt.Println("square area", s.(square).length)
	case circle:
		fmt.Println("circle area", s.(circle).radius)
	default:
		fmt.Println("default value")
	}

	fmt.Println(s.area())

}

func (s square) area() float64 {
	return s.length * s.length
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	c := circle{
		radius: 2,
	}
	s := square{
		length: 5,
	}
	c.area()
	s.area()
	info(c)
	info(s)
}
