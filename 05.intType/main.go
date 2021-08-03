package main

import "fmt"

var a int
var b float64
var c int8

func main() {
	x := 20
	y := 20.11
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	a = 20
	b = 20.11
	fmt.Println("a, b = ", a, ",", b)
	fmt.Printf("%T\n%T\n", a, b)

	c = -128
	fmt.Println(c)
}
