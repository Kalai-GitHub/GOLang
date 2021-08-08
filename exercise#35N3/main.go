package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}
type shape interface {
	area() float64
}

//non-Pointer receiver and non-pointer value OR non-Pointer receiver and pointer value
// func (c circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func info(s shape) {
	fmt.Println("area: ", s.area())

}

func main() {
	c := circle{
		radius: 2,
	}

	// fmt.Println("non-Pointer receiver and non-pointer value")
	// info(c)

	// fmt.Println("non-Pointer receiver and pointer value")
	// info(&c)

	fmt.Println("Pointer receiver and pointer value")
	info(&c)

}
