//iota is a predeclared identifier and it is reset the values to 0 whenever const keyword appears

package main

import "fmt"

const (
	a = iota //0
	b = iota //1
	c = iota //2
	d        //3
)
const (
	x = iota + 1 //0+1
	y = 100
	z = iota // 2
)

func main() {
	fmt.Println(a, "\n", b, "\n", c, "\n", d)
	fmt.Printf("%T\n%T\n%T\n%T\n", a, b, c, d)

	fmt.Println(x, "\n", y, "\n", z)
	fmt.Printf("%T\n%T\n%T\n", x, y, z)
}
