package main

import "fmt"

func main() {
	x := 1.0
	y := 2.0
	c := average(x, y)
	fmt.Println(c)
}

func average(x, y float64) float64 {
	return (x + y) / 2
}
