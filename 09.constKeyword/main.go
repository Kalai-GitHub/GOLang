package main

import "fmt"

const a = 42
const b = "James Bond"
const c = 43.42

const (
	x     = 10
	y     = "playground"
	z     = false
	p int = 85
)

func main() {
	fmt.Println(a, b, c)
	fmt.Printf("%T\n%T\n%T\n", a, b, c)

	fmt.Println(x, y, z)
	fmt.Printf("%T\n%T\n%T\n", x, y, z)

	fmt.Println(p)
	fmt.Printf("%T", p)
}
