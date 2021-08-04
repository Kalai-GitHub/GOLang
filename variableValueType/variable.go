package main

import "fmt"

var z = 30 //var is used when the variable declared outside main function, and scope of z is entire program(package level)

var q int //default value for the int is 0
// func main() {
// 	x := 10 // := is called short declaration operator
// 	fmt.Println("x", x)
// 	x = 20
// 	fmt.Println("again x", x)

// 	y := "string expr"
// 	fmt.Println("y as a string", y)
// 	fmt.Println("z declared outside main", z)
// 	foo()
// 	fmt.Println("q from outsite main with default value", q)
// }

func foo() {
	fmt.Println("z declared outside main", z)
}
